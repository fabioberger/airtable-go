package airtable

import (
	"io/ioutil"
	"net/http"
)

type httpFetcher interface {
	Fetch(req *http.Request) (rawBody []byte, statusCode int, err error)
}

type realHTTPFetcher struct{}

func (realHTTPFetcher) Fetch(req *http.Request) (rawBody []byte, statusCode int, err error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, 0, err
	}
	defer resp.Body.Close()
	rawBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, 0, err
	}
	return rawBody, resp.StatusCode, nil
}

type fakeHTTPFetcher struct {
	statusCode  int
	rawResponse []byte
}

func (f fakeHTTPFetcher) Fetch(req *http.Request) (rawBody []byte, statusCode int, err error) {
	return f.rawResponse, f.statusCode, nil
}
