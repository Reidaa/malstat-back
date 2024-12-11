package database

import (
	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Struct representing a tracked anime record in the database.
type Tracked struct {
	gorm.Model
	MalID    int    `gorm:"unique;column:mal_id"`
	Title    string `gorm:"unique"`
	ImageURL string `gorm:"unique;column:image_url"`
	Rank     int
	Type     string
}

func (Tracked) TableName() string {
	return "tracked"
}

// Upserts a list of tracked anime data into the database.
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

	for _, d := range data {
		utils.Debug.Println("Upserting in database:", d)
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "mal_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"title", "image_url", "rank"}),
		})
	}
}

func RetrieveTracked(db *gorm.DB) []Tracked {
	var result []Tracked

	db.Find(&result)

	return result
}
