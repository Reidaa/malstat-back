package cmd

import (
	"fmt"
	"os"

	"malstat/scrapper/internal"
	"malstat/scrapper/pkg/utils"

	"github.com/urfave/cli"
)

func app() *cli.App {
	app := &cli.App{
		Name: "scrapper",
		Commands: []cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Version and Release information",
				Action: func(_ *cli.Context) error {
					internal.Version()
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

					err := internal.Scrap(top, connStr, csvFile)
					if err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:  "serve",
				Usage: "Start the REST API",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "port",
						Required: false,
						Usage:    "",
						Value:    8080,
					},
				},
				Action: func(ctx *cli.Context) error {
					if err := internal.Serve(ctx.Int("port")); err != nil {
						return err
					} else {
						return nil
					}
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
