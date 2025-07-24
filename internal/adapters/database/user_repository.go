package database

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Exists(userID string) bool {
	var count int64
	r.db.Table("users").Where("id = ?", userID).Count(&count)
	return count > 0
}
