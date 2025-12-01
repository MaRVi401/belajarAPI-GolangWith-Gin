package main

import (
	"github.com/MaRVi401/belajarAPI-GolangWith-Gin.git/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Panggil initiateRouter untuk mendapatkan router yang sudah memiliki semua rute
	router := initiateRouter()

	// 2. Jalankan router HANYA SEKALI
	router.Run(":8080")
}

// initiateRouter sekarang mengembalikan *gin.Engine (router)
func initiateRouter() *gin.Engine {
	// Buat router utama
	router := gin.Default()

	// Definisikan grup API /api/v1
	api := router.Group("/api/v1")

	// Inisialisasi semua modul di bawah grup /api/v1
	station.Initiate(api)

	// Kembalikan router yang sudah diinisialisasi
	return router
}
