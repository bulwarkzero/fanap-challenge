package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// Postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Config define config schema
type Config struct {
	Engine   string `json:"engine"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	SslMode  string `json:"ssl_mode"`
	Log      bool   `json:"log"`
}

// New create new database instance
func New(cnf Config) (db *gorm.DB, err error) {
	switch cnf.Engine {
	case "postgres":
		db, err = gorm.Open(cnf.Engine, fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			cnf.Host,
			cnf.Port,
			cnf.Username,
			cnf.DbName,
			cnf.Password,
			cnf.SslMode))
	case "mysql":
	case "mssql":
	default:
		db, err = gorm.Open(cnf.Engine, fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
			cnf.Host,
			cnf.Port,
			cnf.Username,
			cnf.DbName,
			cnf.Password,
			cnf.SslMode))
	}

	if err != nil {
		return nil, err
	}

	db.LogMode(cnf.Log)

	return db, err
}
