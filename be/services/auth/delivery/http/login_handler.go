package http

import (
	"be/services/auth/model/request"
	"be/services/auth/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)



func LoginHandler(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request payload: " + err.Error(),
        })
        return
    }
	res, token, err := usecase.Login(req)
	if err != nil {c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	return}

	c.JSON(http.StatusOK, gin.H{
        "data":  res,
        "token": token,
    })
}
