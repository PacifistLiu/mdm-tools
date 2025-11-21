package main

import (
	"fmt"
	"mdm-tools/internal/apicook"
)

func main() {
	deviceNumber := "o26-5126-display" // Replace with actual device number

	//settings
	apiPath := fmt.Sprintf("%s/rest/private/settings", apicook.MDM_SERVER)
	res, _ := apicook.HttpReq(apiPath, "GET", nil)
	fmt.Println(string(res))

	//settings
	apiPath = fmt.Sprintf("%s/rest/private/summary/devices", apicook.MDM_SERVER)
	res, _ = apicook.HttpReq(apiPath, "GET", nil)
	fmt.Println(string(res))

	// device info
	apiPath = fmt.Sprintf("%s/rest/private/devices/number/%s", apicook.MDM_SERVER, deviceNumber)

	res, _ = apicook.HttpReq(apiPath, "GET", nil)
	fmt.Println(string(res))
}
