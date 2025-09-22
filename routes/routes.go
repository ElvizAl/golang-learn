package routes

import (
	"userfc/cmd/user/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler handler.UserHandler) {
	// Public Api
	router.GET("/ping", userHandler.Ping)
}
