package database

import (
	"bulwark/fanap-challenge/pkg/db"
	"sync"

	"github.com/jinzhu/gorm"
)

// Singletoned db instance
var _db *gorm.DB

// Init initialize database instance
func Init(cnf db.Config) error {
	var err error
	// Singleton database initialization
	new(sync.Once).Do(func() {
		_db, err = db.New(cnf)
	})

	return err
}

// GetInstance return database singleton instance
func GetInstance() *gorm.DB {
	return _db
}
