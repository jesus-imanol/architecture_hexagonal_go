package main

import (
	"demo/src/products/infraestructure/dependencies"
	"demo/src/users/infraestructure/config"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	dependencies.InitProducts(r)
	config.InitUsers(r)
	if err := r.Run(); err != nil {
		panic(err)
	}
}