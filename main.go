package main

import (
	"log"

	"github.com/sharanrohit7/authService/config"
	"github.com/sharanrohit7/authService/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	config.ConnectDB()

	routes.AuthRoutes(r)
	//  routes.TokenRoutes(r)

	r.Run(":8080")
}
