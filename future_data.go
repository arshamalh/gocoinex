package gocoinex

import "net/http"

type FutureDataClient struct {
	client *http.Client
}

func NewFutureDataClient() *FutureDataClient {
	return &FutureDataClient{client: &http.Client{}}
}

// Ping/Pong
func (c *FutureDataClient) Ping() {

}

// System Time
func (c *FutureDataClient) GetSystemTime() {

}

// Market List
func (c *FutureDataClient) GetMarketList() {

}

// Position Level
func (c *FutureDataClient) GetPositionLevel() {

}

// Market Status
func (c *FutureDataClient) GetMarketStatus(market string) {

}

// All Market Status
func (c *FutureDataClient) GetAllMarketStatus() {

}

// Market Depth
func (c *FutureDataClient) GetMarketDepth(market string, merge string, limit int) {

}

// Latest Transaction in the Market
func (c *FutureDataClient) GetMarketLatestTransaction(market string, last_id int, limit int) {

}

// Market K-Line
func (c *FutureDataClient) GetMarketKLine(market string, limit int, Type string) {

}

// User Transaction
func (c *FutureDataClient) GetUserTransaction(market string, side int, start_time int, end_time int, offset int, limit int) {

}
