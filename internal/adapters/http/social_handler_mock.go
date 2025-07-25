package http

import "net/http"

type MockSocialHandler struct{}

func (m *MockSocialHandler) CreateTweet(w http.ResponseWriter, r *http.Request) {}
func (m *MockSocialHandler) GetTimeline(w http.ResponseWriter, r *http.Request) {}
func (m *MockSocialHandler) FollowUser(w http.ResponseWriter, r *http.Request)  {}
