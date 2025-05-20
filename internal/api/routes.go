package api

import (
	"github.com/WarexDev/vaultlite/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/health", service.HealthCheck)

	return r
}
