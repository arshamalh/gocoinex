package gocoinex

func newFutureDataClient(client Client) *FutureDataClient {
	return &FutureDataClient{
		client: client,
	}
}

/*
func TestGetMarketList(t *testing.T) {
	mockedClient := new(mockedClient)
	client := newFutureDataClient(mockedClient)
	body := io.NopCloser(strings.NewReader("Hello world"))
	httpResponse := &http.Response{Body: body, Status: "200 OK", StatusCode: 200}
	mockedClient.On("Do").Return(httpResponse, nil)
	client.GetMarketList()
}
*/
