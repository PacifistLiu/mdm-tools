package apicook

import (
	"fmt"
)

func PushMsg(messageType string, payload string, deviceNumbers []string, apiPath string) (string, error) {
	pushPayload := map[string]interface{}{
		"messageType":   messageType, //"runApp",
		"payload":       payload,     //`{"pkg": "com.example.adb"}`,
		"broadcast":     false,
		"deviceNumbers": deviceNumbers, //[]string{"n24-132-display"},
	}
	pushURL := fmt.Sprintf("%s%s", MDM_SERVER, apiPath)

	pushBody, err := HttpReq(pushURL, "POST", pushPayload)
	if err != nil {
		return "", fmt.Errorf("Error reading push response:", err)
	}
	return fmt.Sprintf("Push response: %s", string(pushBody)), nil
}
