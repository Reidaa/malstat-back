package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"malstat/scrapper/pkg/jikan"
)

type Tracked struct {
	gorm.Model
	MalID    int    `gorm:"unique"`
	Title    string `gorm:"unique"`
	ImageURL string `gorm:"unique"`
	Rank     int
	Type     string
}

func (Tracked) TableName() string {
	return "tracked"
}

func UpsertTrackedAnimes(db *gorm.DB, animes []jikan.Anime) {
	var data []Tracked

	for _, v := range animes {
		data = append(data, Tracked{
			MalID:    v.MalID,
			Title:    v.Titles[0].Title,
			ImageURL: v.Images.Jpg.ImageURL,
			Rank:     v.Rank,
			Type:     "anime",
		})
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "mal_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "image_url", "rank"}),
	}).Create(&data)
}

func RetrieveTracked(db *gorm.DB) []Tracked {
	var result []Tracked

	db.Find(&result)

	return result
}
