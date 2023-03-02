package gocoinex

type Ping struct {
	Data string `json:"data"` // Pong

}

type SystemTime struct {
	Data int `json:"data"` // Server time, milliseconds

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

type PositionLevel struct {
	Params0 string `json:"params0"` // amount, amount
	Params1 string `json:"params1"` // leverage, leverage
	Params2 string `json:"params2"` // mainten margin, maintenance margin rate

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
