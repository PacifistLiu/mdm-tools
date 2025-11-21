package main

import (
	"flag"
	"fmt"
	"log"
	"mdm-tools/internal/apicook"
	"mdm-tools/internal/crate"
	"os"
)

var (
	flagHelp = false

	flagShutdown = false
	flagReboot   = false
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
	flag.BoolVar(&flagShutdown, "s", flagShutdown, "Shutdown specified agent")
	flag.BoolVar(&flagReboot, "r", flagReboot, "Reboot specified agent")
	flag.StringVar(&flagAgent, "a", flagAgent, "Specify agent")

	flag.Usage = usage
	flag.Parse()

	if flagHelp {
		usage()
		os.Exit(0)
	}

	messageType := "reboot"
	payload := ""
	devices := []string{}
	apiPath := "/rest/private/push"

	if flagShutdown {
		payload = "shutdown"
	}

	if flagReboot {
		payload = "reboot"
		apiPath = fmt.Sprintf("%s/rest/plugins/devicereset/public/reboot/%s", apicook.MdmServer, flagAgent)
		_, err := apicook.HttpReq(apiPath, "POST", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if payload == "" {
		fmt.Println("No command specified")
		os.Exit(1)
	}

	osStdinStats, _ := os.Stdin.Stat()
	if osStdinStats.Mode()&os.ModeNamedPipe != 0 {
		devices = crate.ReadLines(os.Stdin)
	}

	if flagAgent != "" {
		if len(devices) == 0 {
			devices = append(devices, flagAgent)
		} else {
			fmt.Println("can't read from stdin and take agent as option additionally")
			os.Exit(1)
		}
	}
	if len(devices) == 0 {
		fmt.Println("No device specified")
		os.Exit(1)
	}
	_, err := apicook.PushMsg(messageType, payload, devices, apiPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	/*_, err := mockPushMessage(messageType, payload, devices, apiPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}*/

}

func mockPushMessage(messageType string, payload string, deviceNumbers []string, apiPath string) (string, error) {
	pushPayload := map[string]interface{}{
		"messageType":   messageType, //"runApp",
		"payload":       payload,     //`{"pkg": "com.example.adb"}`,
		"broadcast":     false,
		"deviceNumbers": deviceNumbers, //[]string{"n24-132-display"},
	}
	fmt.Println(pushPayload)
	return "", nil
}
