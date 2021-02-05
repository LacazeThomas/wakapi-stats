package utils

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/lacazethomas/wakapi-stats/config"
)

func initRepositoryContentFileOptions(gitProfile config.GitConfig, imageContent []byte) *github.RepositoryContentFileOptions {
	return &github.RepositoryContentFileOptions{
		Branch:  github.String(gitProfile.Branch),
		Message: github.String(gitProfile.Message),
		Committer: &github.CommitAuthor{
			Name:  github.String(gitProfile.CommitName),
			Email: github.String(gitProfile.CommitEmail),
		},
		Author: &github.CommitAuthor{
			Name:  github.String(gitProfile.CommitName),
			Email: github.String(gitProfile.CommitEmail),
		},
		Content: imageContent,
	}
}

//PublishToGithub with git config and file to push. It will automatically create or update file
func PublishToGithub(gitProfile config.GitConfig, filename string) (err error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitProfile.AccessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	imageContent, err := getImageBytes(filename)
	if err != nil{
		return errors.Wrap(err, "error getting bytes from filename")
	}

	commitOption := initRepositoryContentFileOptions(gitProfile, imageContent)
	ctx := context.Background()

	_, response, err := client.Repositories.CreateFile(ctx, gitProfile.UserName, gitProfile.RepoName, filename, commitOption)

	if err != nil {
		zap.S().Debugf("Receive response%+v", response)

		repositoryContent, _, response, err := client.Repositories.GetContents(ctx, gitProfile.UserName, gitProfile.RepoName, filename, nil)
		if err != nil {
			zap.S().Debugf("Receive response & repositoryContent %+v --- %+v", response, repositoryContent)
			return errors.Wrap(err, "error getting file content from github")
		}

		commitOption.SHA = repositoryContent.SHA

		_, response, err = client.Repositories.UpdateFile(ctx, gitProfile.UserName, gitProfile.RepoName, filename, commitOption)
		if err != nil {
			zap.S().Debugf("Receive response%+v", response)
			return errors.Wrap(err, "error updating file")
		}
	}

	return nil
}
