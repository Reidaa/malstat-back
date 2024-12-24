package database

import (
	"time"

	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

// Struct representing a tracked anime record in the database.
type TrackedModel struct {
	gorm.Model
	MalID    int    `gorm:"unique;column:mal_id"`
	Title    string `gorm:"unique"`
	ImageURL string `gorm:"unique;column:image_url"`
	Rank     int
	Type     string
}

func (TrackedModel) TableName() string {
	return "tracked"
}

// Struct representing an anime record in the database.
type animeModel struct {
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

func (animeModel) TableName() string {
	return "animes"
}
