package rules

import (
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"fmt"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/factory"
)

type DeleteResult struct {
	Success string `json:"success"`
}

func DeleteAPI(client *factory.SdkClient, ruleId string) (*[]DeleteResult, *common.ErrorHUE, error) {
	bodyResponse, hueErr, err := internal.Request(client, "DELETE", http.StatusOK, fmt.Sprintf("/api/%s/rules/%s", client.Connection.Username, ruleId), nil)

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