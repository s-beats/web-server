package ginserver

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func authorization(c *gin.Context) {
	// before request
	log.Println("auth...")

	c.Next()
}

func logging(c *gin.Context) {
	start := time.Now() // アクセスを受けた時間

	c.Next()

	latencyDuration := time.Now().Sub(start) // アクセスからのレスポンスまでの時間

	// ログ出力
	var logger *zerolog.Event
	baseLogger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()
	if len(c.Errors) > 0 {
		logger = baseLogger.Error().Err(c.Errors.Last())
	} else {
		logger = baseLogger.Info()
	}
	logger.
		Str("query", c.Request.URL.RawQuery).
		Int("status", c.Writer.Status()).
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.RawPath).
		Dur("latency", latencyDuration).
		Send()
}

func Start() {
	r := gin.Default()

	// Logging middleware
	r.Use(logging)

	// CORS middleware
	r.Use(cors.Default())

	// /auth/*
	// use Authorization middleware
	auth := r.Group("/auth", authorization)

	// /auth/ping
	auth.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
