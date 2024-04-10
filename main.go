package main

import (
	"fmt"
	"log"
	"net/http"

	httpIceBreakers "github.com/dragaoazuljr/ice-breaker-go/internal/app/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	
	if err != nil {
		panic("Error loading .env file")
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /ice-breakers", httpIceBreakers.GetIceBreakers)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server is not running")
		log.Fatal(err)
	}
}
