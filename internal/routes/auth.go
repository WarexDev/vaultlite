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
// - POST /login: to authenticate and obtain an API token.
// - POST /register: to create a new user account.
func RegisterAuth(apiGroup *gin.RouterGroup) {
	// register auth service
	apiGroup.POST("/login", controllers.Login)
	apiGroup.POST("/register", controllers.Register)
}
