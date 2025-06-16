package http

import (
	"be/pkg/db"
	"be/services/auth/model/entity"
	"be/services/auth/model/request"
	"be/services/auth/usecase"
	"fmt"
	"net/http"

	// jwt "be/services/items/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)



func LoginHandler(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request payload: " + err.Error(),
        })
        return
    }
	res, accessToken, refreshToken, err := usecase.Login(req)
	if err != nil {c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	return
}
	fmt.Println("Access Token:", accessToken)
	fmt.Println("Refresh Token:", refreshToken)
	c.JSON(http.StatusOK, gin.H{
        "data":  res,
        "accessToken": accessToken,
		"refreshToken": refreshToken,
    })
}

// func RefreshTokenHandler(c *gin.Context) {
// 	// lấy từ middleware
	
// 	var req request.RefreshTokenRequest
// 	fmt.Println("RefreshTokenHandler called" + req.RefreshToken)
// 	if err := c.ShouldBindJSON(&req); err != nil || req.RefreshToken == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid refreshToken"})
// 		return
// 	}

// 	claims, err := jwt.VerifyToken(req.RefreshToken, []byte(os.Getenv("JWT_SECRET_REFRESHTOKEN")))
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Lấy user từ DB
// 	email, ok := claims["email"].(string)
// 	if !ok {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
// 		return
// 	}

// 	var user entity.User
// 	err = db.DB.Where("email = ?", email).First(&user).Error
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
// 		return
// 	}

// 	// Tạo accessToken và refreshToken mới
// 	newAccessToken, err := usecase.SignAccessToken(user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign access token"})
// 		return
// 	}

// 	newRefreshToken, err := usecase.SignRefreshToken(user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign refresh token"})
// 		return
// 	}

// 	fmt.Println("New Access Token:", newAccessToken)
// 	fmt.Println("New Refresh Token:", newRefreshToken)
// 	c.JSON(http.StatusOK, gin.H{
// 		"accessToken":  newAccessToken,
// 		"refreshToken": newRefreshToken,
// 	})
// }


func RefreshTokenHandler(c *gin.Context) {
	// Lấy refreshToken và claims từ context (đã được middleware ValidateRefreshToken xử lý)
	_, exists := c.Get("refreshToken")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found in context"})
		return
	}

	claimsValue, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims not found in context"})
		return
	}
	claims, ok := claimsValue.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims format"})
		return
	}

	// Lấy email từ claims
	email, ok := claims["email"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
		return
	}

	// Tìm user trong DB
	var user entity.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Sinh accessToken và refreshToken mới
	newAccessToken, err := usecase.SignAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign access token"})
		return
	}

	newRefreshToken, err := usecase.SignRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	})
}




