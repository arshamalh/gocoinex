package gocoinex

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type mockedClient struct {
	mock.Mock
}

func (mc *mockedClient) Do(req *http.Request) (*http.Response, error) {
	args := mc.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}
