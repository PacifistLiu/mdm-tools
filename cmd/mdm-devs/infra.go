package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	flagHelp    = false
	flagOnline  = false
	flagOffline = false

	flagMatch   multipleArgFlag
	flagExclude multipleArgFlag

	flagAsJson  = false
	flagAsTable = false
)

// custom flag to allow multiple usages upon function call. this is a slice of string
type multipleArgFlag []string

// flag packages type value interface requires String() method, actual output may only be relevant for debugging purpose, but return "" would work just fine
func (m *multipleArgFlag) String() string {
	return strings.Join(*m, ",")
}

// method to collect all args for multipleArgFlag
func (m *multipleArgFlag) Set(v string) error {
	*m = append(*m, v)
	return nil
}

func usage() {
	fmt.Fprint(flag.CommandLine.Output(), SYNOPSIS)
	flag.PrintDefaults()
	fmt.Fprint(flag.CommandLine.Output(), EXAMPLE_USAGE)
}

// doCmdline, handle flags and args, return dirs to scan
func doCmdline() DeviceMap { // map[string]string {

	log.SetPrefix("")
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	flag.BoolVar(&flagHelp, "h", flagHelp, "help")
	flag.BoolVar(&flagOnline, "on", flagOnline, "restrict output to online devices")
	flag.BoolVar(&flagOffline, "off", flagOffline, "restrict output to offline devices")

	flag.Var(&flagMatch, "m", "match ANY  of this pattern(s)")
	flag.Var(&flagExclude, "v", "match NONE of this pattern(s)")

	flag.BoolVar(&flagAsJson, "j", flagAsJson, "Print output as JSON")
	flag.BoolVar(&flagAsTable, "t", flagAsTable, "Print output as table")

	flag.Usage = usage
	flag.Parse()

	if flagHelp {
		usage()
		os.Exit(0)
	}

	deviceMap, err := getDevices()
	if err != nil {
		log.Fatal(err)
	}
	return deviceMap
}
