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
// @Success 200
// @Router /auth/login [post]
func Login(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

// Refresh godoc
// @Summary Logging to an account to generate an api token
// @Description Returns users infos on successful creation.
// @Tags auth
// @Produce json
// @Success 200
// @Router /auth/refresh [post]
func Refresh(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

// Revoke godoc
// @Summary Revoke an existing token
// @Description Revoke the current token.
// @Tags auth
// @Produce json
// @Success 200
// @Router /auth/revoke [post]
func Revoke(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
