package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/dkkundu", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server Running with Docker",
		})
	})

	router.Run(":80")
}
