package cmd

import (
	"fmt"
	"os"

	"malstat/scrapper/pkg/csv"
	"malstat/scrapper/pkg/database"
	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"

	"github.com/urfave/cli"
)

var Release struct {
	Version string
	Build   string
}

func run(top int, connectionString string, csvFile string) error {
	var data []jikan.Anime

	db, err := database.DB(connectionString)
	if err != nil {
		utils.Error.Println(err)
		return err
	}

	err = database.Prepare(db)
	if err != nil {
		utils.Error.Println(err)
		return err
	}

	if len(database.RetrieveTracked(db)) >= jikan.MaxSafeHitPerDay {
		utils.Warning.Println("Tracked anime limit reached, skipping new anime retrieval")
	} else {
		utils.Info.Println("Checking the top", top, "anime for any new entry")
		topAnime, err := jikan.TopAnimeByRank(top)
		if err != nil {
			utils.Error.Println(err)
			return err
		}
		database.UpsertTrackedAnimes(db, topAnime)
	}

	tracked := database.RetrieveTracked(db)
	utils.Info.Println("Fetching", len(tracked), "entries")

	for _, v := range tracked {
		d, err := jikan.AnimeByID(v.MalID)
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
			return err
		}
	}

	return nil
}

func app() *cli.App {
	app := &cli.App{
		Name: "scrapper",
		Commands: []cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Version and Release information",
				Action: func(_ *cli.Context) error {
					fmt.Printf("Build:\t%s\n", Release.Build)
					return nil
				},
			},
			{
				Name: "scrap",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "top",
						Required: true,
						Usage:    "Upmost anime to retrieve for storage",
					},
					&cli.StringFlag{
						Name:     "csv",
						Usage:    "Record to a csv `file`",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "db",
						Usage:    "Record to database using the given postgreSQL connection `string`",
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					var connStr string = ctx.String("db")
					var csvFile string = ctx.String("csv")
					var top int = ctx.Int("top")

					if csvFile != "" {
						utils.Info.Println("Output to", csvFile)
					}

					err := run(top, connStr, csvFile)
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
	}

	return app
}

// Run starts the command parsing process
func Run() {
	app := app()

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
