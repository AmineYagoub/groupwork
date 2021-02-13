package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api pong",
		})
	})

	r.GET("/api/auth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth api pong",
		})
	})

	r.GET("/auth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "auth pong",
		})
	})

	log.Print("Hi")
	r.Run() // listen and serve on 0.0.0.0:8080
}
