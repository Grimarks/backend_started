package routes

import (
	"backend_started/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Product routes
	r.GET("/products", controllers.GetAllProducts)
	r.POST("/products", controllers.CreateProduct)

	// User routes
	r.GET("/users", controllers.GetAllUsers)
	r.POST("/users", controllers.CreateUser)

	// Payment routes
	r.GET("/payments", controllers.GetAllPayments)
	r.POST("/payments", controllers.CreatePayment)

	// Transaction routes
	r.GET("/transactions", controllers.GetAllTransactions)
	r.POST("/transactions", controllers.CreateTransaction)
}
