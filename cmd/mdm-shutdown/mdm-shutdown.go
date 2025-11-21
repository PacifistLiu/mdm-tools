package main

import (
	"os"
)

const (
	Synopsis     = "A script to shutdown android devices.\n\n"
	ExampleUsage = `
mdm-shutdown [OPTIONS] -a AGENT
`
)

func main() {

	doCmdline()
	os.Exit(0)
}
