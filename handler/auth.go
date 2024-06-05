package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	AuthLogin(*gin.Context)
}

type AuthImplement struct{}

func Login() AuthInterface {
	return &AuthImplement{}
}

type BodyPayLoadAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthImplement) AuthLogin(c *gin.Context) {
	var bodyPayload BodyPayLoadAuth

	err := c.BindJSON(&bodyPayload)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if bodyPayload.Username == "admin" && bodyPayload.Password == "admin123" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Account retrieved successfully",
			"data":    bodyPayload,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Invalid username or password",
		})
	}
}
