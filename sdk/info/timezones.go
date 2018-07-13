package info

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

// GetAllTimezones GET on /api/<username>/info/timezones
func GetAllTimezones(client *sdk_client.SdkClient) ([]string, *common.ErrorHUE, error) {
	timezones := []string{}
	path := fmt.Sprintf("/api/" + client.Connection.Username + "/info/timezones")
	bodyResponse, errHUE, err := internal.Request(client, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return timezones, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return timezones, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &timezones)
	if err != nil {
		log.Errorf("Error with unmarshalling GetAllTimezones: %s", err.Error())
		return timezones, nil, err
	}
	return timezones, nil, nil
}
