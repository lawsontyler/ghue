package groups

import (
	"encoding/json"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"github.com/lawsontyler/ghue/sdk/common"
	log "github.com/Sirupsen/logrus"
	"fmt"
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

// CreateAPI POST on /api/groups to create a new group
func CreateAPI(connection *common.Connection, create *Create) (*CreateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(create)

	if err != nil {
		log.Errorf("Error with marshalling create: %s", err.Error())
		return &CreateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(connection, "POST", http.StatusOK, fmt.Sprintf("/api/%s/groups", connection.Username), bodyRequest)

	if errHUE != nil {
		log.Errorf("Error with requesting POST on /api/groups (create a new group), HUE Error: %s", errHUE.Error.Description)
		err := fmt.Errorf("error with requesting POST on /api/groups (create a new group), HUE Error: %s", errHUE.Error.Description)

		return &CreateResult{}, errHUE, err
	}

	if err != nil {
		log.Errorf("Error with requesting POST on /api/groups (create a new group): %s", err.Error())
		return &CreateResult{}, errHUE, err
	}

	var creates []CreateResult
	err = json.Unmarshal(bodyResponse, &creates)

	if err != nil {
		log.Errorf("Error with unmarshalling POST on /api/groups (create a new group): %s", err.Error())
		return &CreateResult{}, nil, err
	}

	return &creates[0], nil, nil
}