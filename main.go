package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Struct for JSON response
type Response struct {
	Message      string `json:"message"`
	RandomNumber int    `json:"random_number"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random number between 1 and 100
	randomNumber := rand.Intn(100) + 1

	// Create response with message and random number
	response := Response{
		Message:      "Hello, World!",
		RandomNumber: randomNumber,
	}

	// Send JSON response
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", helloWorld)

	log.Println("Server starting on port 6060...")
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
