package gocoinex

import (
	"net/http"
)

type FutureAccountClient struct {
	client Client
}

func NewFutureAccountClient() *FutureAccountClient {
	return &FutureAccountClient{client: &http.Client{}}
}

func (c *FutureAccountClient) get(endpoint string, params Map) (*http.Response, error) {
	request, err := requestMaker(endpoint, params)
	if err != nil {
		return nil, err
	}
	return c.client.Do(request)
}

// https://viabtc.github.io/coinex_api_en_doc/futures/#docsfutures001_http016_asset_query
func (c *FutureAccountClient) GetAssetQuery() (*AssetQuery, error) {
	return nil, nil
}
