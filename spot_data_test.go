package gocoinex

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newSpotDataClient(client Client) *SpotDataClient {
	return &SpotDataClient{
		client: client,
	}
}

func TestGetAllMarketList(t *testing.T) {
	test_case := struct {
		input string
		want  *AllMarketList
	}{
		`{"code":0,"data":["LTCBCH","ZECBCH","DASHBCH"],"message":"Ok"}`,
		&AllMarketList{
			GeneralResponse: GeneralResponse{
				Code:    0,
				Message: "Ok",
			},
			Data: []string{"LTCBCH", "ZECBCH", "DASHBCH"},
		},
	}

	// Make request
	request, _ := requestMaker("market/list", nil)

	// Make response
	response := &http.Response{
		Body:       io.NopCloser(strings.NewReader(test_case.input)),
		Status:     "200 OK",
		StatusCode: 200,
	}

	// Make mocked client and spot client
	mockedClient := new(mockedClient)
	spot_data_client := newSpotDataClient(mockedClient)

	// Set mock & Test happy path
	mockedClient.On("Do", request).Return(response, nil)
	got, _ := spot_data_client.GetAllMarketList()
	assert.EqualValues(t, test_case.want, got)

	// Set mock & Test sad path
	mockedClient.On("Do", request).Return(nil, errors.New("there is no data"))
	_, err := spot_data_client.GetAllMarketList()
	assert.Error(t, err)
}
