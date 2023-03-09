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

// All Market Status
func (c *FutureDataClient) GetAllMarketStatus() (*AllMarketStatus, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/ticker/all", nil)
	if err != nil {
		return nil, err
	}
	return (&AllMarketStatus{}).Parse(raw_response)
}

// Market Depth
func (c *FutureDataClient) GetMarketDepthFuture(market string, merge string, limit int) (*MarketDepthFuture, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/depth", nil)
	if err != nil {
		return nil, err
	}
	return (&MarketDepthFuture{}).Parse(raw_response)
}

// Latest Transaction in the Market
func (c *FutureDataClient) GetMarketLatestTransaction(market string, last_id int, limit int) (*MarketLatestTransaction, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/deals", nil)
	if err != nil {
		return nil, err
	}
	return (&MarketLatestTransaction{}).Parse(raw_response)
}

// Market K-Line
func (c *FutureDataClient) GetMarketKLine(market string, limit int, Type string) (*MarketKLine, error) {
	raw_response, err := c.get("https://api.coinex.com/perpetual/v1/market/kline", nil)
	if err != nil {
		return nil, err
	}
	return (&MarketKLine{}).Parse(raw_response)
}

// User Transaction
func (c *FutureDataClient) GetUserTransaction(market string, side int, start_time int, end_time int, offset int, limit int) {

}

// Adjust Leverage
func (c *FutureDataClient) PostAdjustLeverage(market string, leverage string, position_type int, timestamp int, windowtime int) {

}
