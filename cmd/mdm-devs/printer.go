package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maps"
	"slices"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const germanDateLayout = "02. Jan 2006, 15:04:05"

func printResult(w io.Writer, deviceMap DeviceMap) {
	switch {
	case flagAsJson:
		printJSON(w, deviceMap)
	case flagAsTable:
		printTabular(w, deviceMap)
	default:
		printList(deviceMap)
	}
}

func printJSON(w io.Writer, deviceMap DeviceMap) {
	js, err := json.MarshalIndent(deviceMap, "", "  ")
	if err != nil {
		log.Printf("ERROR marshaling data: %v, %v", deviceMap, err)
		return
	}

	fmt.Fprintln(w, string(js))
}

func printList(deviceMap DeviceMap) {
	devices := slices.Collect(maps.Keys(deviceMap))
	slices.Sort(devices)
	for _, device := range devices {
		fmt.Println(device)
	}
}

func printTabular(w io.Writer, deviceMap DeviceMap) {
	if len(deviceMap) == 0 {
		return
	}
	const maxColumnWidth = 70

	fmt.Fprintln(w)

	// disable color if STDOUT is no char device
	/*if !crate.IsCharDevice(os.Stdout) {
		text.DisableColors()
	}*/

	tw := table.NewWriter()

	// don't transform the header to UpperCase
	noUpperCaseStyle := table.StyleDefault
	noUpperCaseStyle.Format.Header = text.FormatDefault

	tw.SetStyle(noUpperCaseStyle)

	tw.AppendHeader(table.Row{
		"Device",
		"Status",
		"Last Checkin",
		"Model",
		"Android",
	})

	// range over all device names, sorted
	devices := slices.Collect(maps.Keys(deviceMap))
	slices.Sort(devices)

	for _, device := range devices {
		tw.AppendRow(table.Row{
			device,
			deviceMap[device].Status,
			deviceMap[device].LastUpdate.Format(germanDateLayout),
			deviceMap[device].Model,
			deviceMap[device].Android,
		})
	}

	tw.SortBy([]table.SortBy{{Name: "Last Checkin", Mode: table.DscAlphaNumeric}})

	fmt.Fprintln(w, tw.Render())
}
