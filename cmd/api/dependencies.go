package api

import (
	"net/http"
	"twitter-clone/internal/adapters/database"
	httpAdap "twitter-clone/internal/adapters/http"
	"twitter-clone/internal/services"
)

// BuildApp sets up all application dependencies and returns the configured HTTP handler.
func BuildApp() http.Handler {
	// PostgreSQL connection
	db := database.InitPostgresDB()

	// Repositories (Postgres version)
	tweetRepo := database.NewTweetRepository(db)
	followRepo := database.NewFollowRepository(db)
	userRepo := database.NewUserRepository(db)

	// Service (use case)
	tweetService := services.NewSocialService(tweetRepo, followRepo, userRepo)

	// Handler (HTTP)
	tweetHandler := httpAdap.NewSocialHandler(tweetService)

	// Router
	return NewRouter(tweetHandler)
}
