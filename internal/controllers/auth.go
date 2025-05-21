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

// Login godoc
// @Summary Logging to an account to generate an api token
// @Description Returns users infos on successful creation.
// @Tags auth
// @Produce json
// @Success 200 {object} map[string]string "token"
// @Router /api/login [post]
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": "hello from controller"})
}

// Register godoc
// @Summary Register a new account to the service
// @Description Returns users infos on successful creation.
// @Tags auth
// @Produce json
// @Success 200 {object} map[string]string "OK status"
// @Router /api/register [post]
func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello from controller"})
}
