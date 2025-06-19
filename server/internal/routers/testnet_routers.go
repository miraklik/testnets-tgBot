package routers

import (
	"tg-bot-server/internal/controllers"
	"tg-bot-server/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupTestnetRouters(router *gin.Engine, services *service.TestService) {
	controller := controllers.NewTestnetControllers(services)

	router.POST("/testnet", controller.CreateTestnet)
	router.POST("/testnets", controller.CreateTestnets)
	router.GET("/testnets", controller.GetTestnets)
	router.PUT("/testnet", controller.UpdateTestnet)
	router.DELETE("/testnet", controller.DeleteTestnet)
}
