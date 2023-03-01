package gocoinex

import (
	"encoding/json"
	"net/http"
)

type AssetQueryItem struct {
	Avaiable     string
	Frozen       string
	Transfer     string
	TotalBalance string `json:"balance_total"`
	Margin       string
	UnrealProfit string `json:"profit_unreal"`
}

type AssetQuery struct {
	Data map[string]AssetQueryItem
}

func (r *AssetQuery) Parse(raw_response *http.Response) (*AssetQuery, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}
