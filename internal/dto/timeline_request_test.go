package dto

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewGetTimelineRequest_Valid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/timeline", nil)
	req.Header.Set("X-User-ID", "user123")

	result, err := ParseGetTimelineRequest(req)

	assert.NoError(t, err)
	assert.Equal(t, "user123", result.UserID)
}

func TestNewGetTimelineRequest_MissingHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/timeline", nil)

	result, err := ParseGetTimelineRequest(req)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "X-User-ID header is required")
}
