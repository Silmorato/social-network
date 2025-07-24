package main

import (
	"log"
	"net/http"
	"twitter-clone/cmd/api"
	"twitter-clone/internal/adapters/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("could not initialize db: %v", err)
	}
	router := api.BuildApp(db)
	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
