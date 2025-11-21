package main

import (
	"flag"
	"fmt"
	"log"
	"mdm-tools/internal/apicook"
	"os"
)

var (
	flagHelp = false

	flagAdb      = false
	flagFirefox  = false
	flagSettings = false
	flagAgent    = ""
)

func usage() {
	fmt.Fprint(flag.CommandLine.Output(), Synopsis)
	flag.PrintDefaults()
	fmt.Fprint(flag.CommandLine.Output(), ExampleUsage)
}

// doCmdline, handle flags and args, return dirs to scan
func doCmdline() { // map[string]string {

	log.SetPrefix("")
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	flag.BoolVar(&flagHelp, "h", flagHelp, "help")
	flag.BoolVar(&flagAdb, "adb", flagAdb, "ADB")
	flag.BoolVar(&flagFirefox, "ff", flagFirefox, "Firefox")
	flag.BoolVar(&flagSettings, "set", flagFirefox, "Settings")
	flag.StringVar(&flagAgent, "a", flagAgent, "Specify agent")

	flag.Usage = usage
	flag.Parse()

	if flagHelp {
		usage()
		os.Exit(0)
	}

	messageType := "runApp"
	payload := ""
	devices := []string{}
	apiPath := "/rest/private/push"

	if flagAdb {
		payload = getApp("adb")
	}
	if flagFirefox {
		payload = getApp("ff")
	}

	if flagSettings {
		payload = getApp("set")
	}

	if payload == "" {
		fmt.Println("No app specified")
		os.Exit(1)
	}

	if flagAgent != "" {
		devices = append(devices, flagAgent)
	} else {
		fmt.Println("No device specified")
		os.Exit(1)
	}

	_, err := apicook.PushMsg(messageType, payload, devices, apiPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(resp)
}
