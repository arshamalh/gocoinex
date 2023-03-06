package gocoinex

import "net/http"

type FutureDataClient struct {
	client Client
}

func NewFutureDataClient() *FutureDataClient {
	return &FutureDataClient{client: &http.Client{}}
}
