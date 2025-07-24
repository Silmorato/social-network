package main

import (
	"log"
	"net/http"
	"twitter-clone/cmd/api"
)

func main() {
	router := api.BuildApp() //
	log.Println("🚀 Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
