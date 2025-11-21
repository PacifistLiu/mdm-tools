package main

import (
	"os"
)

const (
	SYNOPSIS      = "A script to print the device names of Legamaster displays.\n\n"
	EXAMPLE_USAGE = `

Example: lega-devs -o -m o26

Prints all online devices hostname includes  "o26"

If options -v and -m OR -d and -m would contradict each other, -v and -d take precedence over -m.
`
)

func main() {
	devs := doCmdline()

	printResult(os.Stdout, devs)
}
