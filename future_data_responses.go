package gocoinex

import (
	"encoding/json"
	"net/http"
)

type Ping struct {
	Data string `json:"data"` // Pong

}

func (r *Ping) Parse(raw_response *http.Response) (*Ping, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type SystemTime struct {
	Data int `json:"data"` // Server time, milliseconds

}

func (r *SystemTime) Parse(raw_response *http.Response) (*SystemTime, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type MarketList struct {
	Name        string   `json:"name"`        // Market name
	Type        int      `json:"type"`        // Contract type, 1: Linear contract, 2: Inverse contract
	Stock       string   `json:"stock"`       // Base coin
	Money       string   `json:"money"`       // Price coin
	Fee_prec    int      `json:"fee_prec"`    // Rate decimal
	Stock_prec  int      `json:"stock_prec"`  // Base coin decimal
	Money_prec  int      `json:"money_prec"`  // Price coin decimal
	Multiplier  int      `json:"multiplier"`  // Multiplier
	Amount_prec int      `json:"amount_prec"` // Quantity decimal
	Amount_min  string   `json:"amount_min"`  // Minimum amount
	Tick_size   string   `json:"tick_size"`   // Price granularity
	Leverages   []string `json:"leverages"`   // Margin list
	Available   bool     `json:"available"`   // Whether the market is open

}

func (r *MarketList) Parse(raw_response *http.Response) (*MarketList, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type PositionLevel struct {
	Params0 string `json:"params0"` // amount, amount
	Params1 string `json:"params1"` // leverage, leverage
	Params2 string `json:"params2"` // mainten margin, maintenance margin rate

}

func (r *PositionLevel) Parse(raw_response *http.Response) (*PositionLevel, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type MarketStatus struct {
	Period               int    `json:"period"`               // Period
	Funding_time         int    `json:"funding_time"`         // The remaining time for the next collection of the funding rate, unit: minute
	Position_amount      string `json:"position_amount"`      // Amount
	Funding_rate_last    string `json:"funding_rate_last"`    // Last funding rate
	Funding_rate_next    string `json:"funding_rate_next"`    // Next funding rate
	Funding_rate_predict string `json:"funding_rate_predict"` // Predicted funding rate
	Last                 string `json:"last"`                 // Latest Price
	Sign_price           string `json:"sign_price"`           // Mark Price
	Index_price          string `json:"index_price"`          // Index Price
	Sell_total           string `json:"sell_total"`           // The number of ask orders in the last 1,000 transactions
	Buy_total            string `json:"buy_total"`            // The number of bid orders in the last 1,000 transactions
	Open                 string `json:"open"`                 // Opening price
	Close                string `json:"close "`               // Closing price
	High                 string `json:"high"`                 // Highest price
	Low                  string `json:"low"`                  // Lowest price
	Vol                  string `json:"vol"`                  // Amount
	Buy                  string `json:"buy"`                  // Bid1 price
	Buy_amount           string `json:"buy_amount"`           // Bid1 amount
	Sell                 string `json:"sell"`                 // Ask1 price
	Sell_amount          string `json:"sell_amount"`          // Ask1 amount
	Date                 int    `json:"date"`                 // Date timestamp
}

func (r *MarketStatus) Parse(raw_response *http.Response) (*MarketStatus, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type AllMarketStatus struct {
	Period               int    `json:"period"`               // Period
	Funding_time         int    `json:"funding_time"`         // The remaining time for the next collection of the funding rate, unit: minute
	Position_amount      string `json:"position_amount"`      // Amount
	Funding_rate_last    string `json:"funding_rate_last"`    // Last funding rate
	Funding_rate_next    string `json:"funding_rate_next"`    // Next funding rate
	Funding_rate_predict string `json:"funding_rate_predict"` // Predicted funding rate
	Last                 string `json:"last"`                 // Latest Price
	Sign_price           string `json:"sign_price"`           // Mark Price
	Index_price          string `json:"index_price"`          // Index Price
	Sell_total           string `json:"sell_total"`           // The number of ask orders in the last 1,000 transactions
	Buy_total            string `json:"buy_total"`            // The number of bid orders in the last 1,000 transactions
	Open                 string `json:"open"`                 // Opening price
	Close                string `json:"close "`               // Closing price
	High                 string `json:"high"`                 // Highest price
	Low                  string `json:"low"`                  // Lowest price
	Vol                  string `json:"vol"`                  // Amount
	Buy                  string `json:"buy"`                  // Bid1 price
	Buy_amount           string `json:"buy_amount"`           // Bid1 amount
	Sell                 string `json:"sell"`                 // Ask1 price
	Sell_amount          string `json:"sell_amount"`          // Ask1 amount
	Date                 int    `json:"date"`                 // Date timestamp
}

func (r *AllMarketStatus) Parse(raw_response *http.Response) (*AllMarketStatus, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type MarketDepthFuture struct {
	Asks0       string `json:"asks00"`      // Ask1 price
	Asks1       string `json:"asks01"`      // Ask1 amount
	Bids00      string `json:"bids00"`      // Bid1 price
	Bids01      string `json:"bids01"`      // Bid1 amount
	Last        string `json:"last"`        // Price
	Sign_price  string `json:"sign_price"`  // Mark Price
	Index_price string `json:"index_price"` // Index Price
}

func (r *MarketDepthFuture) Parse(raw_response *http.Response) (*MarketDepthFuture, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type LatestTransactionInTheMarket struct {
	Id      int    `json:"id"`      // Txid
	Type    string `json:"type"`    // Type, “buy”: buy, “sell”: sell
	Price   string `json:"price"`   // Executed Price
	Amount  string `json:"amount"`  // Amount
	Date    string `json:"date"`    // Transaction time, unit: second
	Data_ms string `json:"data_ms"` // Transaction time, unit: milliseconds
}

type MarketKLine struct {
	Data0 int    `json:"data0"` // Timestamp
	Data1 string `json:"data1"` // Opening price
	Data2 string `json:"data2"` // Closing price
	Data3 string `json:"data3"` // Highest price
	Data4 string `json:"data4"` // Lowest price
	Data5 string `json:"data5"` // Amount
	Data6 string `json:"data6"` // Value

}

type UserTransaction struct {
	Offset          int    `json:"offset"`          // Offset
	Limit           int    `json:"limit"`           // Number of query
	Id              int    `json:"id"`              // Transaction ID
	Time            int    `json:"time"`            // Timestamp
	Deal_type       int    `json:"deal_type"`       // Transaction type, 1: open position, 2: add margin, 3: reduce margin, 4: close position, 5: reduce by system, 6: liquidate position, 7: position adl
	Market          string `json:"market"`          //  Market name
	User_id         int    `json:"user_id"`         // User ID
	Deal_user_id    int    `json:"deal_user_id"`    // Counter-party user id
	Order_id        int    `json:"order_id"`        // Order id
	Deal_order_id   int    `json:"deal_order_id"`   // Counter-party order id
	Position_id     int    `json:"position_id"`     // Position id
	Side            int    `json:"side"`            // 1: sell, 2: buy
	Role            int    `json:"role"`            // 1: maker, 2: taker
	Position_type   int    `json:"position_type"`   // Margin type, 1: Isolated, 2: Cross
	Price           string `json:"price"`           // Price
	Open_price      string `json:"open_price"`      // Opening Price
	Settle_price    string `json:"settle_price"`    // Settlement Price
	Amount          string `json:"amount"`          // Amount
	Position_amount string `json:"position_amount"` // Amount
	Margin_amount   string `json:"margin_amount"`   // Position margin after transaction
	Leverage        string `json:"leverage"`        // Margin
	Deal_stock      string `json:"deal_stock"`      // Number of base coin transactions
	Deal_fee        string `json:"deal_fee"`        // Fee
	Deal_margin     string `json:"deal_margin"`     // Trading margin
	Fee_rate        string `json:"fee_rate"`        // Fee rate
	Deal_profit     string `json:"deal_profit"`     // Realized PNL
	Deal_insurance  string `json:"deal_insurance "` // Consumed or increased insurance fund

}

type AdjustLeverage struct {
	Position_type int    `json:"position_type"` // Position Type
	Leverage      string `json:"leverage"`      // Margin
}

type EstimatedAmountOfPositionsToBeOpened struct {
	Position_expect string `json:"position_expect"` // Estimated Amount of Positions To Be Opened
}

type AssetQueryFuture struct {
	Available     string `json:"available"`     // Available
	Frozen        string `json:"frozen"`        // Frozen
	Tranfer       string `json:"tranfer"`       // Available
	Balance_total string `json:"balance_total"` // Balance
	Margin        string `json:"margin"`        // Margin
	Profit_unreal string `json:"profit_unreal"` // Unrealized PNL

}

type SubmitLimitOrder struct {
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // Position id
	Stop_id          int     `json:"stop_id"`          // Stop order id
	Market           string  `json:"market "`          //  Market name
	Type             int     `json:"type "`            // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: sell, 2: buy
	Effect_type      int     `json:"effect_type"`      // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill
	User_id          int     `json:"user_id"`          // User ID
	Create_time      float64 `json:"create_time"`      // Create time
	Update_time      float64 `json:"update_time"`      // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee"`        // Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left"`             // Executed
	Deal_stock       string  `json:"deal_stock"`       // Executed value
	Deal_fee         string  `json:"deal_fee"`         // Used trading fee
	Deal_profit      string  `json:"deal_profit"`      // Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amount
	Last_deal_price  string  `json:"last_deal_price"`  // Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id"`     // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   //  Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"client_id"`        // Client id
	Leverage         string  `json:"leverage"`         // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross
}

type SubmitMarketOrder struct {
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // Position id
	Stop_id          int     `json:"stop_id"`          // Stop order id
	Market           string  `json:"market"`           // Market name
	Type             int     `json:"type "`            // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: sell, 2: buy
	Effect_type      int     `json:"effect_type"`      // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1.
	User_id          int     `json:"user_id"`          // User ID
	Create_time      float64 `json:"create_time"`      // Create time
	Update_time      float64 `json:"update_time"`      // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee"`        // Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left"`             // Executed
	Deal_stock       string  `json:"deal_stock"`       // Executed value
	Deal_fee         string  `json:"deal_fee"`         // Used trading fee
	Deal_profit      string  `json:"deal_profit"`      // Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amount
	Last_deal_price  string  `json:"last_deal_price "` // Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id"`     // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   // Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"client_id"`        // Client id
	Leverage         string  `json:"leverage"`         // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross

}

type CancelOrderInBatch struct {
	Code             int     `json:"code"`             // Error code
	Message          string  `json:"message"`          // Error message
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // Position id
	Stop_id          int     `json:"top_id"`           // Stop order id
	Market           string  `json:"market"`           // Market name
	Type             int     `json:"type"`             // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: sell, 2: buy
	Effect_type      int     `json:"effect_type "`     // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1.
	User_id          int     `json:"user_id"`          // User ID
	Create_time      float64 `json:"create_time"`      // Create time
	Update_time      float64 `json:"update_time "`     // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee "`       // Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left"`             // Executed
	Deal_stock       string  `json:"deal_stock"`       // Executed value
	Deal_fee         string  `json:"deal_fee"`         // Used trading fee
	Deal_profit      string  `json:"deal_profit"`      // 	Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amount
	Last_deal_price  string  `json:"last_deal_price"`  // 	Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id"`     // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   // Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"client_id "`       // Client id
	Leverage         string  `json:"leverage "`        // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross

}

type CancelOrder struct {
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // Position id
	Stop_id          int     `json:"stop_id "`         // Stop order id
	Market           string  `json:"market   "`        // Market name
	Type             int     `json:"code"`             // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: sell, 2: buy
	Effect_type      int     `json:"effect_type "`     // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1.
	User_id          int     `json:"user_id"`          // User ID
	Create_time      float64 `json:"create_time"`      // Create time
	Update_time      float64 `json:"update_time"`      // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee"`        //  Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left "`            // Executed
	Deal_stock       string  `json:"deal_stock"`       // Executed value
	Deal_fee         string  `json:"deal_fee"`         // Used trading fee
	Deal_profit      string  `json:"deal_profit "`     // Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amount
	Last_deal_price  string  `json:"last_deal_price"`  // Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id "`    // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   // Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"leverage"`         // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross
}

type CancelStopOrder struct {
	Order_id    int     `json:"order_id"`    // Order id
	Market      string  `json:"market"`      // Market name
	Type        int     `json:"type"`        // Order type, 1: limit order, 2: market order
	Side        int     `json:"side"`        // 1: sell, 2: buy
	Effect_type int     `json:"effect_type"` // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill.
	Stop_type   int     `json:"stop_type"`   // Trigger method 1: by latest transaction price, 2: by index price, 3: by mark price
	User_id     int     `json:"user_id "`    // User ID
	Create_time float64 `json:"create_time"` // Create time
	Update_time float64 `json:"update_time"` // Update time
	Source      string  `json:"source"`      // Source
	Stop_price  string  `json:"stop_price"`  // Stop Price
	Price       string  `json:"price"`       // Price
	Amount      string  `json:"amount"`      // Amount
	Taker_fee   string  `json:"taker_fee"`   // Taker rate
	Maker_fee   string  `json:"maker_fee"`   // Maker rate
	Client_id   string  `json:"client_id"`   // Client id
}

type QueryPendingOrders struct {
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // Position id
	Stop_id          int     `json:"stop_id"`          // Stop order id
	Market           string  `json:"market"`           // Market name
	Type             int     `json:"type"`             // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: sell, 2: buy
	Effect_type      int     `json:"effect_type"`      // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1
	User_id          int     `json:"user_id "`         // User ID
	Create_time      float64 `json:"create_time "`     // Create time
	Update_time      float64 `json:"update_time"`      // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee"`        // Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left"`             // Executed
	Deal_stock       string  `json:"deal_stock"`       // Executed value
	Deal_fee         string  `json:"deal_fee "`        // Used trading fee
	Deal_profit      string  `json:"deal_profit"`      // Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amountLatest transaction amount
	Last_deal_price  string  `json:"last_deal_price"`  // Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id"`     // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   // Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"client_id"`        // Client id
	Leverage         string  `json:"leverage"`         // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross
	Total            int     `json:"total"`            // Number of total records
	Offset           int     `json:"offset"`           // Offset
	Limit            int     `json:"limit"`            // Number of records per query
}

type OrderStatus struct {
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // Position id
	Stop_id          int     `json:"stop_id"`          // Stop order id
	Market           string  `json:"market"`           // Market name
	Type             int     `json:"type"`             // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: sell, 2: buy
	Effect_type      int     `json:"effect_type"`      // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1.
	User_id          int     `json:"user_id "`         // User ID
	Create_time      float64 `json:"create_time "`     // Create time
	Update_time      float64 `json:"update_time"`      // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee"`        // Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left"`             // Executed
	Deal_stock       string  `json:"deal_stock"`       // Executed value
	Deal_fee         string  `json:"deal_fee "`        // Used trading fee
	Deal_profit      string  `json:"deal_profit"`      // Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amountLatest transaction amount
	Last_deal_price  string  `json:"last_deal_price"`  // Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id"`     // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   // Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"client_id"`        // Client id
	Leverage         string  `json:"leverage"`         // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross
	Status           string  `json:"status"`           // Status not_deal: not executed, part_deal: partially executed, done: executed, cancel: canceled
}

type QueryPendingStopOrders struct {
	Order_id    int     `json:"order_id"`    // Order id
	Market      string  `json:"market"`      // Market name
	Type        int     `json:"type"`        // Order type, 1: limit order, 2: market order
	Side        int     `json:"side"`        // 1: sell, 2: buy
	Effect_type int     `json:"effect_type"` // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1.
	Stop_type   int     `json:"stop_type"`   // Trigger method 1: by latest transaction price, 2: by index price, 3: by mark price
	User_id     int     `json:"user_id"`     // User ID
	Create_time float64 `json:"create_time"` // Create time
	Update_time float64 `json:"update_time"` // Update time
	Source      string  `json:"source"`      // Source
	State       int     `json:"state"`       // The relationship between the stop price of a stop order and the current market price 1: Lower than the price in the current market 2: Higher than the price in the current market
	Stop_price  string  `json:"stop_price"`  // Stop Price
	Price       string  `json:"price"`       // Price
	Amount      string  `json:"amount"`      // Amount
	Taker_fee   string  `json:"taker_fee"`   // Taker rate
	Maker_fee   string  `json:"maker_fee"`   // Maker rate
	Client_id   string  `json:"client_id"`   // Client id
	Total       int     `json:"total"`       // Number of total records
	Offset      int     `json:"offset"`      // Offset
	Limit       int     `json:"limit"`       // Number of records per query

}

type StopOrderStatus struct {
	Order_id    int     `json:"order_id"`    // Order id
	Market      string  `json:"market"`      // Market name
	Type        int     `json:"type"`        // Order type, 1: limit order, 2: market order
	Side        int     `json:"side"`        // 1: sell, 2: buy
	Effect_type int     `json:"effect_type"` // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1.
	Stop_type   int     `json:"stop_type"`   // Trigger method 1: by latest transaction price, 2: by index price, 3: by mark price
	User_id     int     `json:"user_id"`     // User ID
	Create_time float64 `json:"create_time"` // Create time
	Update_time float64 `json:"update_time"` // Update time
	Source      string  `json:"source"`      // Source
	State       int     `json:"state"`       // The relationship between the stop price of a stop order and the current market price 1: Lower than the price in the current market 2: Higher than the price in the current market
	Stop_price  string  `json:"stop_price"`  // Stop Price
	Price       string  `json:"price"`       // Price
	Amount      string  `json:"amount"`      // Amount
	Taker_fee   string  `json:"taker_fee"`   // Taker rate
	Maker_fee   string  `json:"maker_fee"`   // Maker rate
	Client_id   string  `json:"client_id"`   // Client id
}

type QueryFinishedOrder struct {
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // Position id
	Stop_id          int     `json:"stop_id"`          // Stop order id
	Market           string  `json:"market"`           // Market name
	Type             int     `json:"type"`             // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: sell, 2: buy
	Effect_type      int     `json:"effect_type"`      // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1.
	User_id          int     `json:"user_id"`          // User ID
	Create_time      float64 `json:"create_time"`      // Create time
	Update_time      float64 `json:"update_time"`      // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee"`        // Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left"`             // Executed
	Deal_stock       string  `json:"deal_stock"`       // Number of base coin transactions
	Deal_fee         string  `json:"deal_fee"`         // Fee
	Deal_profit      string  `json:"deal_profit"`      // Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amount
	Last_deal_price  string  `json:"last_deal_price"`  // Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id"`     // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   // Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"client_id"`        // Client id
	Leverage         string  `json:"leverage"`         // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross
	Offset           int     `json:"offset"`           // Offset
	Limit            int     `json:"limit"`            // Number of records per query
}

type LimitClose struct {
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // Position id
	Stop_id          int     `json:"stop_id"`          // Stop order id
	Market           string  `json:"market"`           // Market name
	Type             int     `json:"type"`             // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: sell, 2: buy
	Effect_type      int     `json:"effect_type"`      // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1.
	User_id          int     `json:"user_id"`          // User ID
	Create_time      float64 `json:"create_time"`      // Create time
	Update_time      float64 `json:"update_time"`      // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee"`        // Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left"`             // Executed
	Deal_stock       string  `json:"deal_stock"`       // Number of base coin transactions
	Deal_fee         string  `json:"deal_fee"`         // Fee
	Deal_profit      string  `json:"deal_profit"`      // Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amount
	Last_deal_price  string  `json:"last_deal_price"`  // Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id"`     // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   // Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"client_id"`        // Client id
	Leverage         string  `json:"leverage"`         // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross
}

type MarketClose struct {
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // Position id
	Stop_id          int     `json:"stop_id"`          // Stop order id
	Market           string  `json:"market"`           // Market name
	Type             int     `json:"type"`             // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: sell, 2: buy
	Effect_type      int     `json:"effect_type"`      // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill. Default is 1.
	User_id          int     `json:"user_id"`          // User ID
	Create_time      float64 `json:"create_time"`      // Create time
	Update_time      float64 `json:"update_time"`      // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee"`        // Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left"`             // Executed
	Deal_stock       string  `json:"deal_stock"`       // Number of base coin transactions
	Deal_fee         string  `json:"deal_fee"`         // Fee
	Deal_profit      string  `json:"deal_profit"`      // Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amount
	Last_deal_price  string  `json:"last_deal_price"`  // Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id"`     // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   // Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"client_id"`        // Client id
	Leverage         string  `json:"leverage"`         // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross
}

type AdjustPositionMargin struct {
	Position_id           int     `json:"position_id"`           // Position id
	Create_time           float64 `json:"create_time"`           // Create time
	Update_time           float64 `json:"update_time"`           // Update time
	Market                string  `json:"market"`                // Market name
	User_id               int     `json:"user_id"`               // User ID
	Type                  int     `json:"type"`                  // Order type, 1: limit order, 2: market order
	Side                  int     `json:"side"`                  // 1: Short, 2: Long
	Amount                string  `json:"amount"`                // Amount
	Amount_max            string  `json:"amount_max"`            // ATH position amount
	Amount_max_margin     string  `json:"amount_max_margin"`     // Maximum margin amount
	Close_left            string  `json:"close_left"`            // Available Amount to Liquidate
	Open_price            string  `json:"open_price"`            // Average entry price
	Open_val              string  `json:"open_val"`              // Cumulative opening value
	Open_val_max          string  `json:"open_val_max"`          // Maximum opening value
	Open_margin           string  `json:"open_margin"`           // Margin
	Mainten_margin        string  `json:"mainten_margin"`        // Maintenance Margin Rate
	Mainten_margin_amount string  `json:"mainten_margin_amount"` // Maintenance margin
	Margin_amount         string  `json:"margin_amount"`         // Margin, Initial Margin + Margin Call - Reduced Margin
	Profit_real           string  `json:"profit_real"`           // Realized PNL
	Profit_clearing       string  `json:"profit_clearing"`       // Unsettled PNL
	Take_profit_price     string  `json:"take_profit_price"`     // Take-profit price
	Stop_loss_price       string  `json:"stop_loss_price"`       // Stop-loss price
	Taker_fee             string  `json:"taker_fee"`             // Taker fee
	Maker_fee             string  `json:"maker_fee"`             // Maker fee
	Take_profit_type      int     `json:"take_profit_type"`      // Take-profit price type, 1: transaction price, 3: mark price
	Stop_loss_type        int     `json:"stop_loss_type"`        // Stop-loss price type, 1: transaction price, 3: mark price
	Fee_asset             string  `json:"fee_asset"`             // Payment coin deducted as trading fees
	Deal_asset_fee        string  `json:"deal_asset_fee"`        // Executed trading fee
	Leverage              string  `json:"leverage"`              // Margin
	Liq_price             string  `json:"liq_price"`             // Liquidation price, when the liquidation price is greater than 1000000000000, return “Infinity”
	Bkr_price             string  `json:"bkr_price"`             // Bankruptcy price, when the bankruptcy price is greater than 1000000000000, return “Infinity”
	Profit_unreal         string  `json:"aprofit_unrea"`         // Unrealized PNL
	Settle_price          string  `json:"settle_price "`         // Settlement Price
	Settle_val            string  `json:"ettle_val"`             // Settlement Value
	Adl_sort              int     `json:"adl_sort"`              // Sort by ADL
	Total                 int     `json:"total"`                 // Number of accounts with positions
}

type UserPositions struct {
	Position_id           int     `json:"position_id"`           // Position id
	Create_time           float64 `json:"create_time"`           // Create time
	Update_time           float64 `json:"update_time"`           // Update time
	Market                string  `json:"market"`                // Market name
	User_id               int     `json:"user_id"`               // User ID
	Type                  int     `json:"type"`                  // Order type, 1: limit order, 2: market order
	Side                  int     `json:"side"`                  // 1: Short, 2: Long
	Amount                string  `json:"amount"`                // Amount
	Amount_max            string  `json:"amount_max"`            // ATH position amount
	Amount_max_margin     string  `json:"amount_max_margin"`     // Maximum margin amount
	Close_left            string  `json:"close_left"`            // Available Amount to Liquidate
	Open_price            string  `json:"open_price"`            // Average entry price
	Open_val              string  `json:"open_val"`              // Cumulative opening value
	Open_val_max          string  `json:"open_val_max"`          // Maximum opening value
	Open_margin           string  `json:"open_margin"`           // Margin
	Mainten_margin        string  `json:"mainten_margin"`        // Maintenance Margin Rate
	Mainten_margin_amount string  `json:"mainten_margin_amount"` // Maintenance margin
	Margin_amount         string  `json:"margin_amount"`         // Margin, Initial Margin + Margin Call - Reduced Margin
	Profit_real           string  `json:"profit_real"`           // Realized PNL
	Profit_clearing       string  `json:"profit_clearing"`       // Unsettled PNL
	Take_profit_price     string  `json:"take_profit_price"`     // Take-profit price
	Stop_loss_price       string  `json:"stop_loss_price"`       // Stop-loss price
	Taker_fee             string  `json:"taker_fee"`             // Taker fee
	Maker_fee             string  `json:"maker_fee"`             // Maker fee
	Take_profit_type      int     `json:"take_profit_type"`      // Take-profit price type, 1: transaction price, 3: mark price
	Stop_loss_type        int     `json:"stop_loss_type"`        // Stop-loss price type, 1: transaction price, 3: mark price
	Fee_asset             string  `json:"fee_asset"`             // Payment coin deducted as trading fees
	Deal_asset_fee        string  `json:"deal_asset_fee"`        // Executed trading fee
	Leverage              string  `json:"leverage"`              // Margin
	Liq_price             string  `json:"liq_price"`             // Liquidation price, when the liquidation price is greater than 1000000000000, return “Infinity”
	Bkr_price             string  `json:"bkr_price"`             // Bankruptcy price, when the bankruptcy price is greater than 1000000000000, return “Infinity”
	Profit_unreal         string  `json:"aprofit_unrea"`         // Unrealized PNL
	Settle_price          string  `json:"settle_price "`         // Settlement Price
	Settle_val            string  `json:"ettle_val"`             // Settlement Value
	Adl_sort              int     `json:"adl_sort"`              // Sort by ADL
	Total                 int     `json:"total"`                 // Number of accounts with positions
}

type QueryUserHistoricalFundingRate struct {
	User_id      int    `json:"user_id"`      // User ID
	Time         int    `json:"time"`         // Timestamp
	Market       string `json:"market"`       // Market
	Asset        string `json:"asset"`        // Asset name
	Type         int    `json:"type"`         // Order type, 1: limit order, 2: market order
	Position_id  int    `json:"position_id"`  // Position id
	Side         int    `json:"side"`         // 1: Short, 2: Long
	Amount       string `json:"amount"`       // Amount
	Price        string `json:"price"`        // Price
	Funding_rate string `json:"funding_rate"` // Funding rate
	Funding      string `json:"funding"`      // Funding fee
	Offset       int    `json:"offset"`       // Offset
	Limit        int    `json:"limit"`        // Number of records per query

}

type PositionStopLossSettings struct {
	Position_id           int     `json:"position_id"`           // Position id
	Create_time           float64 `json:"create_time"`           // Create time
	Update_time           float64 `json:"update_time"`           // Update time
	Market                string  `json:"market"`                // Market name
	User_id               int     `json:"user_id"`               // User ID
	Type                  int     `json:"type"`                  // Order type, 1: limit order, 2: market order
	Side                  int     `json:"side"`                  // 1: Short, 2: Long
	Amount                string  `json:"amount"`                // Amount
	Amount_max            string  `json:"amount_max"`            // ATH position amount
	Amount_max_margin     string  `json:"amount_max_margin"`     // Maximum margin amount
	Close_left            string  `json:"close_left"`            // Available Amount to Liquidate
	Open_price            string  `json:"open_price"`            // Average entry price
	Open_val              string  `json:"open_val"`              // Cumulative opening value
	Open_val_max          string  `json:"open_val_max"`          // Maximum opening value
	Open_margin           string  `json:"open_margin"`           // Margin
	Mainten_margin        string  `json:"mainten_margin"`        // Maintenance Margin Rate
	Mainten_margin_amount string  `json:"mainten_margin_amount"` // Maintenance margin
	Margin_amount         string  `json:"margin_amount"`         // Margin, Initial Margin + Margin Call - Reduced Margin
	Profit_real           string  `json:"profit_real"`           // Realized PNL
	Profit_clearing       string  `json:"profit_clearing"`       // Unsettled PNL
	Take_profit_price     string  `json:"take_profit_price"`     // Take-profit price
	Stop_loss_price       string  `json:"stop_loss_price"`       // Stop-loss price
	Taker_fee             string  `json:"taker_fee"`             // Taker fee
	Maker_fee             string  `json:"maker_fee"`             // Maker fee
	Take_profit_type      int     `json:"take_profit_type"`      // Take-profit price type, 1: transaction price, 3: mark price
	Stop_loss_type        int     `json:"stop_loss_type"`        // Stop-loss price type, 1: transaction price, 3: mark price
	Fee_asset             string  `json:"fee_asset"`             // Payment coin deducted as trading fees
	Deal_asset_fee        string  `json:"deal_asset_fee"`        // Executed trading fee
	Leverage              string  `json:"leverage"`              // Margin
	Liq_price             string  `json:"liq_price"`             // Liquidation price, when the liquidation price is greater than 1000000000000, return “Infinity”
	Bkr_price             string  `json:"bkr_price"`             // Bankruptcy price, when the bankruptcy price is greater than 1000000000000, return “Infinity”
	Profit_unreal         string  `json:"aprofit_unrea"`         // Unrealized PNL
	Settle_price          string  `json:"settle_price "`         // Settlement Price
	Settle_val            string  `json:"ettle_val"`             // Settlement Value
	Adl_sort              int     `json:"adl_sort"`              // Sort by ADL
	Total                 int     `json:"total"`                 // Number of accounts with positions

}

type PositionTakeProfitSettings struct {
	Position_id           int     `json:"position_id"`           // Position id
	Create_time           float64 `json:"create_time"`           // Create time
	Update_time           float64 `json:"update_time"`           // Update time
	Market                string  `json:"market"`                // Market name
	User_id               int     `json:"user_id"`               // User ID
	Type                  int     `json:"type"`                  // Order type, 1: limit order, 2: market order
	Side                  int     `json:"side"`                  // 1: Short, 2: Long
	Amount                string  `json:"amount"`                // Amount
	Amount_max            string  `json:"amount_max"`            // ATH position amount
	Amount_max_margin     string  `json:"amount_max_margin"`     // Maximum margin amount
	Close_left            string  `json:"close_left"`            // Available Amount to Liquidate
	Open_price            string  `json:"open_price"`            // Average entry price
	Open_val              string  `json:"open_val"`              // Cumulative opening value
	Open_val_max          string  `json:"open_val_max"`          // Maximum opening value
	Open_margin           string  `json:"open_margin"`           // Margin
	Mainten_margin        string  `json:"mainten_margin"`        // Maintenance Margin Rate
	Mainten_margin_amount string  `json:"mainten_margin_amount"` // Maintenance margin
	Margin_amount         string  `json:"margin_amount"`         // Margin, Initial Margin + Margin Call - Reduced Margin
	Profit_real           string  `json:"profit_real"`           // Realized PNL
	Profit_clearing       string  `json:"profit_clearing"`       // Unsettled PNL
	Take_profit_price     string  `json:"take_profit_price"`     // Take-profit price
	Stop_loss_price       string  `json:"stop_loss_price"`       // Stop-loss price
	Taker_fee             string  `json:"taker_fee"`             // Taker fee
	Maker_fee             string  `json:"maker_fee"`             // Maker fee
	Take_profit_type      int     `json:"take_profit_type"`      // Take-profit price type, 1: transaction price, 3: mark price
	Stop_loss_type        int     `json:"stop_loss_type"`        // Stop-loss price type, 1: transaction price, 3: mark price
	Fee_asset             string  `json:"fee_asset"`             // Payment coin deducted as trading fees
	Deal_asset_fee        string  `json:"deal_asset_fee"`        // Executed trading fee
	Leverage              string  `json:"leverage"`              // Margin
	Liq_price             string  `json:"liq_price"`             // Liquidation price, when the liquidation price is greater than 1000000000000, return “Infinity”
	Bkr_price             string  `json:"bkr_price"`             // Bankruptcy price, when the bankruptcy price is greater than 1000000000000, return “Infinity”
	Profit_unreal         string  `json:"aprofit_unrea"`         // Unrealized PNL
	Settle_price          string  `json:"settle_price "`         // Settlement Price
	Settle_val            string  `json:"ettle_val"`             // Settlement Value
	Adl_sort              int     `json:"adl_sort"`              // Sort by ADL
	Total                 int     `json:"total"`                 // Number of accounts with positions

}

type QueryMarketHistoricalFundingRate struct {
	Time              int    `json:"time"`              // Timestamp
	Market            string `json:"market"`            // Market name
	Asset             string `json:"asset"`             // Asset name
	Funding_rate      string `json:"funding_rate"`      // Funding rate
	Funding_rate_real string `json:"Funding_rate_real"` // actual funding rate
	Offset            int    `json:"offset"`            // Offset
	Limit             int    `json:"limit"`             // Number of query
}

type Modifyorder struct {
	Order_id         int     `json:"order_id"`         // Order id
	Position_id      int     `json:"position_id"`      // position id
	Stop_id          int     `json:"stop_id"`          // Stop order id
	Market           string  `json:"market"`           // Market name
	Type             int     `json:"type"`             // Order type, 1: limit order, 2: market order
	Side             int     `json:"side"`             // 1: Short, 2: Long
	Effect_type      int     `json:"effect_type"`      // Order effective type, 1: always valid, 2: immediate or cancel, 3: fill or kill
	User_id          int     `json:"user_id"`          // User ID
	Create_time      float64 `json:"create_time"`      // Create time
	Update_time      float64 `json:"update_time"`      // Update time
	Source           string  `json:"source"`           // Source
	Price            string  `json:"price"`            // Price
	Amount           string  `json:"amount"`           // Amount
	Taker_fee        string  `json:"taker_fee"`        // Taker rate
	Maker_fee        string  `json:"maker_fee"`        // Maker rate
	Left             string  `json:"left"`             // Executed
	Deal_stock       string  `json:"deal_stock"`       // Executed value
	Deal_fee         string  `json:"deal_fee"`         // Used trading fee
	Deal_profit      string  `json:"deal_profit"`      // Realized PNL
	Last_deal_amount string  `json:"last_deal_amount"` // Latest transaction amount
	Last_deal_price  string  `json:"last_deal_price"`  // Latest transaction price
	Last_deal_time   int     `json:"last_deal_time"`   // Latest transaction time
	Last_deal_id     int     `json:"last_deal_id"`     // Latest txid
	Last_deal_type   int     `json:"last_deal_type"`   // Latest transaction type
	Last_deal_role   int     `json:"last_deal_role"`   // Latest trading roles
	Client_id        string  `json:"client_id"`        // Client id
	Leverage         string  `json:"leverage"`         // Margin
	Position_type    int     `json:"position_type"`    // Margin type, 1: Isolated, 2: Cross

}

type PositionADLHistoryQuery struct {
	Amount      string `json:"amount"`      // Amount
	Deal_id     int    `json:"deal_id"`     // Transaction ID
	Market      string `json:"market"`      // Market name
	Order_id    int    `json:"order_id"`    // Order id
	Price       string `json:"price"`       // Price
	Position_id int    `json:"position_id"` // Position id
	Role        int    `json:"role"`        // 1: maker, 2: taker
	Side        int    `json:"side"`        // 1: sell, 2: buy
	Time        int    `json:"time"`        // Timestamp
}
