package station

import (
	"net/http"

	"github.com/MaRVi401/belajarAPI-GolangWith-Gin.git/common/response"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()
	station := router.Group("/stations")
	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
	station.GET("/:id", func(c *gin.Context) {
		CheckSchedulesByStations(c, stationService)
	})
}
func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Success Get All Station",
			Data:    datas,
		},
	)
}

func CheckSchedulesByStations(c *gin.Context, service Service) {
	id := c.Param("id")

	datas, err := service.CheckSchedulesByStations(id)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Success Check Schedules By Stations",
			Data:    datas,
		},
	)
}
