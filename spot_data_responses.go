package gocoinex

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type AllMarketList struct {
	GeneralResponse
	Data []string
}

func (r *AllMarketList) Parse(raw_response *http.Response) (*AllMarketList, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type MarketInfo struct {
	Name           string `json:"name"`            // Market name
	TakerFeeRate   string `json:"taker_fee_rate"`  // Taker rate
	MakerFeeRate   string `json:"maker_fee_rate"`  // Maker rate
	MinAmount      string `json:"min_amount"`      // Min transaction volume
	TradingName    string `json:"trading_name"`    // Trading currency
	TradingDecimal int    `json:"trading_decimal"` // Trading currency decimal
	PricingName    string `json:"pricing_name"`    // Pricing currency
	PricingDecimal int    `json:"pricing_decimal"` // Pricing currency decimal
}

type AllMarketInfo struct {
	GeneralResponse
	Data map[string]MarketInfo
}

func (r *AllMarketInfo) Parse(raw_response *http.Response) (*AllMarketInfo, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type SingleMarketInfo struct {
	GeneralResponse
	Data MarketInfo
}

func (r *SingleMarketInfo) Parse(raw_response *http.Response) (*SingleMarketInfo, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type BidAsk struct {
	Price  float64 `json:"price"`  // Bid or Ask price
	Amount float64 `json:"amount"` // Bid or Ask amount
}

func (t *BidAsk) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	var bisask []string
	if err := json.Unmarshal(data, &bisask); err != nil {
		return err
	}
	price, err := strconv.ParseFloat(bisask[0], 64)
	if err != nil {
		return err
	}
	amount, err := strconv.ParseFloat(bisask[1], 64)
	if err != nil {
		return err
	}
	*t = BidAsk{
		Price:  price,
		Amount: amount,
	}
	return err
}

type MarketDepth struct {
	GeneralResponse
	Data struct {
		Last   string   `json:"last"` // Latest transaction price
		Update Time     `json:"time"` // Depth update time
		Asks   []BidAsk `json:"asks"` // Ask depth
		Bids   []BidAsk `json:"bids"` // Bid depth
	}
}

func (r *MarketDepth) Parse(raw_response *http.Response) (*MarketDepth, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	defer raw_response.Body.Close()
	return r, err
}

// SingleMarketStatistics
type SingleMarketStatistics struct {
	GeneralResponse
	Data TickerData
}

// AllMarketStatistics
type AllMarketStatistics struct {
	GeneralResponse
	Data TickerData
}

// CurrencyRate
type CurrencyRate struct {
	GeneralResponse
	Data TickerData
}

// AssetAllocation
type AssetAllocation struct {
	GeneralResponse
	Data TickerData
}
type Float64 float64

func (f *Float64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	var f64 string
	if err := json.Unmarshal(data, &f64); err != nil {
		return err
	}
	value, err := strconv.ParseFloat(f64, 64)
	if err != nil {
		return err
	}
	*f = Float64(value)
	return nil
}

/// TODO: Convert strings to float64
type TickerData struct {
	ServerTime Time    `json:"date"`        // server time when returning
	Bid1Price  float64 `json:"buy"`         // First bid price
	Bid1Amount float64 `json:"buy_amount"`  // First bid amount
	Ask1Price  float64 `json:"sell"`        // First ask price
	Ask1Amount float64 `json:"sell_amount"` // First ask amount
	Open       float64 `json:"open"`        // 24H opening price
	Close      float64 `json:"last"`        // Latest transaction price
	High       float64 `json:"high"`        // 24H highest price
	Low        float64 `json:"low"`         // 24H lowest price
	Volume     float64 `json:"vol"`         // 24H volume
}

func (t *TickerData) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	var realTicker struct {
		ServerTime Time `json:"date"` // server time when returning
		Ticker     struct {
			Bid1Price  Float64 `json:"buy"`         // First bid price
			Bid1Amount Float64 `json:"buy_amount"`  // First bid amount
			Ask1Price  Float64 `json:"sell"`        // First ask price
			Ask1Amount Float64 `json:"sell_amount"` // First ask amount
			Open       Float64 `json:"open"`        // 24H opening price
			Close      Float64 `json:"last"`        // Latest transaction price
			High       Float64 `json:"high"`        // 24H highest price
			Low        Float64 `json:"low"`         // 24H lowest price
			Volume     Float64 `json:"vol"`         // 24H volume
		} `json:"ticker"` // Ticker data
	}
	if err := json.Unmarshal(data, &realTicker); err != nil {
		return err
	}

	*t = TickerData{
		ServerTime: realTicker.ServerTime,
		Bid1Price:  float64(realTicker.Ticker.Bid1Price),
		Bid1Amount: float64(realTicker.Ticker.Bid1Amount),
		Ask1Price:  float64(realTicker.Ticker.Ask1Price),
		Ask1Amount: float64(realTicker.Ticker.Ask1Amount),
		Open:       float64(realTicker.Ticker.Open),
		Close:      float64(realTicker.Ticker.Close),
		High:       float64(realTicker.Ticker.High),
		Low:        float64(realTicker.Ticker.Low),
		Volume:     float64(realTicker.Ticker.Volume),
	}

	return nil
}

func (r *SingleMarketStatistics) Parse(raw_response *http.Response) (*SingleMarketStatistics, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	defer raw_response.Body.Close()
	return r, nil
}
