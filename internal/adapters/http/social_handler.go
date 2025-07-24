package http

import (
	"encoding/json"
	"net/http"
	"twitter-clone/internal/dto"
	"twitter-clone/internal/errors"
	"twitter-clone/internal/ports"
	"twitter-clone/internal/services"
)

type SocialHandler struct {
	Service ports.SocialService
}

func NewSocialHandler(service *services.SocialService) *SocialHandler {
	return &SocialHandler{Service: service}
}

func (h *SocialHandler) CreateTweet(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTweetRequest
	if err := req.FromJSON(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tweet, err := h.Service.PublishTweet(req.UserID, req.Content)
	if err != nil {
		respondError(w, err.Status, err.Message)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweet)
}

func (h *SocialHandler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	req, err := dto.ParseGetTimelineRequest(r)
	if err != nil {
		errApi := errors.NewAPIError(http.StatusBadRequest, "Invalid request", err)
		respondError(w, errApi.Status, errApi.Message)
		return
	}

	tweets, apiErr := h.Service.GetTimeline(req.UserID)
	if apiErr != nil {
		respondError(w, apiErr.Status, apiErr.Message)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tweets)
}

func (h *SocialHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	var req dto.FollowRequest
	if err := req.FromJSON(r.Body); err != nil {
		errApi := errors.NewAPIError(http.StatusBadRequest, "Invalid request", err)
		respondError(w, errApi.Status, errApi.Message)
		return
	}

	msg, apiErr := h.Service.FollowUser(req.FollowerID, req.FollowingID)
	if apiErr != nil {
		respondError(w, apiErr.Status, apiErr.Message)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": *msg,
	})
}
