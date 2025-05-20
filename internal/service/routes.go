package service

import (
	_ "github.com/WarexDev/vaultlite/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {

	// register api group
	apiGroup := router.Group("/api")

	// swagger documentation
	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// health services
	apiGroup.GET("/ping", ping)
	apiGroup.GET("/health", healthCheck)

}
