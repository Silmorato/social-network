package database

import (
	"gorm.io/gorm"
	"twitter-clone/internal/adapters/database/model"
)

type TweetRepository struct {
	db *gorm.DB
}

func NewTweetRepository(db *gorm.DB) *TweetRepository {
	return &TweetRepository{db: db}
}

func (r *TweetRepository) Save(tweet *model.Tweet) error {
	return r.db.Create(tweet).Error
}

func (r *TweetRepository) GetAllByUserIDs(userIDs []string) ([]*model.Tweet, error) {
	var tweets []*model.Tweet
	err := r.db.
		Where("user_id IN ?", userIDs).
		Order("created_at DESC").
		Find(&tweets).Error

	return tweets, err
}
