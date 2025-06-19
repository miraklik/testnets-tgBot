package routers

import (
	"tg-bot-server/internal/controllers"
	"tg-bot-server/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupAuthRouter(router *gin.Engine, services *service.AuthService) {
	controller := controllers.NewAuthControllers(services)

	router.POST("/login", controller.Login)
	router.POST("/change-password", controller.ChangePassword)
}
