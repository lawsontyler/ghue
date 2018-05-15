package lights

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"github.com/lawsontyler/ghue/sdk/factory"
)

// GetLight GET on /api/<username>/lights/<id>
func GetLight(client *factory.SdkClient, id string) (*Light, *common.ErrorHUE, error) {
	light := &Light{}
	path := fmt.Sprintf("/api/%s/lights/%s", client.Connection.Username, id)
	bodyResponse, errHUE, err := internal.Request(client, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return light, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return light, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &light)
	if err != nil {
		log.Errorf("Error with unmarshalling GetLight: %s", err.Error())
		return light, nil, err
	}
	return light, nil, nil
}

func GetLightIdByName(client *factory.SdkClient, name string) (string, *common.ErrorHUE, error) {
	var lightId string

	lights, _, _ := GetAllLights(client)

	for aLightId, aLight := range lights {
		if aLight.Name == name {
			lightId = aLightId
			break
		}
	}

	return lightId, nil, nil
}