package cmd

import (
	"errors"
	"fmt"
	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var Release struct {
	Version string
	Build   string
}

func run(connectionString string) error {
	// fmt.Printf("%s", connectionString)
	// db, err := db(connectionString)
	// if err != nil {
	// 	return err
	// }
	// db
	d, err := jikan.TopAnimeByRank(1000)
	if err != nil {
		return err
	}
	err = utils.AnimesToCsv(d, "")
	if err != nil {
		return err
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
				Action: func(ctx *cli.Context) error {
					fmt.Printf("Build:\t%s\n", Release.Build)
					return nil
				},
			},
			{
				Name: "scrap",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "database, db",
						Usage:    "connection string to the database (postgres)",
						Required: false,
					},
				},
				Action: func(ctx *cli.Context) error {
					var connStr string

					if ctx.String("database") != "" {
						connStr = ctx.String("database")
					} else {
						err := godotenv.Load()
						if err != nil {
							return errors.New("error loading .env file")
						}
						connStr = os.Getenv("SCRAPPER_DB_URL")
					}

					err := run(connStr)
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

// starts the command parsing process
func Run() {

	app := app()

	if err := app.Run(os.Args); err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		os.Exit(1)
	}
}
