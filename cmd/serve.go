package cmd

import (
	"github.com/urfave/cli"
)

const DefaultPort = 8080

type Server interface {
	Start(port int) error
}

var ServeCmd = &cli.Command{
	Name:  "serve",
	Usage: "Start the REST API",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:     "port",
			Required: false,
			Usage:    "",
			Value:    DefaultPort,
		},
	},
	Action: runServe,
}

func runServe(ctx *cli.Context) error {
	// var port = ctx.Int("port")

	// server := Server()

	// err := server.St
	// if err := serve(port); err != nil {
	// 	return err
	// }

	return nil
}
