package main

import (
	"log"
	"os"
	"tg-bot-server/internal/config"
	"tg-bot-server/internal/routers"
	"tg-bot-server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set")
	}

	if err := config.InitMongoDB(mongoURI); err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}

	db := config.Client.Database("Testnet")
	authService := service.NewAuthService(db)
	testService := service.NewTestService(db)

	router := gin.Default()
	routers.SetupAuthRouter(router, authService)
	routers.SetupTestnetRouters(router, testService)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
