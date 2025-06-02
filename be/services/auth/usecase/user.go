package usecase

import (
	"be/pkg/db"
	"be/services/auth/model/entity"
	"be/services/auth/model/request"
	"be/services/auth/model/responce"

	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// GenerateToken tạo JWT token cho user
func GenerateToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"name":    user.Name,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

//tạo token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))	
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Login
func Login(req request.LoginRequest) (responce.LoginResponse, string, error) {
	var user entity.User

	// check xem user đã tồn tại trong database chưa
	err := db.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Tạo user mới nếu không tìm thấy
			user = entity.User{
				Name:    req.Name,
				Email:   req.Email,
				Picture: req.Picture,
			}
			if err := db.DB.Create(&user).Error; err != nil {
				return responce.LoginResponse{}, "", errors.New("failed to create user: " + err.Error())
			}
		} else {
			return responce.LoginResponse{}, "", errors.New("database error: " + err.Error())
		}
	}

	// Tạo JWT token
	token, err := GenerateToken(user)
	if err != nil {
		return responce.LoginResponse{}, "", errors.New("failed to generate token: " + err.Error())
	}

	// Tạo response
	res := responce.LoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return res, token, nil
}