package services

import (
	"twitter-clone/internal/domain"
	customerr "twitter-clone/internal/errors"

	"github.com/stretchr/testify/mock"
)

type SocialServiceMock struct {
	mock.Mock
}

func (m *SocialServiceMock) PublishTweet(userID, content string) (*domain.Tweet, *customerr.APIError) {
	args := m.Called(userID, content)
	tweet, _ := args.Get(0).(*domain.Tweet)
	err, _ := args.Get(1).(*customerr.APIError)
	return tweet, err
}

func (m *SocialServiceMock) GetTimeline(userID string) ([]*domain.Tweet, *customerr.APIError) {
	args := m.Called(userID)
	tweets, _ := args.Get(0).([]*domain.Tweet)
	err, _ := args.Get(1).(*customerr.APIError)
	return tweets, err
}

func (m *SocialServiceMock) FollowUser(followerID, followingID string) (*string, *customerr.APIError) {
	args := m.Called(followerID, followingID)
	follow, _ := args.Get(0).(*string)
	err, _ := args.Get(1).(*customerr.APIError)
	return follow, err
}
