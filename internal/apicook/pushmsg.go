package apicook

import (
	"fmt"
)

func PushMsg(messageType string, payload string, deviceNumbers []string, apiPath string) (string, error) {
	pushPayload := map[string]interface{}{
		"messageType":   messageType,
		"payload":       payload,
		"broadcast":     false,
		"deviceNumbers": deviceNumbers,
	}
	pushURL := fmt.Sprintf("%s%s", MdmServer, apiPath)

	pushBody, err := HttpReq(pushURL, "POST", pushPayload)
	if err != nil {
		return "", fmt.Errorf("Error reading push response:", err)
	}
	return fmt.Sprintf("Push response: %s", string(pushBody)), nil
}
