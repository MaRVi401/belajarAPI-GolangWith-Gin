package station

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type Service interface {
	GetAllStation() (Response []StationResponse, err error)
	CheckSchedulesByStations(id string) (Response []ScheduleResponse, err error)
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

func (s *service) CheckSchedulesByStations(id string) (Response []ScheduleResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/jadwal/"

	bytesResponse, err := DoRequest(s.client, url)
	if err != nil {
		return
	}

	var schedule []Schedule
	err = json.Unmarshal(bytesResponse, &schedule)
	if err != nil {
		return
	}

	var scheduleSelected Schedule

	for _, item := range schedule {
		if item.StationId == id {
			scheduleSelected = item
			break
		}
	}

	if scheduleSelected.StationId == "" {
		err = errors.New("Station not found")
		return
	}

	Response, err = ConvertToScheduleResponse(scheduleSelected)
	if err != nil {
		return
	}
	return
}

func ConvertToScheduleResponse(schedule Schedule) (Response []ScheduleResponse, err error) {
	var (
		LebakBulusTripName = "Stasiun Lebak Bulus Grab"
		BundaranHITripName = "Stasiun Bundaran HI Bank DKI"
	)
	scheduleLebakBulus := schedule.ScheduleLebakBulus
	scheduleBundaranHI := schedule.ScheduleBundaranHI

	scheduleLebakBulusParsed, err := ConvertScheduleToTimeFormat(scheduleLebakBulus)
	if err != nil {
		return
	}

	scheduleBundaranHIParsed, err := ConvertScheduleToTimeFormat(scheduleBundaranHI)
	if err != nil {
		return
	}

	//Convert to response
	for _, item := range scheduleLebakBulusParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			Response = append(Response, ScheduleResponse{
				StationName: LebakBulusTripName,
				Time:        item.Format("15:04"),
			})
		}
	}

	for _, item := range scheduleBundaranHIParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			Response = append(Response, ScheduleResponse{
				StationName: BundaranHITripName,
				Time:        item.Format("15:04"),
			})
		}
	}
	return
}

func ConvertScheduleToTimeFormat(schedule string) (respose []time.Time, err error) {
	var parsedTime time.Time

	Schedule := strings.Split(schedule, ",")

	for _, item := range Schedule {
		trimmedTime := strings.TrimSpace(item)
		if trimmedTime == "" {
			continue
		}
		parsedTime, err = time.Parse("15:04", trimmedTime)
		if err != nil {
			err = errors.New("Invalid time format" + trimmedTime)
			return
		}
		respose = append(respose, parsedTime)
	}

	return
}
