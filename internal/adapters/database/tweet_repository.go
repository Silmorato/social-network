package database

import (
	"gorm.io/gorm"
	"twitter-clone/internal/domain"
)

type TweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) *TweetRepository {
	return &TweetRepository{db: db}
}

func (r *TweetRepository) Save(tweet *domain.Tweet) error {
	return r.db.Create(tweet).Error
}

func (r *TweetRepository) GetAllByUserIDs(userIDs []string) ([]*domain.Tweet, error) {
	var tweets []*domain.Tweet
	err := r.db.
		Where("user_id IN ?", userIDs).
		Order("created_at DESC").
		Find(&tweets).Error

	return tweets, err
}
