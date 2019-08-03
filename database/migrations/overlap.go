package migrations

import (
	"bulwark/fanap-challenge/models"
	"fmt"
)

func migrateIntersection() {
	fmt.Println("Migrating overlaps ...")

	DB.AutoMigrate(&models.Overlap{})

	fmt.Println("Overlaps migrated")
}
