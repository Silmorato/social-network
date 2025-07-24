package dto

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFollowRequest_FromJSON_Valid(t *testing.T) {
	jsonBody := `{"follower_id": "user1", "following_id": "user2"}`
	req := FollowRequest{}

	err := req.FromJSON(bytes.NewBufferString(jsonBody))

	assert.NoError(t, err)
	assert.Equal(t, "user1", req.FollowerID)
	assert.Equal(t, "user2", req.FollowingID)
}

func TestFollowRequest_FromJSON_InvalidCases(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantErrMsg string
	}{
		{
			name:       "missing follower_id",
			body:       `{"following_id": "123"}`,
			wantErrMsg: "validation error",
		},
		{
			name:       "missing following_id",
			body:       `{"follower_id": "456"}`,
			wantErrMsg: "validation error",
		},
		{
			name:       "invalid json",
			body:       `{"follower_id": "1", "following_id": 123`,
			wantErrMsg: "decode error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req FollowRequest
			err := req.FromJSON(bytes.NewReader([]byte(tt.body)))
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErrMsg)
		})
	}
}
