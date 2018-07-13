package scenes

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

// GetScene GET on /api/<username>/scenes/<id>
func GetScene(client *sdk_client.SdkClient, id string) (*Scene, *common.ErrorHUE, error) {
	scene := &Scene{}
	path := fmt.Sprintf("/api/%s/scenes/%s", client.Connection.Username, id)

	bodyResponse, errHUE, err := internal.Request(client, "GET", http.StatusOK, path, nil)

	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		err := fmt.Errorf("HUE Error: %s", errHUE.Error.Description)
		return scene, errHUE, err
	}

	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return scene, errHUE, err
	}

	err = json.Unmarshal(bodyResponse, &scene)

	if err != nil {
		log.Errorf("Error with unmarshalling GetScene: %s", err.Error())
		log.Errorf("JSON was: %s", bodyResponse)
		return scene, nil, err
	}

	return scene, nil, nil
}
