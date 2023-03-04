package gocoinex

import "net/http"

type FutureDataClient struct {
	client Client
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

// Cancel Stop Order
func (c *FutureDataClient) PostCancelStopOrder(market string, order_id int, timestamp int, windowtime int) {

}

// Cancel All Stop Orders
func (c *FutureDataClient) PostCancelAllStopOrders(market string, side int, timestamp int, windowtime int) {

}

// Query Pending Orders
func (c *FutureDataClient) GetQueryPendingOrders(market string, side int, client_id string, offset int, limit int, timestamp int, windowtime int) {

}

// Order Status
func (c *FutureDataClient) GetOrderStatus(market string, order_id int, timestamp int, windowtime int) {

}

// Query Pending Stop Orders
func (c *FutureDataClient) GetQueryPendingStopOrders(market string, side int, client_id string, timestamp int, windowtime int) {

}

// Stop Order Status
func (c *FutureDataClient) GetStopOrderStatus(market string, order_id int, timestamp int, windowtime int) {

}

// Query Finished Order
func (c *FutureDataClient) GetQueryFinishedOrder(market string, side int, start_time int, end_time int, offset int, limit int, timestamp int, windowtime int) {

}

// Limit Close
func (c *FutureDataClient) PostLimitClose(market string, position_id int, amount string, price string, effect_type int, option int, client_id string, timestamp int, windowtim int) {

}

// Market Close
func (c *FutureDataClient) PostMarketClose(market string, position_id int, amount string, client_id string, timestamp int, windowtime int) {

}

// Adjust Position Margin
func (c *FutureDataClient) PostAdjustPositionMargin(market string, amount string, Type int, timestamp int, windowtime int) {

}

// User Positions
func (c *FutureDataClient) GetUserPositions(market string, timestamp int, windowtime int) {

}

// Query User Historical Funding Rate
func (c *FutureDataClient) GetQueryUserHistoricalFundingRate(market string, start_time int, end_time int, offset int, limit int, timestamp int, windowtime int) {

}

// Position Stop-Loss Settings
func (c *FutureDataClient) PostPositionStopLossSettings(market string, position_id int, stop_type int, stop_loss_price string, timestamp int, windowtime int) {

}

// Position Take-Profit Settings
func (c *FutureDataClient) PostPositionTakeProfitSettings(market string, position_id int, stop_type int, take_profit_price string, timestamp int, windowtime int) {

}

// Market Close All
func (c *FutureDataClient) PostMarketCloseAll(market string, position_id int, timestamp int, windowtime int) {

}

// Query Market Historical Funding Rate
func (c *FutureDataClient) GetQueryMarketHistoricalFundingRate(market string, start_time int, end_time int, offset int, limit int) {

}

// Modify order
func (c *FutureDataClient) PostModifyOrder(market string, order_id int, amount string, price string, timestamp int, windowtime int) {

}

// Modify Stop Order
func (c *FutureDataClient) PostModifyStopOrder(market string, order_id int, amount string, price string, stop_price string, timestamp int, windowtime int) {

}

// Position ADL History Query
func (c *FutureDataClient) GetPositionADLHistoryQuery(market string, position_id int, start_time int, end_time int, offset int, limit int) {

}

// Cancel Orders By client_id
func (c *FutureDataClient) PostCancelOrdersByClientid(market string, client_id string, timestamp int, windowtime int) {

}

// Cancel Stop Orders By client_id
func (c *FutureDataClient) PostCancelStopOrdersByClientid(market string, client_id string, timestamp int, windowtime int) {

}
