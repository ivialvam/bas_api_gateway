package main

import (
	"api_gateway/handler"
	"api_gateway/proto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	r := gin.Default()

	addrServiceTransactionOpt := client.WithAddress((":1809"))
	clientSrvTransaction := grpc.NewClient()

	srvTransaction := micro.NewService(
		micro.Client(clientSrvTransaction),
	)

	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.POST("/balance", handler.NewAccount().BalanceAccount)

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().AuthLogin)

	transactionRoute := r.Group("/transaction")
	// transactionRoute.POST("/transfer", handler.NewTransaction().CreateTransaction)
	transactionRoute.GET("/get", func(g *gin.Context) {
		clientResponse, err := proto.NewServiceTransactionService("service-transaction", srvTransaction.Client()).
			Login(context.Background(), &proto.LoginRequest{
				Username: "admin",
				Password: "admin123",
			}, addrServiceTransactionOpt)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		g.JSON(http.StatusOK, gin.H{
			"data": clientResponse,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
