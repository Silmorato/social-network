package domain

import (
	"github.com/google/uuid"
	"time"
)

type Tweet struct {
	ID        uuid.UUID
	UserID    string
	Content   string
	CreatedAt time.Time
}

func (t *Tweet) GetKey() string {
	return t.ID.String()
}
