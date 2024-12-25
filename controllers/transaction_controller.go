package controllers

import (
	"backend_started/config"
	"backend_started/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	var transactions []models.Transaction
	if err := config.DB.Preload("Product").Preload("Buyer").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"debug": "Input JSON received", "transaction": transaction})

	var product models.Product
	if err := config.DB.First(&product, transaction.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"debug": "Product Found", "product": product})

	if product.Status == "sold" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product already sold"})
		return
	}

	transaction.TotalPrice = float64(transaction.Quantity) * product.Price

	if err := config.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.Status = "sold"
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product status"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}
