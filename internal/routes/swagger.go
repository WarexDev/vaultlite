package routes

import (
	_ "github.com/WarexDev/vaultlite/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(apiGroup *gin.RouterGroup) {
	// swagger documentation
	apiGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
