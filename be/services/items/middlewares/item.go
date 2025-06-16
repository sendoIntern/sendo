package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"be/pkg/db"
	"be/services/auth/model/request"
	"be/services/items/model/entity"
	"errors"

	jwt "be/services/items/utils"
)

// function kiểm tra ID trước khi gọi handlers
func ItemIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("itemId")
		
		// Kiểm tra xem item có tồn tại không
		var item entity.Item
		result := db.DB.First(&item, "id = ?", id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + result.Error.Error()})
			}
			c.Abort()
			return
		}
		c.Next()
	}
}

// check input của item có hợp lệ không
func ValidateItemFields() gin.HandlerFunc {
	return func(c *gin.Context) {
			name := c.PostForm("name")
			price := c.PostForm("price")
			quantity := c.PostForm("quantity")
			description := c.PostForm("description")

			if name == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
				return
			}
			if price == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Price is required"})
				return
			}
			if quantity == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Quantity is required"})
				return
			}
			if description == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
				return
			}

		c.Next()
	}
}


func RequireExcelFileMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "file is required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

//validate access token
func ValidateAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Loại bỏ tiền tố "Bearer "
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}
		token := splitToken[1]

		// Xác thực token
		_, err := jwt.VerifyToken(token, []byte(os.Getenv("JWT_SECRET_ACCESSTOKEN")))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token: " + err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
	


// validate refresh token lấy từ body request

func ValidateRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.RefreshTokenRequest
		// fmt.Println("ValidateRefreshToken middleware called"+ req.RefreshToken)
		if err := c.ShouldBindJSON(&req); err != nil || req.RefreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is required"})
			c.Abort()
			return
		}

		decoded, err := jwt.VerifyToken(req.RefreshToken, []byte(os.Getenv("JWT_SECRET_REFRESHTOKEN")))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("refreshToken", req) // Lưu refresh token vào context nếu cần thiết
		c.Set("claims", decoded) 
		c.Next()
	}
}


