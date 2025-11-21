package token

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetToken is a function to generate a token for rest api calls
func GetToken(mdmServer string) (string, error) {
	username := "apiuser"
	passMD5 := "00ED03A7E7061FDB61C8F11AA03AD245"
	//password := ;)

	// Create MD5 hash of password in uppercase
	//hash := md5.Sum([]byte(password))
	//passMD5 := strings.ToUpper(hex.EncodeToString(hash[:]))
	//fmt.Println(passMD5)

	// Prepare login payload
	loginData := map[string]string{
		"login":    username,
		"password": passMD5,
	}

	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		return "", err
	}

	// Send login request
	loginURL := fmt.Sprintf("%s/rest/public/jwt/login", mdmServer)
	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse token
	var loginResp map[string]interface{}
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return "", err
	}

	token, ok := loginResp["id_token"].(string)
	if !ok || token == "" {
		return "", fmt.Errorf("API login failed, token not found")
	}
	return token, nil
}
