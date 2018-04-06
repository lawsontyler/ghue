package rules

import (
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"fmt"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
)

type DeleteResult struct {
	Success map[string]interface{} `json:"success"`
}

func DeleteAPI(connection *common.Connection, ruleId string) (*[]DeleteResult, *common.ErrorHUE, error) {
	bodyResponse, hueErr, err := internal.Request(connection, "DELETE", http.StatusOK, fmt.Sprintf("/api/%s/rules/%s", connection.Username, ruleId), nil)

	if err != nil {
		return &[]DeleteResult{}, hueErr, err
	}

	var deletes []DeleteResult

	err = json.Unmarshal(bodyResponse, &deletes)
	if err != nil {
		log.Errorf("Error with unmarshalling PUT on /api/groups/%s (delete a group): %s", ruleId, err.Error())
		return &[]DeleteResult{}, nil, err
	}


	return &deletes, nil, nil

}