package sdk_client

import (
	"net/http"
	"io/ioutil"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SdkResponse struct {
	http.Response
}

func (response *SdkResponse) GetBody() []byte {
	r, _ := ioutil.ReadAll(response.Body)

	return r
}