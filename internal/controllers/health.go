// Package controllers defines the HTTP handlers for VaultLite's REST API.
//
// It includes controller functions responsible for handling authentication,
// health checks, and other service-level endpoints.
//
// Each handler responds with appropriate JSON payloads and status codes
// and is documented using Swagger annotations for automated OpenAPI generation.
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Check the health status of the service
// @Description Returns a simple OK status to verify the service is up and running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string "OK status"
// @Router /api/health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Ping godoc
// @Summary Ping the service
// @Description Returns a simple pong status to verify the service is up and running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string "PONG status"
// @Router /api/ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "pong"})
}
