package main

import (
	"net/http"
	"sa-webapp/handler"
	"sa-webapp/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/sentiment", handler.SentimentHandler)

	r.POST("/analyse/sentiment", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"sentence": "a",
			"polarity": 1.2,
		})
	})
	r.Run(":8080")
}
