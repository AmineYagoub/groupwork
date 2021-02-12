package app

import "github.com/gin-gonic/gin"

// Bootstrap just boot the app
func Bootstrap() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8020") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
