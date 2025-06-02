package handler

import (
	"be/db"
	"be/entity"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

var oauthConf *oauth2.Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
}

func GoogleLoginHandler(c *gin.Context) {
	url := oauthConf.AuthCodeURL("random-state-string")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func generateJWT(user entity.User) (string, error) {
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

func GoogleCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	client := oauthConf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user info"})
		return
	}

	// Chuyển map => struct
	user := entity.User{
		Name:    userInfo["name"].(string),
		Email:   userInfo["email"].(string),
		Picture: userInfo["picture"].(string),
	}

	// Check tồn tại user (theo Google ID hoặc email), nếu chưa có thì tạo
	var existing entity.User
	result := db.DB.Where("email = ?", user.Email).First(&existing)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			db.DB.Create(&user)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
	} else {
		user = existing // đã tồn tại => dùng lại
	}

	// Tạo JWT token
	jwtToken, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Redirect về frontend với token
	frontendURL := os.Getenv("FRONTEND_URL")
	c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/auth/callback?token="+jwtToken)
}

// "email": "anhviettran357@gmail.com",
// "id": "116196705478408426432",
// "name": "Việt Anh Trần",
// "picture": "https://lh3.googleusercontent.com/a/ACg8ocLnTONkVt0KTbptOvm9YVEvnBnrLZb7YVkAaQdrZ18TvhLeXovI=s96-c",
