package groups

import (
	"encoding/json"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"github.com/lawsontyler/ghue/sdk/common"
	log "github.com/Sirupsen/logrus"
)

type Create struct {
	Lights []string `json:"lights"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// CreateResult struct
type CreateResult struct {
	Success struct {
		Id int `json:"id"`
	} `json:"success"`
}

// CreateAPI POST on /api to create a new user
func CreateAPI(connection *common.Connection, create *Create) (*CreateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(create)
	if err != nil {
		log.Errorf("Error with marshalling create: %s", err.Error())
		return &CreateResult{}, nil, err
	}
	bodyResponse, errHUE, err := internal.Request(connection, "POST", http.StatusOK, "/api/", bodyRequest)
	if errHUE != nil {
		log.Errorf("Error with requesting POST on /api (create a new group), HUE Error: %s", errHUE.Error.Description)
		return &CreateResult{}, errHUE, err
	}
	if err != nil {
		log.Errorf("Error with requesting POST on /api (create a new group): %s", err.Error())
		return &CreateResult{}, errHUE, err
	}
	var creates []CreateResult
	err = json.Unmarshal(bodyResponse, &creates)
	if err != nil {
		log.Errorf("Error with unmarshalling POST on /api (create a new group): %s", err.Error())
		return &CreateResult{}, nil, err
	}
	return &creates[0], nil, nil

}