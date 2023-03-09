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
