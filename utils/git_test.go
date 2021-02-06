package utils

import (
	"testing"

	"github.com/google/go-github/github"
	"github.com/lacazethomas/wakapi-stats/config"
	"github.com/stretchr/testify/assert"
)

var (
	testGitConfig = config.GitConfig{
		AccessToken: "dfdjgfdjfgdjfddjgf",
		Branch:      "main",
		Message:     "Update images from wakapi-stats",
		CommitName:  "LACAZE Thomas",
		CommitEmail: "contact@thomaslacaze.fr",
		UserName:    "lacazethomas",
		RepoName:    "testname",
	}
)

func TestInitRepositoryContentFileOptions(t *testing.T) {

	expected := &github.RepositoryContentFileOptions{
		Branch:  github.String(testGitConfig.Branch),
		Message: github.String(testGitConfig.Message),
		Committer: &github.CommitAuthor{
			Name:  github.String(testGitConfig.CommitName),
			Email: github.String(testGitConfig.CommitEmail),
		},
		Author: &github.CommitAuthor{
			Name:  github.String(testGitConfig.CommitName),
			Email: github.String(testGitConfig.CommitEmail),
		},
		Content: []byte("Hello"),
	}

	actual := initRepositoryContentFileOptions(testGitConfig, []byte("Hello"))

	assert.Equal(t, expected, actual)
}

func TestPublishToGithub(t *testing.T) {

	err := PublishToGithub(testGitConfig, "null.png")
	assert.Error(t, err)
	assert.Regexp(t, ".*error getting bytes from filename.*", err.Error())

	err = PublishToGithub(testGitConfig, "projects.png")
	assert.Error(t, err)
	assert.Regexp(t, ".*error getting file content from github.*", err.Error())

}
