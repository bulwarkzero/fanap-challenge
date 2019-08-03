package usecase

import (
	"bulwark/fanap-challenge/app/filter"
	"bulwark/fanap-challenge/models"
	"errors"
	"sync"
)

// FilterUsecase realize filter usecase
type FilterUsecase struct {
	OverlapRepo filter.OverlapRepository
}

// HaveOverlap check that two reactangles have overlaps or not
func (fu *FilterUsecase) HaveOverlap(main, needle *models.Rectangle) bool {
	// Check that if needle is in left or right side of main and out of main box
	if (needle.X+needle.Width) < main.X ||
		needle.X > main.X+main.Width {
		return false
	}

	// Check that if needle is in top or bottom side of main and out of main box
	if needle.Y > (main.Y+main.Height) ||
		(needle.Y+needle.Height) < main.Y {
		return false
	}

	return true
}

// FindOverlapsAndStore find overlaps and store in database
func (fu *FilterUsecase) FindOverlapsAndStore(main *models.Rectangle, haystack []*models.Rectangle) error {
	results := make(chan *models.Overlap, len(haystack))
	errs := make(chan error, len(haystack))

	var wg sync.WaitGroup

	go func() {
		for overlap := range results {
			_, e := fu.OverlapRepo.Store(overlap)
			// TODO add transaction and rollback here
			if e != nil {
				errs <- e
			}
		}
	}()

	for _, needle := range haystack {
		wg.Add(1)
		go func(_needle *models.Rectangle, _wg *sync.WaitGroup) {
			if fu.HaveOverlap(main, _needle) {
				results <- &models.Overlap{
					Rectangle: *_needle,
				}
			}
			_wg.Done()
		}(needle, &wg)
	}
	wg.Wait()
	close(results)

	select {
	case <-errs:
		return errors.New("error occurred when saving overlaps")
	default:
	}

	return nil
}

// FetchOverlaps fetch stored overlaps
func (fu *FilterUsecase) FetchOverlaps() ([]*models.Overlap, error) {
	// TODO add pagination here
	return fu.OverlapRepo.Fetch()
}

// New create new filter usecase
func New(or filter.OverlapRepository) *FilterUsecase {
	return &FilterUsecase{
		OverlapRepo: or,
	}
}
