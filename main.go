package main

import (
	"api_gateway/usecase"
	"fmt"
)

// Login struct untuk menyimpan informasi login
type Login struct {
	Username string
	Password string
}

func main() {
	login := usecase.TaskLogin()
	auth := login.Auth("ivialva", "123456")
	fmt.Println(auth)
}
