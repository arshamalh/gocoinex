package gocoinex

import (
	"encoding/json"
	"net/http"
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

type AllMarketInfo struct {
	GeneralResponse
	Data map[string]struct {
		Name           string `json:"name"`            // Market name
		TakerFeeRate   string `json:"taker_fee_rate"`  // Taker rate
		MakerFeeRate   string `json:"maker_fee_rate"`  // Maker rate
		MinAmount      string `json:"min_amount"`      // Min transaction volume
		TradingName    string `json:"trading_name"`    // Trading currency
		TradingDecimal int    `json:"trading_decimal"` // Trading currency decimal
		PricingName    string `json:"pricing_name"`    // Pricing currency
		PricingDecimal int    `json:"pricing_decimal"` // Pricing currency decimal
	}
}

func (r *AllMarketInfo) Parse(raw_response *http.Response) (*AllMarketInfo, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}
