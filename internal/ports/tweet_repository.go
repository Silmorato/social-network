package ports

import (
	"twitter-clone/internal/adapters/database/model"
)

type TweetRepository interface {
	Save(tweet *model.Tweet) error
	GetAllByUserIDs(userIDs []string) ([]*model.Tweet, error)
}
