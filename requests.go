package gocoinex

import (
	"fmt"
	"net/http"
	"net/url"
)

func requestMaker(endpoint string, params Map) (*http.Request, error) {
	route, err := url.Parse(FuturesBaseURL + endpoint)
	query := route.Query()
	if err != nil {
		return nil, err
	}
	if params != nil {
		for k, v := range params {
			query.Set(k, fmt.Sprintf("%v", v))
		}
		route.RawQuery = query.Encode()
	}
	request, err := http.NewRequest("GET", route.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", UserAgent)
	return request, err
}
