package httpclient

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendGetRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer GET: %v", err)
	}
	return resp, nil
}

func SendPostRequest(url string, payload string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar request POST: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
