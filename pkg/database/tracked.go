package database

import (
	"malstat/scrapper/pkg/jikan"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddToTracked(db *gorm.DB, animes []jikan.Anime) {
	var data []Tracked

	for i := 0; i != len(animes); i++ {
		data = append(data, Tracked{
			MalID: animes[i].Mal_id,
			Title: animes[i].Titles[0].Title,
			Type:  "anime",
		})
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&data)

}

func RetrieveTracked(db *gorm.DB) (result []Tracked) {
	db.Find(&result)
	return
}
