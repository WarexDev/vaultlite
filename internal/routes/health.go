// Package routes defines the route registration logic for VaultLite's API.
//
// It provides functions to group and register API endpoints (e.g., authentication,
// health checks, and documentation) to specific HTTP route paths using Gin's router groups.
//
// This package ensures that all route mappings remain modular, organized, and easy to maintain.
package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/WarexDev/vaultlite/internal/controllers"
)

// RegisterHealth registers the health check endpoints to the given API route group.
//
// It defines the following routes:
// - GET /ping: simple connectivity check.
// - GET /health: full health check to verify the service is running properly.
func RegisterHealth(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/ping", controllers.Ping)
	apiGroup.GET("/health", controllers.HealthCheck)
}
