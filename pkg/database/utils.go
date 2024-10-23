package database

import (
	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"
	"time"

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
	err := db.AutoMigrate(&animeDB{})
	if err != nil {
		utils.Error.Println("Failed to migrate database")
		return err
	}
	utils.Info.Println("Migrated the database")
	return nil
}

func AnimesToDB(db *gorm.DB, animes []jikan.Anime) {
	var data []animeDB
	var now time.Time = time.Now()

	for i := 0; i != len(animes); i++ {
		data = append(data, animeDB{
			Timestamp:  now,
			MalID:      animes[i].Mal_id,
			Title:      animes[i].Titles[0].Title,
			Type:       animes[i].Type,
			Rank:       animes[i].Rank,
			Score:      animes[i].Score,
			ScoredBy:   animes[i].ScoredBy,
			Popularity: animes[i].Popularity,
			Members:    animes[i].Members,
			Favorites:  animes[i].Favorites,
		})
	}

	utils.Info.Println("Writing data to database")
	db.Create(&data)
	utils.Info.Println("Finished writing data to database")
}
