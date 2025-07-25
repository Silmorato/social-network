package database

import (
	"github.com/stretchr/testify/mock"
	"twitter-clone/internal/adapters/database/model"
)

type TweetRepositoryMock struct {
	mock.Mock
}

func (m *TweetRepositoryMock) Save(tweet *model.Tweet) error {
	args := m.Called(tweet)
	return args.Error(0)
}

func (m *TweetRepositoryMock) GetAllByUserIDs(userIDs []string) ([]*model.Tweet, error) {
	args := m.Called(userIDs)
	return args.Get(0).([]*model.Tweet), args.Error(1)
}
