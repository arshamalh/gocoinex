package gocoinex

type Ping struct {
	Data string `json:"data"` // Pong

}

type SystemTime struct {
	Data int `json:"data"` // Server time, milliseconds

}
