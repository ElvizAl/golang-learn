package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public Api
	router.Get("/ping", userHandler.Ping)
}
