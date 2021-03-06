package scenes

import (
	"github.com/lawsontyler/ghue/sdk/common"
	"encoding/json"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"fmt"
	log "github.com/Sirupsen/logrus"

)

// Update is the exact same as create.  Except it's Update.
type Update Create

type UpdateResult struct {
	Success map[string]interface{} `json:"success"`
}

func UpdateAPI(connection *common.Connection, sceneId string, update *Update) (*[]UpdateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(update)

	if err != nil {
		log.Errorf("Error with marshalling update: %s", err.Error())
		return &[]UpdateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(connection, "PUT", http.StatusOK, fmt.Sprintf("/api/%s/scenes/%s", connection.Username, sceneId), bodyRequest)

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

func UpdateSceneLightState(connection *common.Connection, sceneId string, lightId string, state *LightState) (*[]UpdateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(state)

	if err != nil {
		log.Errorf("Error with marshalling update light state: %s", err.Error())
		return &[]UpdateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(connection, "PUT", http.StatusOK, fmt.Sprintf("/api/%s/scenes/%s/lightstates/%s", connection.Username, sceneId, lightId), bodyRequest)

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
