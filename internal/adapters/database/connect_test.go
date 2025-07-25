package database

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	os.Setenv("SCOPE", "local")
	assert.Equal(t, "config/local.properties", getConfigPath())

	os.Setenv("SCOPE", "render")
	assert.Equal(t, "config/render.properties", getConfigPath())

	os.Unsetenv("SCOPE")
	assert.Equal(t, "config/render.properties", getConfigPath())
}
