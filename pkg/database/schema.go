package database

import (
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type animeDB struct {
	gorm.Model
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
