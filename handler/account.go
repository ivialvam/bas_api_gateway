package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler --> folder
// acccount go --> file

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	BalanceAccount(*gin.Context)
}

type acccountImplement struct{}

func BalanceAccount() AccountInterface {
	return &acccountImplement{}
}

func (a *acccountImplement) GetAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()
	name := queryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"Username": "Get account successfully :)",
		"Password": name,
	})
}

func NewAccount() AccountInterface {
	return &acccountImplement{}
}

type BodyPayLoadAccount struct {
	AccountID string
	Name      string
	Address   string
}

func (a *acccountImplement) CreateAccount(g *gin.Context) {
	bodyPayload := BodyPayLoadAccount{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this rest api for later :)",
		"data":    bodyPayload,
	})
}

func (a *acccountImplement) UpdateAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()
	name := queryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "Update account successfully :)",
		"data":    map[string]string{"name": name},
	})
}

func (a *acccountImplement) DeleteAccount(g *gin.Context) {
	id := g.Param("id")

	g.JSON(http.StatusOK, gin.H{
		"message": "Delete account successfully",
		"data":    map[string]string{"name": id},
	})
}

func (a *acccountImplement) BalanceAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()
	balance := queryParam.Get("balance")

	g.JSON(http.StatusOK, gin.H{
		"message": "Account balance successfully",
		"data":    map[string]string{"name": balance},
	})
}
