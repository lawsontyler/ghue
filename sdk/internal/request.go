package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

func initRequest(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")
}

// Request executes a request of method (POST, PUT, DELETE) on path, checks
// if return HTTP Code is equals to wantCode
func Request(client *sdk_client.SdkClient, method string, wantCode int, path string, jsonStr []byte) ([]byte, *common.ErrorHUE, error) {

	var req *http.Request

	fullURL := fmt.Sprintf("http://%s%s", client.Connection.Host, path)
	if jsonStr != nil {
		req, _ = http.NewRequest(method, fullURL, bytes.NewReader(jsonStr))
	} else {
		req, _ = http.NewRequest(method, fullURL, nil)
	}

	initRequest(req)
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var errors []common.ErrorHUE
	inError := false
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &errors)
	if err == nil && len(errors) > 0 && errors[0].Error.Description != "" {
		inError = true
	}

	if resp.StatusCode != wantCode || client.Connection.Verbose || inError {
		//log.Errorf("Response Status: %s and we want %d", resp.Status, wantCode)
		log.Errorf("In HUE Error:%t", inError)
		log.Errorf("HUE Error:%+v", errors)
		log.Errorf("Request path: %s on %s", method, fullURL)
		log.Errorf("Request: %s", string(jsonStr))
		log.Errorf("Response Headers: s%s", resp.Header)
		log.Errorf("Response Body: %s", string(body))
	}

	if inError {
		return body, &errors[0], nil
	}

	return body, nil, nil
}
