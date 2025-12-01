package client

import (
	"io"
	"net/http"
)

func DoRequest(url string) ([]byte, error) {
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
