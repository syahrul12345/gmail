package main

import (
	"fmt"
	"gmailclient/email"
	"log"
	"os"
	"gmailclient/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func setupRouter() *gin.Engine {
	err := godotenv.Load("../.env")
  if err != nil {
    log.Fatal("Error loading .env file")
	}

	app := gin.Default()
	app.Use(auth.AuthMiddleware())

	verification := app.Group("/verification")
	email.Register(verification)

	return app
}

func main() {
	app := setupRouter()
	app.Run(fmt.Sprintf("%s:%s","localhost",os.Getenv("PORT")))
}