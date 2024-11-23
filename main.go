package main

import (
	"malstat/scrapper/cmd"
	"malstat/scrapper/internal"
)

// Build is the last GIT commit
var Build string

func main() {
	internal.Release.Build = Build
	cmd.Run()
}
