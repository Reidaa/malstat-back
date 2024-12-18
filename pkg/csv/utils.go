package csv

import (
	"fmt"
	"os"
	"time"

	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"

	"github.com/gocarina/gocsv"
)

func AnimesToCsv(animes []jikan.Anime, filename string) error {
	var file *os.File
	var err error
	var data []anime
	var now string = time.Now().UTC().Format(time.DateTime)

	if filename == "" {
		filename = "malstat.csv"
	}

	if utils.FileExists(filename) {
		file, err = os.OpenFile(filename, os.O_RDWR, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", filename, err)
		}
		err = gocsv.UnmarshalFile(file, &data)
		if err != nil {
			return fmt.Errorf("failed to unmarshal the CSV file %s: %w", filename, err)
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			return fmt.Errorf("failed to seek to start of file %s: %w", filename, err)
		}
	} else {
		file, err = os.Create(filename)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", filename, err)
		}
	}

	defer file.Close()

	for i := 0; i != len(animes); i++ {
		data = append(data, anime{
			Datetime:   now,
			MalID:      animes[i].MalID,
			Title:      animes[i].Titles[0].Title,
			Type:       animes[i].Type,
			Rank:       animes[i].Rank,
			Score:      animes[i].Score,
			ScoredBy:   animes[i].ScoredBy,
			Popularity: animes[i].Popularity,
			Members:    animes[i].Members,
			Favorites:  animes[i].Favorites,
		})
	}

	err = gocsv.MarshalFile(&data, file)
	if err != nil {
		return fmt.Errorf("failed to marshal data to CSV file %s: %w", filename, err)
	}

	return nil
}
