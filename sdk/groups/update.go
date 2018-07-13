package groups

import (
	"encoding/json"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"github.com/lawsontyler/ghue/sdk/common"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

type Update struct {
	Lights []string `json:"lights"`
	Name string `json:"name"`
}

// UpdateResult struct
type UpdateResult struct {
	Success map[string]interface{} `json:"success"`
}

// UpdateAPI PUT on /api/groups/<group_id> to update a group
func UpdateAPI(client *sdk_client.SdkClient, groupId string, update *Update) (*[]UpdateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(update)

	if err != nil {
		log.Errorf("Error with marshalling update: %s", err.Error())
		return &[]UpdateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(client, "PUT", http.StatusOK, fmt.Sprintf("/api/%s/groups/%s", client.Connection.Username, groupId), bodyRequest)

	if errHUE != nil {
		log.Errorf("Error with requesting PUT on /api/groups/%s (delete a group), HUE Error: %s", groupId, errHUE.Error.Description)
		return &[]UpdateResult{}, errHUE, err
	}

	if err != nil {
		log.Errorf("Error with requesting PUT on /api/groups/%s (delete a group): %s", groupId, err.Error())
		return &[]UpdateResult{}, errHUE, err
	}

	var updates []UpdateResult
	err = json.Unmarshal(bodyResponse, &updates)

	if err != nil {
		log.Errorf("Error with unmarshalling PUT on /api/groups/%s (delete a group): %s", groupId, err.Error())
		return &[]UpdateResult{}, nil, err
	}

	return &updates, nil, nil
}