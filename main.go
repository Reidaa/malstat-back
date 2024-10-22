package main

import (
	"malstat/scrapper/cmd"
)

// Build is the last GIT commit
var Build string

func main() {
	cmd.Release.Build = Build
	cmd.Run()
}
