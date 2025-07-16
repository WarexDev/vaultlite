// Package controllers defines the HTTP handlers for VaultLite's REST API.
//
// It includes controller functions responsible for handling authentication,
// health checks, and other service-level endpoints.
//
// Each handler responds with appropriate JSON payloads and status codes
// and is documented using Swagger annotations for automated OpenAPI generation.
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck godoc
// @Summary Check the health status of the service
// @Description Returns a simple OK status to verify the service is up and running
// @Tags health
// @Produce json
// @Success 200
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

// Ping godoc
// @Summary Ping the service
// @Description Returns a simple pong status to verify the service is up and running
// @Tags health
// @Produce json
// @Success 200
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
