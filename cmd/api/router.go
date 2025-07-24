package api

import (
	"github.com/gorilla/mux"
	"net/http"
	httpAdap "twitter-clone/internal/adapters/http"
)

func NewRouter(socialHandler *httpAdap.SocialHandler) http.Handler {
	router := mux.NewRouter()

	// POST /tweets → create tweet
	router.HandleFunc("/tweets", socialHandler.CreateTweet).Methods("POST")

	// GET /timeline → get timeline
	router.HandleFunc("/timeline", socialHandler.GetTimeline).Methods("GET")

	// POST /follow → follow user
	router.HandleFunc("/follow", socialHandler.FollowUser).Methods("POST")

	return router
}
