package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"malstat/scrapper/pkg/utils"
)

func DB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.Info.Println("Connecting to database")
	if err != nil {
		utils.Error.Println("Failed to connect to database")
		return nil, err
	}

	return db, nil
}

func Prepare(db *gorm.DB) error {
	err := db.AutoMigrate(&animeDB{}, &Tracked{})
	utils.Info.Println("Migrating the database")
	if err != nil {
		utils.Error.Println("Failed to migrate database")
		return err
	}

	return nil
}
