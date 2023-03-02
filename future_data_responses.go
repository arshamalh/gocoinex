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
