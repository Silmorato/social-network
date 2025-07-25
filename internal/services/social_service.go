package services

import (
	"github.com/google/uuid"
	"net/http"
	"sort"
	"time"
	"twitter-clone/internal/adapters/database/model"
	customErr "twitter-clone/internal/errors"
	"twitter-clone/internal/ports"
)

type SocialService struct {
	tweetRepository  ports.TweetRepository
	followRepository ports.FollowRepository
	userRepository   ports.UserRepository
}

// NewSocialService creates a new instance of SocialService.
func NewSocialService(tweetRepo ports.TweetRepository, followRepo ports.FollowRepository, userRepo ports.UserRepository) *SocialService {
	return &SocialService{
		tweetRepository:  tweetRepo,
		followRepository: followRepo,
		userRepository:   userRepo,
	}
}

func (s *SocialService) PublishTweet(userID, content string) (*model.Tweet, *customErr.APIError) {
	if !s.userRepository.Exists(userID) {
		return nil, customErr.NewAPIError(http.StatusBadRequest, customErr.ErrUserNotFound, nil)
	}

	tweet := &model.Tweet{
		ID:        uuid.New(),
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := s.tweetRepository.Save(tweet); err != nil {
		return nil, customErr.NewAPIError(http.StatusInternalServerError, customErr.ErrStorage, err)
	}

	return tweet, nil
}

func (s *SocialService) GetTimeline(userID string) ([]*model.Tweet, *customErr.APIError) {
	userIDs, err := s.followRepository.GetFollowings(userID)
	if err != nil {
		return nil, customErr.NewAPIError(http.StatusInternalServerError, "could not get followings", err)
	}

	userIDs = append(userIDs, userID)

	tweets, err := s.tweetRepository.GetAllByUserIDs(userIDs)
	if err != nil {
		return nil, customErr.NewAPIError(http.StatusInternalServerError, "could not fetch tweets for timeline", err)
	}

	tweets = sortTweetsByDateDesc(tweets)
	if len(tweets) == 0 {
		return nil, customErr.NewAPIError(http.StatusNotFound, "no tweets available", nil)
	}

	return tweets, nil
}

func (s *SocialService) FollowUser(followerID, followingID string) (*string, *customErr.APIError) {
	if !s.userRepository.Exists(followerID) || !s.userRepository.Exists(followingID) {
		return nil, customErr.NewAPIError(http.StatusNotFound, "one or both users not found", nil)
	}

	if s.followRepository.IsFollowing(followerID, followingID) {
		return nil, customErr.NewAPIError(http.StatusConflict, "user is already followed", nil)
	}

	if err := s.followRepository.AddFollow(followerID, followingID); err != nil {
		return nil, customErr.NewAPIError(http.StatusInternalServerError, "could not follow user", err)
	}
	success := "user followed successfully"
	return &success, nil
}

func sortTweetsByDateDesc(tweets []*model.Tweet) []*model.Tweet {
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].CreatedAt.After(tweets[j].CreatedAt)
	})

	return tweets
}
