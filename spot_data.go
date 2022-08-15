package gocoinex

import (
	"fmt"
	"net/http"
	"net/url"
)

// SpotDataClient

type SpotDataClient struct {
	client *http.Client
}

func NewSpotDataClient() *SpotDataClient {
	return &SpotDataClient{client: &http.Client{}}
}

func (c *SpotDataClient) get(endpoint string, params Map) (*http.Response, error) {
	route, err := url.Parse(BaseURL + endpoint)
	query := route.Query()
	if err != nil {
		return nil, err
	}
	if params != nil {
		for k, v := range params {
			query.Set(k, fmt.Sprintf("%v", v))
		}
		route.RawQuery = query.Encode()
	}
	request, err := http.NewRequest("GET", route.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", UserAgent)
	return c.client.Do(request)
}

// Get all market list, applicable to spot and margin markets.
func (c *SpotDataClient) GetAllMarketList() (*AllMarketList, error) {
	raw_response, err := c.get("market/list", nil)
	if err != nil {
		return nil, err
	}
	return (&AllMarketList{}).Parse(raw_response)
}

// Acquire all market details,
// applicable to spot and margin markets.
func (c *SpotDataClient) GetAllMarketInfo() (*AllMarketInfo, error) {
	raw_response, err := c.get("market/info", nil)
	if err != nil {
		return nil, err
	}
	return (&AllMarketInfo{}).Parse(raw_response)
}

// Get detailed information on a single spot or margin market.
func (c *SpotDataClient) GetSingleMarketInfo(market string) (*SingleMarketInfo, error) {
	raw_response, err := c.get("market/detail", Map{
		"market": market,
	})
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
	raw_response, err := c.get("market/depth", Map{
		"market": market,
		"merge":  depth,
		"limit":  limit,
	})
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
func (c *SpotDataClient) GetSingleMarketStatistics(market string) (*SingleMarketStatistics, error) {
	raw_response, err := c.get("market/ticker", Map{
		"market": market,
	})
	if err != nil {
		return nil, err
	}
	return (&SingleMarketStatistics{}).Parse(raw_response)
}

// Get all market statistics, applicable to spot and margin markets.
func (c *SpotDataClient) GetAllMarketStatistics() (*AllMarketStatistics, error) {
	raw_response, err := c.get("market/ticker/all", Map{})
	if err != nil {
		return nil, err
	}
	return (&AllMarketStatistics{}).Parse(raw_response)
}

// Get the exchange rate of all cryptocurrencies to USD
func (c *SpotDataClient) GetCurrencyRate() (*CurrencyRate, error) {
	raw_response, err := c.get("common/currency/rate", Map{})
	if err != nil {
		return nil, err
	}
	return (&CurrencyRate{}).Parse(raw_response)
}

// Get all asset allocation
func (c *SpotDataClient) GetAssetAllocation() (*AssetAllocation, error) {
	raw_response, err := c.get("common/asset/config", Map{})
	if err != nil {
		return nil, err
	}
	return (&AssetAllocation{}).Parse(raw_response)
}

// Get the list of AMM markets
func (c *SpotDataClient) GetAMMMarketList() (*AMMMarketList, error) {
	raw_response, err := c.get("amm/market", Map{})
	if err != nil {
		return nil, err
	}
	return (*SpotDataClient{}).Parse(raw_response)
}

// Get the list of margin markets
func (c *SpotDataClient) GetMarginMarketList() (*MarginMarketList, error) {
	raw_response, err := c.get("margin/market", Map{})
	if err != nil {
		return nil, err
	}
	return (*MarginMarketList{}).Parse(raw_response)
}
