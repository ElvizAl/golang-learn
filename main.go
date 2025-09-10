package main

import (
	"userfc/cmd/user/handler"
	"userfc/config"
	"userfc/infrastructure/log"
	"userfc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Memuat konfigurasi aplikasi dari file config.yml
	cfg := config.LoadConfig()

	log.SetupLogger()

	userHandler := handler.NewUserHandler()

	port := cfg.App.Port // Mengambil nilai port dari konfigurasi aplikasi
	router := gin.Default()
	routes.SetupRoutes(router, userHandler)
	router.Run(":" + port) // Menjalankan server pada port yang ditentukan

	log.Logger.Printf("Server berjalan pada port %s", port)

}
