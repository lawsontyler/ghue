package sensors

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
)

// GetSensor GET on /api/<username>/sensors/<id>
func GetSensor(connection *common.Connection, id string) (*Sensor, *common.ErrorHUE, error) {
	sensor := &Sensor{}
	path := fmt.Sprintf("/api/%s/sensors/%s", connection.Username, id)
	bodyResponse, errHUE, err := internal.Request(connection, "GET", http.StatusOK, path, nil)
	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		err := fmt.Errorf("HUE Error: %s", errHUE.Error.Description)
		return sensor, errHUE, err
	}
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return sensor, errHUE, err
	}
	err = json.Unmarshal(bodyResponse, &sensor)
	if err != nil {
		log.Errorf("Error with unmarshalling GetSensor: %s", err.Error())
		return sensor, nil, err
	}
	return sensor, nil, nil
}

func GetSensorIdByName(connection *common.Connection, name string) (string, *common.ErrorHUE, error) {
	var sensorId string

	sensors, _, _ := GetAllSensors(connection)

	for aSensorId, aSensor := range sensors {
		if aSensor.Name == name {
			return aSensorId, nil, nil
		}
	}

	return sensorId, nil, nil
}