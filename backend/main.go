package main

import (
	"log"

	"example.com/httpserver"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	httpserver.Start()
}
