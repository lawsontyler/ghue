package groups

import (
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"fmt"
	log "github.com/Sirupsen/logrus"

	"encoding/json"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

type DeleteResult struct {
	Success string `json:"success"`
}

func DeleteAPI(client *sdk_client.SdkClient, groupId string) (*[]DeleteResult, *common.ErrorHUE, error) {
	bodyResponse, errHUE, err := internal.Request(client, "DELETE", http.StatusOK, fmt.Sprintf("/api/%s/groups/%s", client.Connection.Username, groupId), nil)

	if errHUE != nil {
		log.Errorf("Error with requesting DELETE on /api/groups/%s (delete a group), HUE Error: %s", groupId, errHUE.Error.Description)
		err := fmt.Errorf("error with requesting DELETE on /api/groups/%s (delete a group), HUE Error: %s", groupId, errHUE.Error.Description)

		return &[]DeleteResult{}, errHUE, err
	}

	if err != nil {
		log.Errorf("Error with requesting DELETE on /api/groups/%s (delete a group): %s", groupId, err.Error())
		return &[]DeleteResult{}, errHUE, err
	}

	var deletes []DeleteResult

	err = json.Unmarshal(bodyResponse, &deletes)

	if err != nil {
		log.Errorf("Error with unmarshalling DELETE on /api/groups/%s (delete a group): %s", groupId, err.Error())
		return &[]DeleteResult{}, nil, err
	}

	return &deletes, nil, nil

}
