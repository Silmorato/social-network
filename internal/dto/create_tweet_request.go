package dto

import (
	"encoding/json"
	"fmt"
	"io"
)

type CreateTweetRequest struct {
	UserID  string `json:"user_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// FromJSON parses the request body into the CreateTweetRequest struct.
func (r *CreateTweetRequest) FromJSON(body io.Reader) error {
	if err := json.NewDecoder(body).Decode(r); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	if err := Validate.Struct(r); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}
