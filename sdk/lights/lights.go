package lights

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
)

// Light struct
type Light struct {
	Manufacturername string        `json:"manufacturername"`
	Modelid          string        `json:"modelid"`
	Name             string        `json:"name"`
	State            *State        `json:"state"`
	Swversion        string        `json:"swversion"`
	Type             string        `json:"type"`
	Uniqueid         string        `json:"uniqueid"`
	Pointsymbol      []interface{} `json:"pointsymbol"`
}

// GetAllLights GET on /api/<username>/lights
func GetAllLights(connection *common.Connection) (map[string]*Light, *common.ErrorHUE, error) {
	lights := map[string]*Light{}
	path := fmt.Sprintf("/api/%s/lights", connection.Username)
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		return lights, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return lights, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &lights)
	if err != nil {
		log.Errorf("Error with unmarshalling GetAllLights: %s", err.Error())
		return lights, nil, err
	}
	return lights, nil, nil
}
