package factory

import (
	"net/http"
	"github.com/lawsontyler/ghue/sdk/common"
)

type SdkClient struct {
	Connection *common.Connection
	Client *http.Client
}

func GetSdkClient(connection *common.Connection) *SdkClient {
	httpClient := &http.Client{}

	client := &SdkClient{
		Client: httpClient,
	}

	if connection != nil {
		client.Connection = connection
	}

	return client
}