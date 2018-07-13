package groups

import (
	"testing"
	"github.com/lawsontyler/ghue/sdk/sdk_client"
	"net/http"
	"fmt"
)

type singleGroupData struct {
	index *int
	body []byte
}

func (g *singleGroupData) Read(p []byte) (n int, err error) {
	if g.index == nil {
		i := 0
		g.index = &i
	}

	buffer := make([]byte, len(p))

	if *g.index + len(p) >= len(g.body) {
		// How many bytes do we need to read?
		bytesToRead := len(g.body) - *g.index
		n = bytesToRead
		buffer = g.body[*g.index : *g.index + bytesToRead]

		copy(p, buffer)
	} else {
		p = g.body[*g.index : *g.index + len(p)]
	}

	// I feel like there must be some better way to do this.
	i := *g.index
	i += len(p)
	g.index = &i

	println(fmt.Sprintf(""))

}
func (g *singleGroupData) Close() error { return nil }

type ClientMock struct {}

func (client *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Body: &singleGroupData{},
	}, nil
}

func TestGetAllGroups(t *testing.T) {
	client := sdk_client.SdkClient{
		Client: &ClientMock{},
	}


}