package database

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Model base model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// JSONTime special time type
type JSONTime struct {
	time time.Time
}

// MarshalJSON ...
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t.time).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// Value ...
func (t JSONTime) Value() (driver.Value, error) {
	return t.time, nil
}

// Scan ...
func (t *JSONTime) Scan(value interface{}) error {
	v, _ := value.(time.Time)
	t.time = v

	return nil
}
