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
		Env:    "dev",
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
	_ = os.Setenv("PLUGIN_API_URL", "http://localhost")
	_ = os.Setenv("PLUGIN_API_KEY", "fdsg-sfdg-dsfg-sdfg")
	_ = os.Setenv("PLUGIN_ENVIRONMENT", "dev")

	_ = os.Setenv("PLUGIN_ACCESSTOKEN", "dfdjgfdjfgdjfddjgf")
	_ = os.Setenv("PLUGIN_BRANCH", "main")
	_ = os.Setenv("PLUGIN_MESSAGE", "Update images from wakapi-stats")
	_ = os.Setenv("PLUGIN_COMMITNAME", "LACAZE Thomas")
	_ = os.Setenv("PLUGIN_COMMITEMAIL", "contact@thomaslacaze.fr")
	_ = os.Setenv("PLUGIN_USERNAME", "lacazethomas")
	_ = os.Setenv("PLUGIN_REPO", "testname")

	appConfig, gitConfig := Load()

	assert.Equal(t, appConfig, AppConfig{
		APIKey: "fdsg-sfdg-dsfg-sdfg",
		APIURL: "http://localhost",
		Env:    "dev",
	})

	assert.Equal(t, gitConfig, GitConfig{
		AccessToken: "dfdjgfdjfgdjfddjgf",
		Branch:      "main",
		Message:     "Update images from wakapi-stats",
		CommitName:  "LACAZE Thomas",
		CommitEmail: "contact@thomaslacaze.fr",
		UserName:    "lacazethomas",
		RepoName:    "testname",
	})
}
