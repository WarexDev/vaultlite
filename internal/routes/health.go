package routes

import (
	controllers "github.com/WarexDev/vaultlite/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterHealth(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/ping", controllers.Ping)
	apiGroup.GET("/health", controllers.HealthCheck)
}
