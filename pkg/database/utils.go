package database

import (
	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Error.Println("failed to connect to database")
		return nil, err
	}
	utils.Info.Println("connected to database")
	return db, nil
}

func Prepare(db *gorm.DB) error {
	err := db.AutoMigrate(&animeDB{})
	if err != nil {
		utils.Error.Println("failed to migrate database")
		return err
	}
	utils.Info.Println("migrated the database")
	return nil
}

func AnimesToDB(db *gorm.DB, animes []jikan.Anime) {
	var data []animeDB

	for i := 0; i != len(animes); i++ {
		data = append(data, animeDB{
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
}
