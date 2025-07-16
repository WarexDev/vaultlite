// Package routes defines the route registration logic for VaultLite's API.
//
// It provides functions to group and register API endpoints (e.g., authentication,
// health checks, and documentation) to specific HTTP route paths using Gin's router groups.
//
// This package ensures that all route mappings remain modular, organized, and easy to maintain.
package routes

import (
	"github.com/WarexDev/vaultlite/internal/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterAuth registers the authentication-related endpoints to the given API route group.
//
// It defines the following routes:
// - POST /auth/login: to authenticate and obtain an API token.
// - POST /auth/revoke: to revoke an API token.
// - POST /auth/refresh: to renew an API token.
func RegisterAuth(apiGroup *gin.RouterGroup) {
	authGroup := apiGroup.Group("/auth")

	// user auth service
	authGroup.POST("/login", controllers.Login)
	authGroup.POST("/revoke", controllers.Revoke)
	authGroup.POST("/refresh", controllers.Refresh)
}
