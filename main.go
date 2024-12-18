package main

import (
	"fmt"
	"os"

	"malstat/scrapper/cmd"

	"github.com/urfave/cli"
)

// Populated by goreleaser during build.
var (
	Version = "unknown"
	Build   = "unknown"
	Name    = "anoce"
)

const (
	port = 8080
)

func app() *cli.App {
	app := &cli.App{
		Name: Name,
		Commands: []cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Version and Release information",
				Action:  cmd.VersionCmd,
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
				Action: cmd.ScrapCmd,
			},
			{
				Name:  "serve",
				Usage: "Start the REST API",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "port",
						Required: false,
						Usage:    "",
						Value:    port,
					},
				},
				Action: cmd.ServeCmd,
			},
		},
	}

	return app
}

func main() {
	cmd.Version.Build = Build
	cmd.Version.Version = Version

	app := app()

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
