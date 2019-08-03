package models

import "bulwark/fanap-challenge/database"

// Overlap represent rectangles with overlap
type Overlap struct {
	ID        uint              `gorm:"primary_key" json:"id"`
	CreatedAt database.JSONTime `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"time"`
	Rectangle
}
