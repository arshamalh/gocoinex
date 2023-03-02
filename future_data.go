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

// Adjust Leverage
func (c *FutureDataClient) PostAdjustLeverage(market string, leverage string, position_type int, timestamp int, windowtime int) {

}

// Estimated Amount of Positions To Be Opened
func (c *FutureDataClient) PostEstimatedAmountOfPositionsToBeOpened(market string, price string, side int, timestamp int, windowtime int) {

}

// Asset Query
func (c *FutureDataClient) GetAssetQuery() {

}

// Submit Limit Order
func (c *FutureDataClient) PostSubmitLimitOrder(market string, side int, amount string, price string, effect_type int, option int, client_id string, timestamp int, windowtime int) {

}

// Submit Market Order
func (c *FutureDataClient) PostSubmitMarketOrder(market string, side int, amount string, client_id string, timestamp int, windowtime int) {

}

// Submit Stop-Limit Order
func (c *FutureDataClient) PostSubmitStopLimitOrder(market string, side int, stop_type int, amount string, stop_price string, price string, effect_type int, option int, client_id string, timestamp int, windowtime int) {

}

// Submit Stop-Market Order
func (c *FutureDataClient) PostSubmitStopMarketOrder(market string, side int, stop_type int, amount string, stop_price string, client_id string, timestamp int, windowtime int) {

}

// Cancel Order In Batch
func (c *FutureDataClient) PostCancelOrderInBatch(market string, order_ids string, timestamp int, windowtime int) {

}

// Cancel Order
func (c *FutureDataClient) PostCancelOrder(market string, order_id int, timestamp int, windowtime int) {

}

// Cancel All Orders
func (c *FutureDataClient) PostCancelAllOrders(market string, side int, timestamp int, windowtime int) {

}
