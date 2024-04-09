package main

import (
	"log"
	"os"

	"github.com/dragaoazuljr/ice-breaker-go/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	//get name from args
	name := os.Args[1]

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.IceBreaker(name)
}
