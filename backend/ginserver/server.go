package ginserver

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func authorization(c *gin.Context) {
	// before request
	log.Println("auth...")

	c.Next()

	// after request
	log.Println("after handling...")
}

func Start() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.Default())

	// Authorization middleware
	r.Use(authorization)
	// /auth/*
	auth := r.Group("/auth")

	// /auth/ping
	auth.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
