// Package routes defines the route registration logic for VaultLite's API.
//
// It provides functions to group and register API endpoints (e.g., authentication,
// health checks, and documentation) to specific HTTP route paths using Gin's router groups.
//
// This package ensures that all route mappings remain modular, organized, and easy to maintain.
package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterSwagger registers the Swagger documentation endpoint to the given API route group.
//
// It defines the following route:
// - GET /swagger/*any: serves the interactive Swagger UI for API documentation.
func RegisterSwagger(apiGroup *gin.RouterGroup) {
	// swagger documentation
	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
