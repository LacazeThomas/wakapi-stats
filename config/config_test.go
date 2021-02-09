package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDev(t *testing.T) {
	assert.True(t, IsDev("dev"))
	assert.True(t, IsDev("development"))
	assert.False(t, IsDev("prod"))
	assert.False(t, IsDev("production"))
	assert.False(t, IsDev("anything else"))
}

func TestAppConfigIsDev(t *testing.T) {
	c := AppConfig{
		Env: "dev",
	}
	assert.True(t, c.IsDev())

	c.Env = "development"
	assert.True(t, c.IsDev())

	c.Env = "prod"
	assert.False(t, c.IsDev())

	c.Env = "production"
	assert.False(t, c.IsDev())

	c.Env = "anything else"
	assert.False(t, c.IsDev())
}

func TestLoad(t *testing.T) {
	_ = os.Setenv("ENVIRONMENT", "dev")

	appConfig := GetAppConfig()

	assert.Equal(t, appConfig, AppConfig{
		Env: "dev",
	})
}
