package main

import (
	"backend_started/config"
	"backend_started/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	config.ConnectDatabase()

	routes.SetupRoutes(r)

	port := "8080"
	if p := config.GetEnv("PORT", "8080"); p != "" {
		port = p
	}
	r.Run(":" + port)
}
