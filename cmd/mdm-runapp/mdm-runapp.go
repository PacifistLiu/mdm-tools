package main

import (
	"os"
)

const (
	Synopsis     = "A script to print the device names.\n\n"
	ExampleUsage = `
mdm-runapp [OPTIONS] -a AGENT
`
)

func main() {

	doCmdline()
	os.Exit(0)
}
