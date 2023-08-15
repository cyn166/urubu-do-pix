package config

import (
	"github.com/cyn166/urubu-do-pix/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", controllers.IndexHandler)
	}
}
