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

//
func (c *FutureDataClient) GetPositionLevel() {}

//
func (c *FutureDataClient) GetMarketStatus() {}

//
func (c *FutureDataClient) GetAllMarketStatus() {}

//
func (c *FutureDataClient) GetMarketDepth() {}

//
func (c *FutureDataClient) GetMarketLatestTransaction() {}

//
func (c *FutureDataClient) GetMarketKLine() {}
