package station

import (
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

	//

	return
}
