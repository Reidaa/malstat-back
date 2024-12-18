package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var Version struct {
	Version string
	Build   string
}

// VersionCmd prints version information to stdout.
// It displays the build number and version string of the application.
// Returns nil on successful execution.
func VersionCmd(_ *cli.Context) error {
	fmt.Printf("Build:\t\t%s\n", Version.Build)
	fmt.Printf("Version:\t%s\n", Version.Version)
	return nil
}
