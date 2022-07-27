package gocoinex

import "net/http"

type SpotDataClient struct {
	client *http.Client
}

func NewSpotDataClient() *SpotDataClient {
	return &SpotDataClient{client: &http.Client{}}
}

//
func (c *SpotDataClient) GetAllMarketList() {
}

//
func (c *SpotDataClient) GetAllMarketInfo() {}

//
func (c *SpotDataClient) GetSingleMarketInfo() {}

//
func (c *SpotDataClient) GetMarketDepth() {}

//
func (c *SpotDataClient) GetLatestTransactionData() {}

//
func (c *SpotDataClient) GetKLineData() {}

//
func (c *SpotDataClient) GetSingleMarketStatistics() {}

//
func (c *SpotDataClient) GetAllMarketStatistics() {}

//
func (c *SpotDataClient) GetCurrencyRate() {}

//
func (c *SpotDataClient) GetAssetAllocation() {}

//
func (c *SpotDataClient) GetAMMMarketList() {}

//
func (c *SpotDataClient) GetMarginMarketList() {}
