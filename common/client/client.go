package client

import (
	"errors"
	"io"
	"net/http"
)

func DoRequest(client http.Client, url string) ([]byte, error) {
	resp.err := client.Get(url)
	if resp.err != nil {
		return nil, resp.err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code: " + resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
