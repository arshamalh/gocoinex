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
