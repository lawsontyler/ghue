package rules

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

// GetRule GET on /api/<username>/rules/<id>
func GetRule(client *sdk_client.SdkClient, id string) (*Rule, *common.ErrorHUE, error) {
	rule := &Rule{}
	path := fmt.Sprintf("/api/%s/rules/%s", client.Connection.Username, id)

	bodyResponse, errHUE, err := internal.Request(client, "GET", http.StatusOK, path, nil)

	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		err := fmt.Errorf("HUE Error: %s", errHUE.Error.Description)

		return rule, errHUE, err
	}

	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return rule, errHUE, err
	}

	err = json.Unmarshal(bodyResponse, &rule)

	if err != nil {
		log.Errorf("Error with unmarshalling GetRule: %s", err.Error())
		return rule, nil, err
	}

	return rule, nil, nil
}
