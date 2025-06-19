package controllers

import (
	"context"
	"net/http"
	"tg-bot-server/internal/models"
	"tg-bot-server/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthControllers struct {
	services *service.AuthService
}

func NewAuthControllers(services *service.AuthService) *AuthControllers {
	return &AuthControllers{services: services}
}

func (cont *AuthControllers) Login(c *gin.Context) {
	var Input models.User
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := cont.services.AuthenticateUser(context.Background(), Input.Username, Input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (cont *AuthControllers) ChangePassword(c *gin.Context) {
	var Input models.ChangePassword
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := cont.services.ChangePassword(context.Background(), Input.Username, Input.OldPassword, Input.NewPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"responce": response})
}
