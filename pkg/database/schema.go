package database

import (
	"time"

	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

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

type Tracked struct {
	gorm.Model
	MalID int    `gorm:"unique"`
	Title string `gorm:"unique"`
	Type  string
}

func (Tracked) TableName() string {
	return "tracked"
}
