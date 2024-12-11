package database

import (
	"malstat/scrapper/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connects to the database using the provided DSN.
func DB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.Info.Println("Connecting to database")
	if err != nil {
		utils.Error.Println("Failed to connect to database")
		return nil, err
	}

	return db, nil
}

// Prepares the database by migrating the necessary tables.
func Prepare(db *gorm.DB) error {
	err := db.AutoMigrate(&animeDB{}, &Tracked{})
	utils.Info.Println("Migrating the database")
	if err != nil {
		utils.Error.Println("Failed to migrate database")
		return err
	}

	return nil
}
