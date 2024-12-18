package internal

import (
	"fmt"
	"malstat/scrapper/pkg/csv"
	"malstat/scrapper/pkg/database"
	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"
	"time"
)

func Scrap(top int, connectionString string, csvFile string) error {
	var data []jikan.Anime

	db, err := database.DB(connectionString)
	if err != nil {
		utils.Error.Println(err)
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = database.Prepare(db)
	if err != nil {
		utils.Error.Println(err)
		return fmt.Errorf("failed to prepare the database: %w", err)
	}

	if len(database.RetrieveTracked(db)) >= jikan.MaxSafeHitPerDay {
		utils.Warning.Println("Tracked anime limit reached, skipping new anime retrieval")
	} else {
		utils.Info.Println("Checking the top", top, "anime for any new entry")
		topAnime, err := jikan.TopAnimeByRank(top)
		if err != nil {
			utils.Error.Println(err)
			return fmt.Errorf("failed retrieve the top %d anime: %w", top, err)
		}
		database.UpsertTrackedAnimes(db, topAnime)
	}

	tracked := database.RetrieveTracked(db)
	utils.Info.Println("Fetching", len(tracked), "entries")

	for _, v := range tracked {
		d, err := jikan.AnimeByID(v.MalID)
		// To prevent -> 429 Too Many Requests
		time.Sleep(jikan.Cooldown)
		if err != nil {
			utils.Warning.Println(err, "| Skipping this entry")
		} else {
			data = append(data, *d)
		}
	}

	for _, v := range data {
		utils.Debug.Println(v.Titles[0].Title)
	}

	database.InsertAnimes(db, data)

	if csvFile != "" {
		err = csv.AnimesToCsv(data, csvFile)
		if err != nil {
			utils.Error.Println(err)
			return fmt.Errorf("failed write data to csv file: %w", err)
		}
	}

	return nil
}
