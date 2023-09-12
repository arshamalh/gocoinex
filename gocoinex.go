package gocoinex

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

const (
	BaseURL        = "https://api.coinex.com/v1/"
	FuturesBaseURL = "https://api.coinex.com/perpetual/v1/"
	UserAgent      = "gocoinex"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type GeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	i, err := strconv.ParseInt(string(data), 10, 64)
	update := time.UnixMilli(i)
	*t = Time{update}
	return err
}

type Map map[string]interface{}

type MaintenanceInformation struct {
	GeneralResponse
	Data struct {
		StartTime Time   `json:"start_time"` // Integer Maintenance start time
		EndTime   Time   `json:"end_time"`   // Integer Maintenance end time
		URL       string `json:"url"`        // URL for maintenance announcement
	}
}

type PartialMaintenanceInfo struct {
	GeneralResponse
	Data struct {
		StartedAt Time     `json:"started_at"` // Integer Maintenance start time
		EndedAt   Time     `json:"ended_at"`   // Integer Maintenance end time
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
