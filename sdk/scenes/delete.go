package scenes

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

func DeleteAPI(client *factory.SdkClient, sceneId string) (*[]DeleteResult, *common.ErrorHUE, error) {
	bodyResponse, errHUE, err := internal.Request(client, "DELETE", http.StatusOK, fmt.Sprintf("/api/%s/scenes/%s", client.Connection.Username, sceneId), nil)

	if errHUE != nil {
		log.Errorf("Error with requesting DELETE on /api/scenes/%s (delete a scene), HUE Error: %s", sceneId, errHUE.Error.Description)
		return &[]DeleteResult{}, errHUE, err
	}

	if err != nil {
		log.Errorf("Error with requesting DELETE on /api/scenes/%s (delete a scene): %s", sceneId, err.Error())
		return &[]DeleteResult{}, errHUE, err
	}

	var deletes []DeleteResult

	err = json.Unmarshal(bodyResponse, &deletes)

	if err != nil {
		log.Errorf("Error with unmarshalling DELETE on /api/scenes/%s (delete a scene): %s", sceneId, err.Error())
		return &[]DeleteResult{}, nil, err
	}

	return &deletes, nil, nil
}
