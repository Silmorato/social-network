package ports

import (
	"twitter-clone/internal/domain"
)

type TweetRepository interface {
	Save(tweet *domain.Tweet) error
	GetAllByUserIDs(userIDs []string) ([]*domain.Tweet, error)
}
