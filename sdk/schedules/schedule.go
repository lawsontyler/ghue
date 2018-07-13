package schedules

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
)

// GetSchedule GET on /api/<username>/schedules/<id>
func GetSchedule(client *sdk_client.SdkClient, id string) (*Schedule, *common.ErrorHUE, error) {
	schedule := &Schedule{}
	path := fmt.Sprintf("/api/%s/schedules/%s", client.Connection.Username, id)
	bodyResponse, errHUE, err := internal.Request(client, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		err := fmt.Errorf("HUE Error: %s", errHUE.Error.Description)
		return schedule, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return schedule, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &schedule)
	if err != nil {
		log.Errorf("Error with unmarshalling GetSchedule: %s", err.Error())
		return schedule, nil, err
	}
	return schedule, nil, nil
}
