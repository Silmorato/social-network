package dto

import (
	"fmt"
	"net/http"
	"strings"
)

type GetTimelineRequest struct {
	UserID string
}

func ParseGetTimelineRequest(r *http.Request) (*GetTimelineRequest, error) {
	userID := r.Header.Get("X-User-ID")
	if strings.TrimSpace(userID) == "" {
		return nil, fmt.Errorf("X-User-ID header is required")
	}
	return &GetTimelineRequest{UserID: userID}, nil
}
