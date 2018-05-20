package api

import (
	"github.com/gin-gonic/gin"
)

func StartApi () {
	r := gin.Default()

	r.GET("/getSimulation", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}