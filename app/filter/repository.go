package filter

import "bulwark/fanap-challenge/models"

// OverlapRepository overlaps repository interface
type OverlapRepository interface {
	Store(*models.Overlap) (*models.Overlap, error)
	// BulkStore([]*models.Overlap) ([]*models.Overlap, error)
	Fetch() ([]*models.Overlap, error)
	GetFirstUnique() (*models.Overlap, error)
}
