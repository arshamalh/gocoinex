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
		JSONResp["AllMarketList"],
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

func TestGetAllMarketInfo(t *testing.T) {
	test_case := struct {
		input string
		want  *AllMarketInfo
	}{
		JSONResp["AllMarketInfo"],
		&AllMarketInfo{
			GeneralResponse: GeneralResponse{
				Code:    0,
				Message: "Ok",
			},
			Data: map[string]MarketInfo{
				"XRPBTC": {
					TakerFeeRate:   "0.001",
					PricingName:    "BTC",
					TradingName:    "XRP",
					MinAmount:      "0.001",
					Name:           "XRPBTC",
					TradingDecimal: 8,
					MakerFeeRate:   "0.001",
					PricingDecimal: 8,
				},
				"CETUSDC": {
					TakerFeeRate:   "0.001",
					PricingName:    "USDC",
					TradingName:    "CET",
					MinAmount:      "0.001",
					Name:           "CETUSDC",
					TradingDecimal: 8,
					MakerFeeRate:   "0.001",
					PricingDecimal: 8,
				},
			},
		},
	}

	// Make request
	request, _ := requestMaker("market/info", nil)

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
	got, _ := spot_data_client.GetAllMarketInfo()
	assert.EqualValues(t, test_case.want, got)

	// Set mock & Test sad path
	mockedClient.On("Do", request).Return(nil, errors.New("there is no data"))
	_, err := spot_data_client.GetAllMarketInfo()
	assert.Error(t, err)
}
