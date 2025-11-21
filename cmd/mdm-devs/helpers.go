package main

import (
	"mdm-tools/internal/apicook"
	"strings"
	"time"
)

const (
	ONLINE = "green"
)

type DeviceInfo struct {
	Status     string    `json:"status"`
	Android    string    `json:"android_version"`
	Model      string    `json:"model"`
	LastUpdate time.Time `json:"last_update"`
}
type DeviceMap map[string]DeviceInfo

// getDevices() fetches info all requested devices
func getDevices() (DeviceMap, error) { //map[string]string {
	devMap := DeviceMap{}
	data := apicook.Welcome{}
	deviceMap := map[string]string{}
	var err error
	var match bool
	var exclude bool

	data, err = data.Get()
	if err != nil {
		return nil, err
	}
	//data.Get()

	if len(flagExclude) > 0 {
		exclude = true
	}

	if len(flagMatch) > 0 {
		match = true
	}

	for _, device := range data.Data.Devices.Items {

		deviceWasSkipped := false

		// skip offline devices if online was requested
		if flagOnline && string(device.StatusCode) != ONLINE {
			continue
		}

		// skip offline devices if online was requested
		if flagOffline && string(device.StatusCode) == ONLINE {
			continue
		}

		// if -v was requested, look for matchpattern and EXCLUDE it if match
		// is found. break to avoid double additions
		if exclude {
			for _, matchpattern := range flagExclude {
				if strings.Contains(device.Number, matchpattern) {
					deviceWasSkipped = true
					break
				}
			}
		}

		// if -m was requested, look for matchpattern and INCLUDE if match
		// is found. Break to avoid double additions
		if match && !deviceWasSkipped {
			for _, matchpattern := range flagMatch {
				if strings.Contains(device.Number, matchpattern) {
					deviceMap[device.Number] = string(device.StatusCode)
					devMap[device.Number] = makeDeviceInfo(string(device.StatusCode), string(device.Groups[0].Name), device.AndroidVersion, device.LastUpdate)
					break
				}
			}
		}

		// add if neither device was skipped nor matching is active
		if !match && !deviceWasSkipped {
			devMap[device.Number] = makeDeviceInfo(string(device.StatusCode), string(device.Groups[0].Name), device.AndroidVersion, device.LastUpdate)
			deviceMap[device.Number] = string(device.StatusCode)
		}
	}
	//return deviceMap
	return devMap, nil
}

// makeDeviceInfo maps status colors to on/off
func makeDeviceInfo(statusCode string, model string, androidVersion string, lastUpdate int64) DeviceInfo {
	statusMap := map[string]string{
		"green":  "on",
		"yellow": "off",
		"red":    "off",
	}
	var newDevice DeviceInfo
	newDevice.Status = statusMap[statusCode]
	newDevice.Model = model
	newDevice.Android = androidVersion
	newDevice.LastUpdate = time.UnixMilli(lastUpdate)
	return newDevice
}
