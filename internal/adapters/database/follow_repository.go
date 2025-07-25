package database

import (
	"gorm.io/gorm"
	"twitter-clone/internal/adapters/database/model"
)

type FollowRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) *FollowRepository {
	return &FollowRepository{db: db}
}

// Save a follow relationship in the database
func (r *FollowRepository) AddFollow(followerID, followingID string) error {
	follow := &model.Follow{
		FollowerID:  followerID,
		FollowingID: followingID,
	}
	return r.db.Create(follow).Error
}

// Retrieve all user IDs that the given user follows
func (r *FollowRepository) GetFollowings(userID string) ([]string, error) {
	var followings []string
	err := r.db.
		Model(&model.Follow{}).
		Where("follower_id = ?", userID).
		Pluck("following_id", &followings).Error

	return followings, err
}

func (r *FollowRepository) IsFollowing(followerID, followingID string) bool {
	var count int64
	err := r.db.
		Model(&model.Follow{}).
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Count(&count).Error
	return err == nil && count > 0
}
