package main

import (
	"log"

	"example.com/ginserver"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// httpserver.Start()
	ginserver.Start()
	// netserver.Start()
	// syscallserver.Start()
}
