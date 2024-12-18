package database

import (
	"time"

	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"

	"gorm.io/gorm"
)

const (
	max_batch = 100
)

// Struct representing an anime record in the database.
type animeDB struct {
	gorm.Model
	Timestamp  time.Time
	MalID      int
	Title      string
	Type       string
	Rank       int
	Score      float32
	ScoredBy   int
	Popularity int
	Members    int
	Favorites  int
}

func (animeDB) TableName() string {
	return "animes"
}

// Inserts a list of anime data into the database.
func InsertAnimes(db *gorm.DB, animes []jikan.Anime) {
	var data []animeDB
	var now time.Time = time.Now()

	for _, v := range animes {
		data = append(data, animeDB{
			Timestamp:  now,
			MalID:      v.MalID,
			Title:      v.Titles[0].Title,
			Type:       v.Type,
			Rank:       v.Rank,
			Score:      v.Score,
			ScoredBy:   v.ScoredBy,
			Popularity: v.Popularity,
			Members:    v.Members,
			Favorites:  v.Favorites,
		})
	}

	utils.Info.Println("Writing data to database")
	db.CreateInBatches(&data, max_batch)
	utils.Info.Println("Finished writing data to database")
}
