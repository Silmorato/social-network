package dto

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTweetRequest_FromJSON_Valid(t *testing.T) {
	body := `{"user_id": "123", "content": "Hello, this is a test tweet"}`
	var req CreateTweetRequest

	err := req.FromJSON(bytes.NewReader([]byte(body)))

	assert.NoError(t, err)
	assert.Equal(t, "123", req.UserID)
	assert.Equal(t, "Hello, this is a test tweet", req.Content)
}

func TestCreateTweetRequest_FromJSON_InvalidCases(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantErrMsg string
	}{
		{
			name:       "missing user_id",
			body:       `{"content": "Hello, this is a test tweet"}`,
			wantErrMsg: "validation error",
		},
		{
			name:       "missing content",
			body:       `{"user_id": "123"}`,
			wantErrMsg: "validation error",
		},
		{
			name:       "empty content",
			body:       `{"user_id": "123", "content": ""}`,
			wantErrMsg: "validation error",
		},
		{
			name:       "invalid json format",
			body:       `{"user_id": "123", "content": "Hello"`,
			wantErrMsg: "invalid JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req CreateTweetRequest
			err := req.FromJSON(bytes.NewReader([]byte(tt.body)))
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErrMsg)
		})
	}
}
