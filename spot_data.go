package gocoinex

import (
	"fmt"
	"net/http"
)

type SpotDataClient struct {
	client *http.Client
}

func NewSpotDataClient() *SpotDataClient {
	return &SpotDataClient{client: &http.Client{}}
}

func (c *SpotDataClient) get(endpoint string) (*http.Response, error) {
	request, err := http.NewRequest("GET", BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", UserAgent)
	return c.client.Do(request)
}

// Get all market list, applicable to spot and margin markets.
func (c *SpotDataClient) GetAllMarketList() (*AllMarketList, error) {
	raw_response, err := c.get("market/list")
	if err != nil {
		return nil, err
	}
	return (&AllMarketList{}).Parse(raw_response)
}

// Acquire all market details,
// applicable to spot and margin markets.
func (c *SpotDataClient) GetAllMarketInfo() (*AllMarketInfo, error) {
	raw_response, err := c.get("market/info")
	if err != nil {
		return nil, err
	}
	return (&AllMarketInfo{}).Parse(raw_response)
}

// Get detailed information on a single spot or margin market,
func (c *SpotDataClient) GetSingleMarketInfo() (*SingleMarketInfo, error) {
	raw_response, err := c.get("market/detail")
	if err != nil {
		return nil, err
	}
	return (&SingleMarketInfo{}).Parse(raw_response)
}

// Acquire market depth in a single spot or margin market,
// Max depth: 50
//
// Sample depth options:
// "10", "0", "0.01", "0.001", "0.00000001"
//
// Limits: 5/10/20/50
func (c *SpotDataClient) GetMarketDepth(market, depth string, limit int) (*MarketDepth, error) {
	route := fmt.Sprintf("market/depth?market=%s&merge=%s&limit=%d", market, depth, limit)
	raw_response, err := c.get(route)
	if err != nil {
		return nil, err
	}
	return (&MarketDepth{}).Parse(raw_response)
}

// Get the latest transaction data of a single spot or margin market,
// Max return of 1000 records
func (c *SpotDataClient) GetLatestTransactionData() {}

// Get k-line data of a single spot or margin market,
// Max return of 1000 records.
func (c *SpotDataClient) GetKLineData() {}

// Get statistics on a single spot or margin market,
func (c *SpotDataClient) GetSingleMarketStatistics() {}

// Get all market statistics, applicable to spot and margin markets.
func (c *SpotDataClient) GetAllMarketStatistics() {}

// GET /common/currency/rate
// Get the exchange rate of all cryptocurrencies to USD
func (c *SpotDataClient) GetCurrencyRate() {}

// GET /common/asset/config
// Get all asset allocation
func (c *SpotDataClient) GetAssetAllocation() {}

// GET /amm/market
// Get the list of AMM markets
func (c *SpotDataClient) GetAMMMarketList() {}

// GET /margin/market
// Get the list of margin markets
func (c *SpotDataClient) GetMarginMarketList() {}
