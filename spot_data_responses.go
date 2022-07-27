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
