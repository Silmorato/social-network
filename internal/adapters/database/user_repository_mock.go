package database

import (
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Exists(userID string) bool {
	args := m.Called(userID)
	return args.Bool(0)
}
