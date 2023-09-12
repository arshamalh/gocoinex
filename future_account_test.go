package gocoinex

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func newFutureAccountClient(client Client) *FutureAccountClient {
	return &FutureAccountClient{
		client: client,
	}
}

func TestGetAssetQuery(t *testing.T) {
	mockedClient := new(mockedClient)
	client := newFutureAccountClient(mockedClient)
	body := io.NopCloser(strings.NewReader("Hello world"))
	httpResponse := &http.Response{Body: body, Status: "200 OK", StatusCode: 200}
	mockedClient.On("Do").Return(httpResponse, nil)
	client.GetAssetQuery()
}
