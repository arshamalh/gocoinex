package gocoinex

import "net/http"

type FutureDataClient struct {
	client Client
}

func NewFutureDataClient() *FutureDataClient {
	return &FutureDataClient{client: &http.Client{}}
}

func (c *FutureDataClient) get(endpoint string, params Map) (*http.Response, error) {
	request, err := requestMaker(endpoint, params)
	if err != nil {
		return nil, err
	}
	return c.client.Do(request)
}

// Ping/Pong
func (c *FutureDataClient) GetPing() (*Ping, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/ping", nil)
	if err != nil {
		return nil, err
	}
	return (&Ping{}).Parse(raw_response)
}

// System Time
func (c *FutureDataClient) GetSystemTime() (*SystemTime, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/time", nil)
	if err != nil {
		return nil, err
	}
	return (&SystemTime{}).Parse(raw_response)
}

// Market List
func (c *FutureDataClient) GetMarketList() (*MarketList, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/list", nil)
	if err != nil {
		return nil, err
	}
	return (&MarketList{}).Parse(raw_response)
}

// Position Level
func (c *FutureDataClient) GetPositionLevel() (*PositionLevel, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/limit_config", nil)
	if err != nil {
		return nil, err
	}
	return (&PositionLevel{}).Parse(raw_response)
}

// Market Status
func (c *FutureDataClient) GetMarketStatus() (*MarketStatus, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/ticker", nil)
	if err != nil {
		return nil, err
	}
	return (&MarketStatus{}).Parse(raw_response)
}
