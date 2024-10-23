package database

import (
	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"
	"time"

	"gorm.io/gorm"
)

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
