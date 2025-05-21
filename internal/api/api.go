// Package api provides the initialization and routing logic for VaultLite's HTTP API.
//
// It sets up the Gin engine and registers all API route groups, including:
// - Swagger documentation endpoints
// - Health check endpoints
// - Authentication and secrets management endpoints
//
// This package acts as the entry point for configuring and exposing the REST API server.
package api

import (
	"github.com/WarexDev/vaultlite/internal/routes"
	"github.com/gin-gonic/gin"
)

// SetupRouter Create gin router & Register api routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Create api group
	apiGroup := r.Group("/api")
	// Register routes
	routes.RegisterSwagger(apiGroup)
	routes.RegisterHealth(apiGroup)
	routes.RegisterAuth(apiGroup)
	return r
}
