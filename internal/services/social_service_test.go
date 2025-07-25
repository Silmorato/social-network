package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"time"
	"twitter-clone/internal/adapters/database"
	"twitter-clone/internal/adapters/database/model"
	customerr "twitter-clone/internal/errors"
)

func TestPublishTweet_Valid(t *testing.T) {
	userID := "user123"
	content := "Hello, this is a test tweet"

	mockUserRepo := new(database.UserRepositoryMock)
	mockTweetRepo := new(database.TweetRepositoryMock)

	mockUserRepo.On("Exists", userID).Return(true)
	mockTweetRepo.On("Save", mock.Anything).Return(nil)

	service := NewSocialService(mockTweetRepo, nil, mockUserRepo)
	tweet, err := service.PublishTweet(userID, content)

	assert.Nil(t, err)
	assert.NotNil(t, tweet)
	assert.Equal(t, userID, tweet.UserID)
	assert.Equal(t, content, tweet.Content)
	assert.WithinDuration(t, time.Now(), tweet.CreatedAt, time.Second)
}

func TestPublishTweet_InvalidCases(t *testing.T) {
	userID := "user123"
	content := "test tweet"

	tests := []struct {
		name            string
		mockUserExists  bool
		mockSaveError   error
		expectedStatus  int
		expectedMessage string
	}{
		{
			name:            "user does not exist",
			mockUserExists:  false,
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: customerr.ErrUserNotFound,
		},
		{
			name:            "tweet save fails",
			mockUserExists:  true,
			mockSaveError:   errors.New("DB error"),
			expectedStatus:  http.StatusInternalServerError,
			expectedMessage: customerr.ErrStorage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := new(database.UserRepositoryMock)
			mockTweetRepo := new(database.TweetRepositoryMock)

			mockUserRepo.On("Exists", userID).Return(tt.mockUserExists)

			if tt.mockUserExists {
				mockTweetRepo.
					On("Save", mock.Anything).
					Return(tt.mockSaveError)
			}

			service := NewSocialService(mockTweetRepo, nil, mockUserRepo)
			tweet, apiErr := service.PublishTweet(userID, content)

			assert.Nil(t, tweet)
			assert.NotNil(t, apiErr)
			assert.Equal(t, tt.expectedStatus, apiErr.Status)
			assert.Equal(t, tt.expectedMessage, apiErr.Message)

			mockUserRepo.AssertExpectations(t)
			mockTweetRepo.AssertExpectations(t)
		})
	}
}

func TestGetTimeline_Valid(t *testing.T) {
	userID := "user123"
	followings := []string{"user234", "user345"}
	expectedUserIDs := append(followings, userID)

	mockTweetRepo := new(database.TweetRepositoryMock)
	mockFollowRepo := new(database.FollowRepositoryMock)

	expectedTweet := &model.Tweet{
		ID:        uuid.New(),
		UserID:    "user234",
		Content:   "Hello, this is a test tweet",
		CreatedAt: time.Now(),
	}

	mockFollowRepo.On("GetFollowings", userID).Return(followings, nil)

	mockTweetRepo.On("GetAllByUserIDs", expectedUserIDs).Return([]*model.Tweet{expectedTweet}, nil)

	service := NewSocialService(mockTweetRepo, mockFollowRepo, nil)
	tweets, err := service.GetTimeline(userID)

	assert.Nil(t, err)
	assert.Len(t, tweets, 1)
	assert.Equal(t, expectedTweet.Content, tweets[0].Content)

	mockTweetRepo.AssertExpectations(t)
	mockFollowRepo.AssertExpectations(t)
}

func TestGetTimeline_InvalidCases(t *testing.T) {
	userID := "user123"

	tests := []struct {
		name            string
		mockFollowError error
		mockTweetError  error
		mockFollowings  []string
		mockTweets      []*model.Tweet
		expectedStatus  int
		expectedMessage string
	}{
		{
			name:            "error getting followings",
			mockFollowError: errors.New("db error"),
			expectedStatus:  http.StatusInternalServerError,
			expectedMessage: "could not get followings",
		},
		{
			name:            "error getting tweets",
			mockFollowings:  []string{"234"},
			mockTweetError:  errors.New("db error"),
			expectedStatus:  http.StatusInternalServerError,
			expectedMessage: "could not fetch tweets for timeline",
		},
		{
			name:            "no tweets available",
			mockFollowings:  []string{"234"},
			mockTweets:      []*model.Tweet{},
			expectedStatus:  http.StatusNotFound,
			expectedMessage: "no tweets available",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTweetRepo := new(database.TweetRepositoryMock)
			mockFollowRepo := new(database.FollowRepositoryMock)

			mockFollowRepo.On("GetFollowings", userID).Return(tt.mockFollowings, tt.mockFollowError)

			if tt.mockFollowError == nil {
				allUserIDs := append(tt.mockFollowings, userID)
				mockTweetRepo.On("GetAllByUserIDs", allUserIDs).Return(tt.mockTweets, tt.mockTweetError)
			}

			service := NewSocialService(mockTweetRepo, mockFollowRepo, nil)
			tweets, err := service.GetTimeline(userID)

			assert.Nil(t, tweets)
			assert.NotNil(t, err)
			assert.Equal(t, tt.expectedStatus, err.Status)
			assert.Equal(t, tt.expectedMessage, err.Message)

			mockTweetRepo.AssertExpectations(t)
			mockFollowRepo.AssertExpectations(t)
		})
	}
}

func TestFollowUser_Valid(t *testing.T) {
	followerID := "user1"
	followingID := "user2"

	mockUserRepo := new(database.UserRepositoryMock)
	mockFollowRepo := new(database.FollowRepositoryMock)

	mockUserRepo.On("Exists", followerID).Return(true)
	mockUserRepo.On("Exists", followingID).Return(true)
	mockFollowRepo.On("IsFollowing", "user1", "user2").Return(false)
	mockFollowRepo.On("AddFollow", followerID, followingID).Return(nil)

	service := NewSocialService(nil, mockFollowRepo, mockUserRepo)
	msg, err := service.FollowUser(followerID, followingID)

	assert.Nil(t, err)
	assert.Equal(t, "user followed successfully", *msg)
	mockUserRepo.AssertExpectations(t)
	mockFollowRepo.AssertExpectations(t)
}

func TestFollowUser_InvalidCases(t *testing.T) {
	type testCase struct {
		name           string
		setupMocks     func(*database.UserRepositoryMock, *database.FollowRepositoryMock)
		expectedStatus int
		expectedMsg    string
	}

	tests := []testCase{
		{
			name: "user not found",
			setupMocks: func(userRepo *database.UserRepositoryMock, followRepo *database.FollowRepositoryMock) {
				userRepo.On("Exists", "user1").Return(true)
				userRepo.On("Exists", "user2").Return(false)
			},
			expectedStatus: http.StatusNotFound,
			expectedMsg:    "one or both users not found",
		},
		{
			name: "user already followed",
			setupMocks: func(userRepo *database.UserRepositoryMock, followRepo *database.FollowRepositoryMock) {
				userRepo.On("Exists", "user1").Return(true)
				userRepo.On("Exists", "user2").Return(true)
				followRepo.On("IsFollowing", "user1", "user2").Return(true)
			},
			expectedStatus: http.StatusConflict,
			expectedMsg:    "user is already followed",
		},
		{
			name: "repository save error",
			setupMocks: func(userRepo *database.UserRepositoryMock, followRepo *database.FollowRepositoryMock) {
				userRepo.On("Exists", "user1").Return(true)
				userRepo.On("Exists", "user2").Return(true)
				followRepo.On("IsFollowing", "user1", "user2").Return(false)
				followRepo.On("AddFollow", "user1", "user2").Return(errors.New("db error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedMsg:    "could not follow user",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockUserRepo := new(database.UserRepositoryMock)
			mockFollowRepo := new(database.FollowRepositoryMock)

			tc.setupMocks(mockUserRepo, mockFollowRepo)

			service := NewSocialService(nil, mockFollowRepo, mockUserRepo)
			_, err := service.FollowUser("user1", "user2")

			assert.NotNil(t, err)
			assert.Equal(t, tc.expectedStatus, err.Status)
			assert.Equal(t, tc.expectedMsg, err.Message)
		})
	}
}
