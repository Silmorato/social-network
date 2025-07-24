package ports

import (
	"twitter-clone/internal/domain"
	customerr "twitter-clone/internal/errors"
)

type SocialService interface {
	PublishTweet(userID, content string) (*domain.Tweet, *customerr.APIError)
	GetTimeline(userID string) ([]*domain.Tweet, *customerr.APIError)
	FollowUser(followerID, followingID string) (*string, *customerr.APIError)
}
