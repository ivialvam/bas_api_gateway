package handler

import (
	"api_gateway/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	AuthLogin(*gin.Context)
}

type authImplement struct{}

func NewAuth() AuthInterface {
	return &authImplement{}
}

type bodyPayloadAuth struct {
	Username string
	Password string
}

func (pi *authImplement) AuthLogin(g *gin.Context) {
	BodyPayLoad := bodyPayloadAuth{}
	err := g.BindJSON(&BodyPayLoad)
	if err != nil {
		log.Fatal(err)
	}

	if usecase.NewLogin().Auth(BodyPayLoad.Username, BodyPayLoad.Password) {
		g.JSON(http.StatusOK, gin.H{
			"message": "Congratss!! Anda berhasil login",
			"data":    BodyPayLoad,
		})
	} else {
		g.JSON(http.StatusUnauthorized, gin.H{
			"message": "Maaf!! Anda gagal login :(",
			"data":    BodyPayLoad,
		})
	}
}
