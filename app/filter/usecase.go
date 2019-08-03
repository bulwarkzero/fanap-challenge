package filter

import "bulwark/fanap-challenge/models"

// Usecase filter service usecase
type Usecase interface {
	HaveOverlap(main, needle *models.Rectangle) bool
	FindOverlapsAndStore(main *models.Rectangle, haystack []*models.Rectangle) error
	FetchOverlaps() ([]*models.Overlap, error)
	GetFirstUnique() (*models.Overlap, error)
}
