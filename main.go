package main

import (
	"github.com/MaRVi401/belajarAPI-GolangWith-Gin.git/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	router.Run(":8080")
}
func initiateRouter() {
	var (
		router = gin.Default()
		api    = router.Group("/api/v1")
	)
	station.Initiate(api)
	router.Run(":8080")
}
