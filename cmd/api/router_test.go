package api

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"testing"
	httpAdap "twitter-clone/internal/adapters/http"
)

func TestRouterRoutesCount(t *testing.T) {
	mockHandler := &httpAdap.MockSocialHandler{}
	router := NewRouter(mockHandler)

	count := 0
	err := router.(*mux.Router).Walk(func(route *mux.Route, r *mux.Router, ancestors []*mux.Route) error {
		count++
		return nil
	})
	assert.Nil(t, err)
	assert.Equal(t, 3, count)
}
