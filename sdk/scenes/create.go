package scenes

import (
	"github.com/lawsontyler/ghue/sdk/common"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/lawsontyler/ghue/sdk/internal"
	"net/http"
	"fmt"
	"github.com/lawsontyler/ghue/sdk/factory"
)

type Create struct {
	Lights []string `json:"lights"`
	Name string `json:"name"`
	Recycle bool `json:"recycle"`
}

type CreateResult struct {
	Success struct {
		Id string `json:"id"`
	} `json:"success"`
}

func CreateApi(client *factory.SdkClient, create *Create) (*CreateResult, *common.ErrorHUE, error) {
	bodyRequest, err := json.Marshal(create)

	if err != nil {
		log.Errorf("Error with marshalling create scene: %s", err.Error())
		return &CreateResult{}, nil, err
	}

	bodyResponse, errHUE, err := internal.Request(client, "POST", http.StatusOK, fmt.Sprintf("/api/%s/scenes", client.Connection.Username), bodyRequest)

	if errHUE != nil {
		log.Errorf("Error with requesting POST on /api/scenes (create a new scene), HUE Error: %s", errHUE.Error.Description)
		err := fmt.Errorf("error with requesting POST on /api/scenes (create a new scene), HUE Error: %s", errHUE.Error.Description)
		return &CreateResult{}, errHUE, err
	}

	if err != nil {
		log.Errorf("Error with requesting POST on /api/scenes (create a new scene): %s", err.Error())
		return &CreateResult{}, errHUE, err
	}

	var creates []CreateResult
	err = json.Unmarshal(bodyResponse, &creates)

	if err != nil {
		log.Errorf("Error with unmarshalling POST on /api/scenes (create a new scene): %s", err.Error())
		return &CreateResult{}, nil, err
	}

	return &creates[0], nil, nil
}