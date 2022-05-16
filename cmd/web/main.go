package main

import (
	"log"
	"net/http"
	"ws/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting web server on port 8081")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8081")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
