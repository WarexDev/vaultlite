package api

import (
	"github.com/WarexDev/vaultlite/internal/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Create api group
	apiGroup := r.Group("/api")
	// Register routes
	routes.RegisterSwagger(apiGroup)
	routes.RegisterHealth(apiGroup)
	return r
}
