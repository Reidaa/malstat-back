package cmd

import (
	"fmt"
	"malstat/scrapper/pkg/jikan"
	"malstat/scrapper/pkg/utils"
	"os"

	"github.com/urfave/cli"
)

var Release struct {
	Version string
	Build   string
}

func run(top int, connectionString string, csvFile string) error {
	d, err := jikan.TopAnimeByRank(top)
	if err != nil {
		return err
	}
	if csvFile != "" {
		err = utils.AnimesToCsv(d, csvFile)
		if err != nil {
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
				Action: func(ctx *cli.Context) error {
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
				},
				Action: func(ctx *cli.Context) error {
					var connStr string = ctx.String("database")
					var csvFile string = ctx.String("csv")
					var top int = ctx.Int("top")

					if csvFile != "" {
						utils.Info.Println("Will output to", csvFile)
					}

					if connStr != "" {
						utils.Info.Println("Will try store in the given database")
					}

					utils.Info.Println("Will fetch up to the top", top, "anime")
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

// starts the command parsing process
func Run() {

	app := app()

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
