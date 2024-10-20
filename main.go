package main

import (
	"malstat/scrapper/cmd"
)

// Build is the last GIT commit
var Build string

// func db(connectionString string) (*gorm.DB, error) {
// 	sqlDB, err := sql.Open("pgx", connectionString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: sqlDB,
// 	}), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return gormDB, nil
// }

func main() {
	cmd.Release.Build = Build
	cmd.Run()
}
