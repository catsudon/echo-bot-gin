package main

import (
	"echo-bot-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic",
		})
	})

	router.POST("/webhook", controllers.HandleWebhook())

	router.Run("localhost:6000")
}
