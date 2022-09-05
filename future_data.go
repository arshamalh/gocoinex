package gocoinex

import (
	"fmt"
	"net/http"
	"net/url"
)

type FutureDataClient struct {
	client *http.Client
}

func NewFutureDataClient() *FutureDataClient {
	return &FutureDataClient{client: &http.Client{}}
}

func (c *FutureDataClient) get(endpoint string, params Map) (*http.Response, error) {
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

func (c *FutureDataClient) Ping() (*Ping, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/ping", nil)
	if err != nil {
		return nil, err
	}
	return (&Ping{}).Parse(raw_response)
}

func (c *FutureDataClient) GetSystemTime() (*SystemTime, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/time", nil)
	if err != nil {
		return nil, err
	}
	return (&SystemTime{}).Parse(raw_response)
}

func (c *FutureDataClient) GetMarketList() (*MarketList, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/list", nil)
	if err != nil {
		return nil, err
	}
	return (&MarketList{}).Parse(raw_response)
}

func (c *FutureDataClient) GetPositionLevel() (*PositionLevel, error) {
	raw_response, err := c.get(" https://api.coinex.com/perpetual/v1/market/limit_config", nil)
	if err != nil {
		return nil, err
	}
	return (&PositionLevel{}).Parse(raw_response)
}

func (c *FutureDataClient) GetMarketStatus(market string) (*MarketStatus, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/ticker", Map{
		"market": market,
	})
	if err != nil {
		return nil, err
	}
	return (&MarketStatus{}).Parse(raw_response)
}

func (c *FutureDataClient) GetAllMarketStatus() (*AllMarketStatus, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/ticker/all", nil)
	if err != nil {
		return nil, err
	}
	return (&AllMarketStatus{}).Parse(raw_response)
}

func (c *FutureDataClient) GetMarketDepth(market, depth string, limit int) (*MarketDepth, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/depth", Map{
		"market": market,
		"merge":  depth,
		"limit":  limit,
	})
	if err != nil {
		return nil, err
	}
	return (&MarketDepth{}).Parse(raw_response)
}

//
func (c *FutureDataClient) GetMarketLatestTransaction() {}

//
func (c *FutureDataClient) GetMarketKLine() {}
