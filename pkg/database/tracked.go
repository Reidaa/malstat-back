package database

import (
	"malstat/scrapper/pkg/jikan"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Tracked struct {
	gorm.Model
	MalID    int    `gorm:"unique"`
	Title    string `gorm:"unique"`
	ImageUrl string `gorm:"unique"`
	Rank     int
	Type     string
}

func (Tracked) TableName() string {
	return "tracked"
}

func AddToTracked(db *gorm.DB, animes []jikan.Anime) {
	var data []Tracked

	for _, v := range animes {
		data = append(data, Tracked{
			MalID:    v.Mal_id,
			Title:    v.Titles[0].Title,
			ImageUrl: v.Images.Jpg.ImageUrl,
			Rank:     v.Rank,
			Type:     "anime",
		})
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "mal_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"image_url", "rank"}),
	}).Create(&data)

}

func RetrieveTracked(db *gorm.DB) (result []Tracked) {
	db.Find(&result)
	return
}
