package gocoinex

import (
	"encoding/json"
	"net/http"
)

const (
	BaseURL   = "https://api.coinex.com/v1/"
	UserAgent = "gocoinex"
)

type GeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type MaintenanceInformation struct {
	GeneralResponse
	Data struct {
		StartTime int    `json:"start_time"` // Integer Maintenance start time
		EndTime   int    `json:"end_time"`   // Integer Maintenance end time
		URL       string `json:"url"`        // URL for maintenance announcement
	}
}

type PartialMaintenanceInfo struct {
	GeneralResponse
	Data struct {
		StartedAt int      `json:"started_at"` // Integer Maintenance start time
		EndedAt   int      `json:"ended_at"`   // Integer Maintenance end time
		Scope     []string `json:"scope"`      // Maintenance scope: PERPETUAL, SPOT.
	}
}

// Obtain site-wide maintenance information,
// which is applicable to scenarios where all functions of the site are out of service
func GetMaintenanceInformation() (*MaintenanceInformation, error) {
	raw_response, err := http.DefaultClient.Get(BaseURL + "common/maintain/info")
	if err != nil {
		return nil, err
	}
	data := &MaintenanceInformation{}
	json.NewDecoder(raw_response.Body).Decode(data)
	defer raw_response.Body.Close()
	return data, nil
}

// Acquire partial maintenance information,
// applicable to the scenarios where some functions are out of service
func GetPartialMaintenanceInfo() (*PartialMaintenanceInfo, error) {
	raw_response, err := http.DefaultClient.Get(BaseURL + "common/temp-maintain/info")
	if err != nil {
		return nil, err
	}
	data := &PartialMaintenanceInfo{}
	json.NewDecoder(raw_response.Body).Decode(data)
	defer raw_response.Body.Close()
	return data, nil
}
