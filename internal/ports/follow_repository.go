package ports

type FollowRepository interface {
	GetFollowings(userID string) ([]string, error)
	AddFollow(followerID, followingID string) error
	IsFollowing(followerID, followingID string) bool
}
