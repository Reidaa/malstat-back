package database

import (
	"fmt"

	"github.com/reidaa/ano/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connects to the database using the provided DSN.
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.Info.Println("Connecting to database")
	if err != nil {
		utils.Error.Println("Failed to connect to database")
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

// Prepares the database by migrating the necessary tables.
func Prepare(db *gorm.DB) error {
	utils.Info.Println("Migrating the database")

	err := db.AutoMigrate(&animeModel{}, &TrackedModel{})
	if err != nil {
		utils.Error.Println("Failed to migrate database")
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
