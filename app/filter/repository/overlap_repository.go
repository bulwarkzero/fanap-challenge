package repository

import (
	"bulwark/fanap-challenge/models"

	"github.com/jinzhu/gorm"
)

// OverlapRepo overlaps repository
type OverlapRepo struct {
	Conn *gorm.DB
}

// Store stores overlaps in database
func (overlapRepo *OverlapRepo) Store(m *models.Overlap) (*models.Overlap, error) {
	err := overlapRepo.Conn.Create(m).Error

	if err != nil {
		return nil, err
	}

	return m, nil
}

// Fetch fetches stored overlaps
func (overlapRepo *OverlapRepo) Fetch() ([]*models.Overlap, error) {
	overlaps := []*models.Overlap{}

	// TODO pagination support
	err := overlapRepo.Conn.Model(models.Overlap{}).
		Order("created_at asc").
		Find(&overlaps).
		Error

	return overlaps, err
}

// GetFirstUnique get first unique overlap
func (overlapRepo *OverlapRepo) GetFirstUnique() (*models.Overlap, error) {
	firstUnique := &models.Overlap{}
	err := overlapRepo.Conn.Model(models.Overlap{}).
		Order("created_at asc").
		First(&firstUnique).
		Error

	return firstUnique, err
}

// NewOverlapRepo make new overlap repo
func NewOverlapRepo(conn *gorm.DB) *OverlapRepo {
	return &OverlapRepo{
		Conn: conn,
	}
}
