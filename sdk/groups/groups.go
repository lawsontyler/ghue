package groups

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"github.com/lawsontyler/ghue/sdk/factory"
)

// Group struct
type Group struct {
	Action struct {
		Alert string `json:"alert"`
		Bri   int    `json:"bri"`
		On    bool   `json:"on"`
	} `json:"action"`
	Lights   []string `json:"lights"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	ModelID  string   `json:"modelid"`
	UniqueID string   `json:"uniqueid"`
	Class    string   `json:"class"`
}

// GetAllGroups GET on /api/<username>/groups
func GetAllGroups(client *factory.SdkClient) (map[string]*Group, *common.ErrorHUE, error) {
	groups := map[string]*Group{}
	path := fmt.Sprintf("/api/%s/groups", client.Connection.Username)

	bodyResponse, errHUE, err := internal.Request(client, "GET", http.StatusOK, path, nil)

	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		err := fmt.Errorf("HUE Error: %s", errHUE.Error.Description)
		return groups, errHUE, err
	}

	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return groups, errHUE, err
	}

	err = json.Unmarshal(bodyResponse, &groups)

	if err != nil {
		log.Errorf("Error with unmarshalling GetAllGroups: %s", err.Error())
		return groups, nil, err
	}

	return groups, nil, nil
}
