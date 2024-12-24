package main

import (
	"log"
	"os"

	"github.com/reidaa/ano/cmd"
	"github.com/reidaa/ano/internal/app"
)

// Populated by goreleaser during build.
var (
	Version = "unknown"
	Build   = "unknown"
	Name    = "ano"
)

type IApp interface {
	Start(args []string) error
}

func main() {
	cmd.Version.Build = Build
	cmd.Version.Version = Version

	var app IApp = app.New(Version, Build, Name)

	err := app.Start(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
