package ports

type UserRepository interface {
	Exists(userID string) bool
}
