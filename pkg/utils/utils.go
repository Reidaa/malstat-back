package utils

import (
	"log"
	"malstat/scrapper/pkg/jikan"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type animeCSV struct {
	Date      time.Time `csv:"date"`
	Name      string    `csv:"name"`
	Rank      int       `csv:"rank"`
	Score     float32   `csv:"score"`
	Members   int       `csv:"members"`
	Favorites int       `csv:"favorites"`
}

func AnimesToCsv(animes []jikan.Anime, filename string) error {
	var file *os.File
	var err error
	var data []*animeCSV
	var now time.Time = time.Now()

	if filename == "" {
		filename = "malstat.csv"
	}

	if FileExists(filename) {
		file, err = os.OpenFile(filename, os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		err = gocsv.UnmarshalFile(file, &data)
		if err != nil {
			return err
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			return err
		}
	} else {
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
	}

	defer file.Close()

	for i := 0; i != len(animes); i++ {
		data = append(data, &animeCSV{
			now, animes[i].Titles[0].Title, animes[i].Rank, animes[i].Score, animes[i].Members, animes[i].Favorites,
		})
	}

	err = gocsv.MarshalFile(&data, file)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Info writes logs in the color blue with "INFO: " as prefix
var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)

// Warning writes logs in the color yellow with "WARNING: " as prefix
var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)

// Error writes logs in the color red with "ERROR: " as prefix
var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)

// Debug writes logs in the color cyan with "DEBUG: " as prefix
var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)
