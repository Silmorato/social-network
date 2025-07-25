package ports

import (
	"twitter-clone/internal/adapters/database/model"
	customerr "twitter-clone/internal/errors"
)

type SocialService interface {
	PublishTweet(userID, content string) (*model.Tweet, *customerr.APIError)
	GetTimeline(userID string) ([]*model.Tweet, *customerr.APIError)
	FollowUser(followerID, followingID string) (*string, *customerr.APIError)
}
