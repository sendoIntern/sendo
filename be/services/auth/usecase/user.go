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






// RefreshToken kiểm tra và tạo mới JWT token cho user
func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	},
	)
	if err != nil {
		return "", errors.New("invalid token: " + err.Error())
	}
	// Kiểm tra xem token có hợp lệ và chưa hết hạn
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			// So sánh thời gian hết hạn với thời gian hiện tại
			if time.Now().Unix() > int64(exp) {
				return "", errors.New("token has expired")
			}
		} else {
			return "", errors.New("invalid token claims")
		}

		// Tạo mới token với cùng thông tin claims
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": claims["user_id"],
			"email":   claims["email"],
			"name":    claims["name"],
			"role":    claims["role"],
			"exp":     time.Now().Add(time.Hour).Unix(),
		})

		newTokenString, err := newToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return "", errors.New("failed to sign new token: " + err.Error())
		}

		return newTokenString, nil
	}
	return "", errors.New("invalid token claims or token is not valid")	
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
	accessToken, err := GenerateToken(user)
	if err != nil {
		return responce.LoginResponse{}, "", "", errors.New("failed to generate token: " + err.Error())
	}

	refreshToken, err := GenerateToken(user)
	if err != nil {
		return responce.LoginResponse{}, "", "", errors.New("failed to generate refresh token: " + err.Error())
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

