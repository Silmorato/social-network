package database

import "github.com/stretchr/testify/mock"

type FollowRepositoryMock struct {
	mock.Mock
}

func (m *FollowRepositoryMock) AddFollow(followerID, followingID string) error {
	args := m.Called(followerID, followingID)
	return args.Error(0)
}

func (m *FollowRepositoryMock) GetFollowings(userID string) ([]string, error) {
	args := m.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}

func (m *FollowRepositoryMock) IsFollowing(followerID, followingID string) bool {
	args := m.Called(followerID, followingID)
	return args.Bool(0)
}
