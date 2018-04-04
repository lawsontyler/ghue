package rules

import (
	"github.com/lawsontyler/ghue/sdk/common"
	"encoding/json"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	log "github.com/Sirupsen/logrus"
)

type Condition struct {
	Address string `json:"address"`
	Operator string `json:"operator"`
	Value string `json:"value,omitempty"`
}

type Action struct {
	Address string `json:"address"`
	Method string `json:"method"`
	Body map[string]interface{} `json:"body"`
}

type Create struct {
	Name string `json:"name"`
	Conditions []Condition `json:"conditions"`
	Actions []Action `json:"actions"`
}

type CreateResult struct {
	Success struct {
		Id int `json:"id"`
	} `json:"success"`
}

func createAPI(connection *common.Connection, create *Create) (*CreateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(create)

	if err != nil {
		log.Errorf("Error with marshalling create rule: %s", err.Error())
		return &CreateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(connection, "POST", http.StatusOK, "/api/rules", bodyRequest)

	if errHUE != nil {
		log.Errorf("Error with requesting POST on /api/rules (create a new rule), HUE Error: %s", errHUE.Error.Description)
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