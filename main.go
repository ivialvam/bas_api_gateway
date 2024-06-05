package main

import (
	"api_gateway/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.GET("/balance", handler.NewAccount().BalanceAccount)

	authRoute := r.Group("/auth")
	authRoute.POST("/post", handler.Login().AuthLogin)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/post", handler.Transaction().CreateTransaction)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
