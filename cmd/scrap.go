package cmd

import (
	"fmt"
	"time"

	"github.com/reidaa/ano/pkg/database"
	"github.com/reidaa/ano/pkg/jikan"
	"github.com/reidaa/ano/pkg/utils"

	"github.com/urfave/cli"
)

var ScrapCmd = &cli.Command{
	Name: "scrap",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "top",
			Required: true,
			Usage:    "Upmost anime to retrieve for storage",
		},
		&cli.StringFlag{
			Name:     "db",
			Usage:    "Record to database using the given postgreSQL connection `string`",
			Required: false,
			Value:    "none",
		},
	},
	Action: runScrap,
}

func runScrap(ctx *cli.Context) error {
	var connStr string = ctx.String("db")
	var top int = ctx.Int("top")

	err := scrap(top, connStr)
	if err != nil {
		return fmt.Errorf("failed to scrap the data: %w", err)
	}

	return nil
}

type IDatabase interface {
	UpsertTrackedAnimes(animes []jikan.Anime)
	RetrieveTrackedAnimes() []database.TrackedModel
	InsertAnimes(animes []jikan.Anime)
}

func scrap(top int, dbURL string) error {
	var data []jikan.Anime
	var db IDatabase

	db, err := database.New(dbURL)
	if err != nil {
		return fmt.Errorf("failed to initialize database connection: %w", err)
	}

	if len(db.RetrieveTrackedAnimes()) >= jikan.MaxSafeHitPerDay {
		utils.Warning.Println("Tracked anime limit reached, skipping new anime retrieval")
	} else {
		utils.Info.Println("Checking the top", top, "anime for any new entry")
		topAnime, err := jikan.TopAnimeByRank(top)
		if err != nil {
			utils.Error.Println(err)
			return fmt.Errorf("failed retrieve the top %d anime: %w", top, err)
		}
		db.UpsertTrackedAnimes(topAnime)
	}

	tracked := db.RetrieveTrackedAnimes()
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

	db.InsertAnimes(data)

	return nil
}
