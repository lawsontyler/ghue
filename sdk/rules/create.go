package rules

import (
	"github.com/lawsontyler/ghue/sdk/common"
	"encoding/json"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)


type Create struct {
	Name string `json:"name"`
	Conditions []Condition `json:"conditions"`
	Actions []Action `json:"actions"`
}

type CreateResult struct {
	Success struct {
		Id string `json:"id"`
	} `json:"success"`
}

func CreateAPI(client *sdk_client.SdkClient, create *Create) (*CreateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(create)

	log.Errorf("JSON: %s", bodyRequest)

	if err != nil {
		log.Errorf("Error with marshalling create rule: %s", err.Error())
		return &CreateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(client, "POST", http.StatusOK, fmt.Sprintf("/api/%s/rules", client.Connection.Username), bodyRequest)

	if errHUE != nil {
		log.Errorf("Error with requesting POST on /api/rules (create a new rule), HUE Error: %s", errHUE.Error.Description)
		err := fmt.Errorf("error with requesting POST on /api/rules (create a new rule), HUE Error: %s", errHUE.Error.Description)
		return &CreateResult{}, errHUE, err
	}

	if err != nil {
		log.Errorf("Error with requesting POST on /api/rules (create a new rule): %s", err.Error())
		return &CreateResult{}, errHUE, err
	}

	var creates []CreateResult
	err = json.Unmarshal(bodyResponse, &creates)

	if err != nil {
		log.Errorf("Error with unmarshalling POST on /api/rules (create a new rule): %s", err.Error())
		return &CreateResult{}, nil, err
	}

	return &creates[0], nil, nil
}