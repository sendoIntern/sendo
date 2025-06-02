package http

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	api := router.Group("/auth")
	{
		api.POST("/login", LoginHandler)
	}
}