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
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	//tạo token
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func SignAccessToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"name":    user.Name,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_ACCESSTOKEN")))
	if err != nil {
		return "", errors.New("failed to sign access token: " + err.Error())
	}
	return tokenString, nil
}

func SignRefreshToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{		
		"user_id": user.ID.String(),
		"email":   user.Email,
		"name":    user.Name,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token expires in 7 days
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_REFRESHTOKEN")))
	if err != nil {
		return "", errors.New("failed to sign refresh token: " + err.Error())
	}
	return tokenString, nil
}

// Login
func Login(req request.LoginRequest) (responce.LoginResponse, string, string, error) {
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
				Role:    "user",
			}
			if err := db.DB.Create(&user).Error; err != nil {
				return responce.LoginResponse{}, "", "", errors.New("failed to create user: " + err.Error())
			}
		} else {
			return responce.LoginResponse{}, "", "", errors.New("database error: " + err.Error())
		}
	}

	// Tạo JWT token
	accessToken, err := SignAccessToken(user)
	if err != nil {
		return responce.LoginResponse{}, "", "", errors.New("failed to generate access_token: " + err.Error())
	}

	refreshToken, err := SignRefreshToken(user)
	if err != nil {
		return responce.LoginResponse{}, "", "", errors.New("failed to generate refresh_token: " + err.Error())
	}

	// Tạo response
	res := responce.LoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return res, accessToken, refreshToken, nil
}

