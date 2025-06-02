package handler

import (
	"be/db"
	"be/dto"
	"be/entity"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func generateToken(user entity.User) (string, error) {
	// Tạo JWT token với thời hạn 1 giờ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"name":    user.Name,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	// Ký token với secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func LoginHandler(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{
		Name:    req.Name,
		Email:   req.Email,
		Picture: req.Picture,
	}
	result := db.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			db.DB.Create(&user)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot connect to database"})
			return
		}
	}
	jwtToken, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	resp := dto.LoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Login success",
		"user":   resp,
		"token":  jwtToken,
	})
}
