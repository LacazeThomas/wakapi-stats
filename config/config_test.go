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
	os.Setenv("API_URL", "http://localhost")
	os.Setenv("API_KEY", "fdsg-sfdg-dsfg-sdfg")
	os.Setenv("ENVIRONMENT", "dev")

	os.Setenv("AccessToken", "dfdjgfdjfgdjfddjgf")
	os.Setenv("Branch", "main")
	os.Setenv("Message", "Update images from wakapi-stats")
	os.Setenv("CommitName", "LACAZE Thomas")
	os.Setenv("CommitEmail", "contact@thomaslacaze.fr")
	os.Setenv("UserName", "lacazethomas")
	os.Setenv("RepoName", "testname")

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
