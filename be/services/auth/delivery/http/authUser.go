package http

import (
	middleware "be/services/items/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	api := router.Group("/auth")
	{
		api.POST("/login", LoginHandler)
		api.POST("/refreshToken", middleware.ValidateRefreshToken(), RefreshTokenHandler)
		// api.DELETE("/logout", LogoutHandler)
	}
}