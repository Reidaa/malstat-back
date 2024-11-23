package internal

import "fmt"

var Release struct {
	Version string
	Build   string
}

func Version() {
	fmt.Printf("Build:\t%s\n", Release.Build)
}
