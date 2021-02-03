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

func TestLoad(t *testing.T) {
	os.Setenv("PLUGIN_API_URL", "http://localhost")
	os.Setenv("PLUGIN_API_KEY", "fdsg-sfdg-dsfg-sdfg")
	os.Setenv("PLUGIN_ENVIRONMENT", "dev")

	os.Setenv("PLUGIN_ACCESSTOKEN", "dfdjgfdjfgdjfddjgf")
	os.Setenv("PLUGIN_BRANCH", "main")
	os.Setenv("PLUGIN_MESSAGE", "Update images from wakapi-stats")
	os.Setenv("PLUGIN_COMMITNAME", "LACAZE Thomas")
	os.Setenv("PLUGIN_COMMITEMAIL", "contact@thomaslacaze.fr")
	os.Setenv("PLUGIN_USERNAME", "lacazethomas")
	os.Setenv("PLUGIN_REPO", "testname")

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
