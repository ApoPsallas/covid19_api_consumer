package mocks

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type MockHttpClient struct {
	ClientResponse interface{}
	ApiKey         string
}

func (client *MockHttpClient) Do(req *http.Request) (*http.Response, error) {

	reqBody, _ := json.Marshal(client.ClientResponse)
	var err error
	key := req.Header.Get("x-rapidapi-key")
	if key != client.ApiKey {
		reqBody, _ = json.Marshal("")
		err = errors.New("Wrong API key")
	}

	return &http.Response{Body: ioutil.NopCloser(bytes.NewReader(reqBody))}, err
}
