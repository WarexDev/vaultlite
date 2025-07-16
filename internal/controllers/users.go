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

// CreateUserSecret godoc
// @Summary Create a new secret into user's private vault.
// @Description Create a new credential.
// @Tags Users/Secrets
// @Produce json
// @Success 200
// @Router /users/me/secrets [post]
func CreateUserSecret(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

// ListUserSecret godoc
// @Summary Get a list of secret ids from user's private vault.
// @Description Returns the list of available credentials.
// @Tags Users/Secrets
// @Produce json
// @Success 200
// @Router /users/me/secrets [get]
func ListUserSecret(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

// GetUserSecret godoc
// @Summary Get a secret by key from user's private vault.
// @Description Returns the value for a credential.
// @Tags Users/Secrets
// @Produce json
// @Success 200
// @Router /users/me/secrets/:key [get]
func GetUserSecret(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

// UpdateUserSecret godoc
// @Summary Update an existing secret from user's private vault.
// @Description Update the value of an existing credential.
// @Tags Users/Secrets
// @Produce json
// @Success 200
// @Router /users/me/secrets/:key [put]
func UpdateUserSecret(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

//	DeleteUserSecret godoc
//
// @Summary Delete an existing secret from user's private vault.
// @Description Delete the value of an existing credential.
// @Tags Users/Secrets
// @Produce json
// @Success 200
// @Router /users/me/secrets/:key [delete]
func DeleteUserSecret(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
