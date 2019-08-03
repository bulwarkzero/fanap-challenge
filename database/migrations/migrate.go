package migrations

import (
	"bulwark/fanap-challenge/pkg/db"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// DB database instance
var DB *gorm.DB

// Init ...
func Init(config db.Config) {
	var e error

	DB, e = db.New(config)

	if e != nil {
		log.Fatalln(e)

		os.Exit(1)
	}
}

// Migrate ...
func Migrate() {
	fmt.Println("run migrations ...")
	migrateIntersection()
}
