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

// RegisterUsers registers the users endpoints to the given API route group.
//
// It defines the following routes:
// TODO : Add routes description
func RegisterUsers(apiGroup *gin.RouterGroup) {
	usersGroup := apiGroup.Group("/users/me")

	// user's secrets routes
	usersGroup.POST("/secrets", controllers.CreateUserSecret)
	usersGroup.GET("/secrets", controllers.ListUserSecret)
	usersGroup.GET("/secrets/:key", controllers.GetUserSecret)
	usersGroup.PUT("/secrets/:key", controllers.UpdateUserSecret)
	usersGroup.DELETE("/secrets/:key", controllers.DeleteUserSecret)
}
