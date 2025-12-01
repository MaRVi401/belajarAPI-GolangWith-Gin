package station

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Service interface {
	GetAllStation() (Response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStation() (Response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	bytesResponse, err := DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(bytesResponse, &stations)
	if err != nil {
		return
	}

	for _, item := range stations {
		Response = append(Response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return
}

func DoRequest(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
