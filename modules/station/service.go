package station

import (
	"encoding/json"
	"net/http"
	"time"
)

type Service interface {
	GetAllStation() (Response []StationResponse, err error)
}
type service struct {
	client http.Client
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

	bytesResponse, err := DoRequest(*s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(bytesResponse, &stations)

	for _, item := range stations {
		Response = append(Response, StationResponse{
			Id:   item.id,
			Name: item.name,
		})
	}

	return
}
