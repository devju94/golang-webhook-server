package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type webhookBody struct {
	CODE	string	`json:"code"`
	MESSAGE	string	`json:"message"`
}

func main() {
	// get env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server_host := os.Getenv("SERVER_HOST")
	server_port := os.Getenv("SERVER_PORT")
	gin_mode := os.Getenv("GIN_MODE")

	// run server
	gin.SetMode(gin_mode)
	router := gin.Default()
	router.POST("/webhook", postWebhook)
	router.Run(server_host + ":" + server_port)
}

func postWebhook(c *gin.Context) {
	var data webhookBody

	if err := c.BindJSON(&data); err != nil {
		return
	}

	/**
	* do something
	*/
}