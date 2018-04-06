package rules

import (
	"github.com/lawsontyler/ghue/sdk/common"
	"encoding/json"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

type Update Create

// UpdateResult struct
type UpdateResult struct {
	Success map[string]interface{} `json:"success"`
}

// UpdateAPI PUT on /api/rules/<rule_id> to update a rule
func UpdateAPI(connection *common.Connection, ruleId string, update *Update) (*[]UpdateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(update)

	log.Errorf("JSON: %s", bodyRequest)

	if err != nil {
		log.Errorf("Error with marshalling update: %s", err.Error())
		return &[]UpdateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(connection, "PUT", http.StatusOK, fmt.Sprintf("/api/%s/rules/%s", connection.Username, ruleId), bodyRequest)

	if errHUE != nil {
		log.Errorf("Error with requesting PUT on /api/groups/%s (delete a group), HUE Error: %s", ruleId, errHUE.Error.Description)
		return &[]UpdateResult{}, errHUE, err
	}

	if err != nil {
		log.Errorf("Error with requesting PUT on /api/groups/%s (delete a group): %s", ruleId, err.Error())
		return &[]UpdateResult{}, errHUE, err
	}

	var updates []UpdateResult
	err = json.Unmarshal(bodyResponse, &updates)

	if err != nil {
		log.Errorf("Error with unmarshalling PUT on /api/groups/%s (delete a group): %s", ruleId, err.Error())
		return &[]UpdateResult{}, nil, err
	}

	return &updates, nil, nil
}
