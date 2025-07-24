package http

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"twitter-clone/internal/domain"
	customerr "twitter-clone/internal/errors"
	"twitter-clone/internal/services"
)

func TestCreateTweetHandler_Valid(t *testing.T) {
	mockService := new(services.SocialServiceMock)

	handler := &SocialHandler{Service: mockService}

	body := `{"user_id": "user123", "content": "¡Hello!"}`
	req := httptest.NewRequest(http.MethodPost, "/tweets", strings.NewReader(body))
	w := httptest.NewRecorder()

	expectedTweet := &domain.Tweet{
		ID:        uuid.New(),
		UserID:    "user123",
		Content:   "¡Hello!",
		CreatedAt: time.Now(),
	}

	mockService.On("PublishTweet", "user123", "¡Hello!").Return(expectedTweet, nil)

	handler.CreateTweet(w, req)

	res := w.Result()
	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func TestCreateTweetHandler_InvalidCases(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		mockSetup      func(s *services.SocialServiceMock)
		expectedStatus int
	}{
		{
			name:           "invalid json",
			body:           `{"user_id": "user123", "content":`, // mal cerrado
			mockSetup:      func(s *services.SocialServiceMock) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "validation error - empty user_id",
			body:           `{"user_id": "", "content": "hola"}`,
			mockSetup:      func(s *services.SocialServiceMock) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "service error",
			body: `{"user_id": "user123", "content": "hola"}`,
			mockSetup: func(s *services.SocialServiceMock) {
				s.On("PublishTweet", "user123", "hola").
					Return(nil, customerr.NewAPIError(http.StatusInternalServerError, "error save", nil))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(services.SocialServiceMock)
			tt.mockSetup(mockService)

			handler := &SocialHandler{Service: mockService}
			req := httptest.NewRequest(http.MethodPost, "/tweets", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			handler.CreateTweet(w, req)

			res := w.Result()

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}

func TestGetTimelineHandler_Valid(t *testing.T) {
	userID := "user123"
	tweets := []*domain.Tweet{
		{
			ID:        uuid.New(),
			UserID:    userID,
			Content:   "Hello world",
			CreatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			UserID:    userID,
			Content:   "Hello, first tweets",
			CreatedAt: time.Now(),
		},
	}

	mockService := new(services.SocialServiceMock)
	mockService.On("GetTimeline", userID).Return(tweets, nil)

	handler := &SocialHandler{Service: mockService}

	req := httptest.NewRequest(http.MethodGet, "/timeline", nil)
	req.Header.Set("X-User-ID", userID)
	w := httptest.NewRecorder()

	handler.GetTimeline(w, req)

	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetTimelineHandler_InvalidCases(t *testing.T) {
	tests := []struct {
		name           string
		headerValue    string
		mockSetup      func(s *services.SocialServiceMock)
		expectedStatus int
	}{
		{
			name:           "missing header",
			headerValue:    "",
			mockSetup:      func(s *services.SocialServiceMock) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:        "service error",
			headerValue: "user123",
			mockSetup: func(s *services.SocialServiceMock) {
				s.On("GetTimeline", "user123").
					Return(nil, customerr.NewAPIError(http.StatusInternalServerError, "some error", nil))
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(services.SocialServiceMock)
			tt.mockSetup(mockService)

			handler := &SocialHandler{Service: mockService}
			req := httptest.NewRequest(http.MethodGet, "/timeline", nil)

			if tt.headerValue != "" {
				req.Header.Set("X-User-ID", tt.headerValue)
			}
			w := httptest.NewRecorder()
			handler.GetTimeline(w, req)

			res := w.Result()
			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}
