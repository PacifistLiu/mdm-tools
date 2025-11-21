package apicook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mdm-tools/internal/token"
	"net/http"
)

// HttpReq is a function to make and retrieve rest api get and post requests
func HttpReq(apiPath string, requestMethod string, requestBody map[string]interface{}) ([]byte, error) {
	var req *http.Request
	var err error
	if requestBody != nil {
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(requestMethod, apiPath, bytes.NewBuffer(bodyBytes))
	} else {
		req, err = http.NewRequest(requestMethod, apiPath, nil)
		if err != nil {
			return nil, err
		}
	}

	bearerToken, err := token.GetToken(MDM_SERVER)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
