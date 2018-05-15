package scenes

import (
	"github.com/lawsontyler/ghue/sdk/common"
	"encoding/json"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"fmt"
	log "github.com/Sirupsen/logrus"

	"github.com/lawsontyler/ghue/sdk/factory"
)

// Update is the exact same as create.  Except it's Update.
type Update Create

type UpdateResult struct {
	Success map[string]interface{} `json:"success"`
}

func UpdateAPI(client *factory.SdkClient, sceneId string, update *Update) (*[]UpdateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(update)

	if err != nil {
		log.Errorf("Error with marshalling update: %s", err.Error())
		return &[]UpdateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(client, "PUT", http.StatusOK, fmt.Sprintf("/api/%s/scenes/%s", client.Connection.Username, sceneId), bodyRequest)

	if errHUE != nil {
		log.Errorf("Error with requesting PUT on /api/scenes/%s (delete a group), HUE Error: %s", sceneId, errHUE.Error.Description)
		return &[]UpdateResult{}, errHUE, err
	}

	if err != nil {
		log.Errorf("Error with requesting PUT on /api/scenes/%s (delete a group): %s", sceneId, err.Error())
		return &[]UpdateResult{}, errHUE, err
	}

	var updates []UpdateResult
	err = json.Unmarshal(bodyResponse, &updates)

	if err != nil {
		log.Errorf("Error with unmarshalling PUT on /api/scenes/%s (delete a group): %s", sceneId, err.Error())
		return &[]UpdateResult{}, nil, err
	}

	return &updates, nil, nil
}

func UpdateSceneLightState(client *factory.SdkClient, sceneId string, lightId string, state *LightState) (*[]UpdateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(state)

	if err != nil {
		log.Errorf("Error with marshalling update light state: %s", err.Error())
		return &[]UpdateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(client, "PUT", http.StatusOK, fmt.Sprintf("/api/%s/scenes/%s/lightstates/%s", client.Connection.Username, sceneId, lightId), bodyRequest)

	if errHUE != nil {
		log.Errorf("Error with requesting PUT on /api/scenes/%s/lightstate/%s (delete a group), HUE Error: %s", sceneId, lightId, errHUE.Error.Description)
		return &[]UpdateResult{}, errHUE, err
	}

	if err != nil {
		log.Errorf("Error with requesting PUT on /api/scenes/%s/lightstate/%s (delete a group): %s", sceneId, lightId, err.Error())
		return &[]UpdateResult{}, errHUE, err
	}

	var updates []UpdateResult
	err = json.Unmarshal(bodyResponse, &updates)

	if err != nil {
		log.Errorf("Error with unmarshalling PUT on /api/scenes/%s/lightstate/%s (delete a group): %s", sceneId, lightId, err.Error())
		return &[]UpdateResult{}, nil, err
	}

	return &updates, nil, nil
}
