package dto

import (
	"encoding/json"
	"fmt"
	"io"
)

type FollowRequest struct {
	FollowerID  string `json:"follower_id" validate:"required"`
	FollowingID string `json:"following_id" validate:"required"`
}

func (r *FollowRequest) FromJSON(body io.Reader) error {
	if err := json.NewDecoder(body).Decode(r); err != nil {
		return fmt.Errorf("decode error: %w", err)
	}
	if err := Validate.Struct(r); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}
