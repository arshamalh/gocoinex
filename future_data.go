package gocoinex

import "net/http"

type FutureDataClient struct {
	client *http.Client
}

func NewFutureDataClient() *FutureDataClient {
	return &FutureDataClient{client: &http.Client{}}
}

//
func (c *FutureDataClient) Ping() {}

//
func (c *FutureDataClient) GetSystemTime() {}

//
func (c *FutureDataClient) GetMarketList() {}

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
