package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize router using the default gin configuration
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8080")
}
