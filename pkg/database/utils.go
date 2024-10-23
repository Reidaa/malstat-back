package database

import (
	"malstat/scrapper/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Error.Println("Failed to connect to database")
		return nil, err
	}
	utils.Info.Println("Connected to database")
	return db, nil
}

func Prepare(db *gorm.DB) error {
	err := db.AutoMigrate(&animeDB{}, &Tracked{})
	if err != nil {
		utils.Error.Println("Failed to migrate database")
		return err
	}
	utils.Info.Println("Migrated the database")
	return nil
}
