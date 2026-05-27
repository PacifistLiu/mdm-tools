package apicook

import (
	"encoding/json"
)

const (
	MdmServer      = "https://mdm-v4.media.uni-ulm.de"
	DeviceInfoPath = "/rest/private/devices/search"
)

type Welcome struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Data    Data        `json:"data"`
}

type Data struct {
	Configurations map[string]Configuration `json:"configurations"`
	Devices        Devices                  `json:"devices"`
}

type Configuration struct {
	Applications []ConfigurationApplication `json:"applications"`
	Files        []interface{}              `json:"files"`
	Name         string                     `json:"name"`
	ID           int64                      `json:"id"`
	BaseURL      *string                    `json:"baseUrl,omitempty"`
	QrCodeKey    *string                    `json:"qrCodeKey,omitempty"`
}

type ConfigurationApplication struct {
	Version     string  `json:"version"`
	Name        string  `json:"name"`
	ID          int64   `json:"id"`
	Action      int64   `json:"action"`
	URL         *string `json:"url,omitempty"`
	Pkg         string  `json:"pkg"`
	Selected    bool    `json:"selected"`
	SkipVersion bool    `json:"skipVersion"`
}

type Devices struct {
	Items           []Item `json:"items"`
	TotalItemsCount int64  `json:"totalItemsCount"`
}

type Item struct {
	Number          string       `json:"number"`
	ID              int64        `json:"id"`
	StatusCode      StatusCode   `json:"statusCode"`
	Groups          []Group      `json:"groups"`
	ConfigurationID int64        `json:"configurationId"`
	Info            Info         `json:"info"`
	LastUpdate      int64        `json:"lastUpdate"`
	MdmMode         bool         `json:"mdmMode"`
	KioskMode       bool         `json:"kioskMode"`
	AndroidVersion  string       `json:"androidVersion"`
	EnrollTime      int64        `json:"enrollTime"`
	Serial          string       `json:"serial"`
	LauncherVersion *string      `json:"launcherVersion,omitempty"`
	LauncherPkg     *LauncherPkg `json:"launcherPkg,omitempty"`
	Description     *string      `json:"description,omitempty"`
}

type Group struct {
	ID   int64 `json:"id"`
	Name Name  `json:"name"`
}

type Info struct {
	Files           []interface{}     `json:"files"`
	Permissions     []int64           `json:"permissions"`
	DeviceID        string            `json:"deviceId"`
	BatteryLevel    int64             `json:"batteryLevel"`
	Model           Model             `json:"model"`
	MdmMode         bool              `json:"mdmMode"`
	KioskMode       bool              `json:"kioskMode"`
	AndroidVersion  string            `json:"androidVersion"`
	Serial          string            `json:"serial"`
	Applications    []InfoApplication `json:"applications"`
	DefaultLauncher bool              `json:"defaultLauncher"`
}

type InfoApplication struct {
	Version string `json:"version"`
	Pkg     string `json:"pkg"`
}

type Name string

const (
	Legamaster6520 Name = "Legamaster 6520"
	Legamaster7530 Name = "Legamaster 7530"
	Legamaster8620 Name = "Legamaster 8620"
	Legamaster8630 Name = "Legamaster 8630"
	Legamaster8640 Name = "Legamaster 8640"
)

type Model string

const (
	Etx     Model = "ETX"
	Prima   Model = "PRIMA"
	T7An400 Model = "t7_an400"
)

type LauncherPkg string

const (
	COMHmdmLauncher LauncherPkg = "com.hmdm.launcher"
)

type StatusCode string

const (
	Green  StatusCode = "green"
	Red    StatusCode = "red"
	Yellow StatusCode = "yellow"
)

// Public Get() method to generate the Welcome struct
func (w *Welcome) Get() (Welcome, error) {
	var err error
	*w, err = makeWelcome()
	if err != nil {
		return *w, err
	}
	return *w, nil
}

func makeWelcome() (Welcome, error) {
	url := MdmServer + DeviceInfoPath
	var w Welcome

	// Request body to fetch all devices with large pageSize
	requestBody := map[string]interface{}{
		"pageSize": 1000000,
		"pageNum":  1,
		//"value":    "deviceId",
	}

	responseData, err := HttpReq(url, "POST", requestBody)
	if err != nil {
		return w, err
	}

	json.Unmarshal(responseData, &w)
	return w, nil
}
