package ginserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
