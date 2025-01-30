package main

import (
	"demo/src/products/infraestructure/dependencies"
	"demo/src/users/infraestructure/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
    "log"
	
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	r := gin.Default()
	dependencies.InitProducts(r)
	config.InitUsers(r)
	if err := r.Run(); err != nil {
		panic(err)
	}
}