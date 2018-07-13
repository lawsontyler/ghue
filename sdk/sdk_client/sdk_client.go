package sdk_client

import (
	"net/http"
	"github.com/lawsontyler/ghue/sdk/common"
)

type SdkClient struct {
	Connection *common.Connection
	Client HttpClient
}

func GetSdkClient(connection *common.Connection) *SdkClient {
	client := &SdkClient{
		Client: &http.Client{ Transport: &http.Transport{} },
	}

	if connection != nil {
		client.Connection = connection
	}

	return client
}

