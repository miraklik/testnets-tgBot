package controllers

import (
	"context"
	"errors"
	"net/http"
	"tg-bot-server/internal/models"
	"tg-bot-server/internal/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestnetControllers struct {
	services service.TestService
}

func NewTestnetControllers(services *service.TestService) *TestnetControllers {
	return &TestnetControllers{services: *services}
}

func (tc *TestnetControllers) CreateTestnet(c *gin.Context) {
	var testnet models.Testnet
	if err := c.ShouldBindJSON(&testnet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.services.CreateTestnet(c.Request.Context(), testnet.Name, testnet.Description, testnet.Link, testnet.DataAirdrop); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Testnet created successfully"})
}

func (tc *TestnetControllers) CreateTestnets(c *gin.Context) {
	var testnets []models.Testnet
	if err := c.ShouldBindJSON(&testnets); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	names := make([]string, 0, len(testnets))
	descriptions := make([]string, 0, len(testnets))
	links := make([]string, 0, len(testnets))
	airdropdates := make([]string, 0, len(testnets))

	for _, t := range testnets {
		names = append(names, t.Name)
		descriptions = append(descriptions, t.Description)
		links = append(links, t.Link)
		airdropdates = append(airdropdates, t.DataAirdrop)
	}

	newTestnets, err := tc.services.CreateTestnets(c.Request.Context(), names, descriptions, links, airdropdates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Testnets created successfully", "testnets": newTestnets})
}

func (tc *TestnetControllers) GetTestnets(c *gin.Context) {
	response, err := tc.services.GetTestnets(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Testnets fetched successfully",
		"testnets": response,
	})
}

func (tc *TestnetControllers) UpdateTestnet(c *gin.Context) {
	var testnet models.UpdateTestnet
	if err := c.ShouldBindJSON(&testnet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.services.UpdateTestnet(context.Background(), testnet.TestnetName, testnet.NewName, testnet.NewDescription, testnet.NewLink, testnet.NewDataAirdrop); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Testnet not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Testnet updated successfully"})
}

func (tc *TestnetControllers) DeleteTestnet(c *gin.Context) {
	var testnet models.Testnet
	if err := c.ShouldBindJSON(&testnet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.services.DeleteTestnet(c.Request.Context(), testnet.Name); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusConflict, gin.H{"error": "Testnet do not exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Testnet deleted successfully"})
}
