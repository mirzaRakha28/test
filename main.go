package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Struct for JSON response
type Response struct {
	Message string `json:"message"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", helloWorld)

	log.Println("Server starting on port 6060...")
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
