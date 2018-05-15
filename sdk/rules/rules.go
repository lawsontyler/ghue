package rules

import (
	"fmt"
	"net/http"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/common"
	"github.com/lawsontyler/ghue/sdk/internal"
	"github.com/lawsontyler/ghue/sdk/factory"
)

type Condition struct {
	Address string `json:"address"`
	Operator string `json:"operator"`
	Value *string `json:"value,omitempty"`
}

type ActionBody struct {
	On *bool `json:"on,omitempty"`
	Bri *int `json:"bri,omitempty"`
	Hue *int `json:"hue,omitempty"`
	Sat *int `json:"sat,omitempty"`
	XY *[2]float64 `json:"xy,omitempty"`
	CT *int `json:"ct,omitempty"`
	Alert *string `json:"alert,omitempty"`
	Effect *string `json:"effect,omitempty"`
	TransitionTime *int `json:"transitiontime,omitempty"`
	BriInc *int `json:"bri_inc,omitempty"`
	HueInc *int `json:"hue_inc,omitempty"`
	SatInc *int `json:"sat_inc,omitempty"`
	CTInc *int `json:"ct_inc,omitempty"`
	XYInc *float64 `json:"xy_inc,omitempty"`
	Scene *string `json:"scene,omitempty"`
}

type Action struct {
	Address string `json:"address"`
	Method string `json:"method"`
	Body ActionBody `json:"body"`
}

// Rule struct
type Rule struct {
	Actions []Action `json:"actions"`
	Conditions []Condition `json:"conditions"`
	Created        string `json:"created"`
	Lasttriggered  string `json:"lasttriggered"`
	Name           string `json:"name"`
	Owner          string `json:"owner"`
	Status         string `json:"status"`
	Timestriggered int    `json:"timestriggered"`
}

// GetAllRules GET on /api/<username>/rules
func GetAllRules(client *factory.SdkClient) (map[string]*Rule, *common.ErrorHUE, error) {
	rules := map[string]*Rule{}
	path := fmt.Sprintf("/api/%s/rules", client.Connection.Username)

	bodyResponse, errHUE, err := internal.Request(client, "GET", http.StatusOK, path, nil)

	if errHUE != nil {
		log.Errorf("HUE Error: %s", errHUE.Error.Description)
		err := fmt.Errorf("HUE Error: %s", errHUE.Error.Description)
		return rules, errHUE, err
	}

	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return rules, errHUE, err
	}

	err = json.Unmarshal(bodyResponse, &rules)

	if err != nil {
		log.Errorf("Error with unmarshalling GetAllRules: %s", err.Error())
		return rules, nil, err
	}

	return rules, nil, nil
}
