package scm

import (
	"context"
	"log"
	"sync"

	"github.com/google/go-github/v42/github"
	"github.com/slim-ai/gitadm/pkg/config"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"
)

var (
	gitlabClient     *gitlab.Client
	gitlabClientOnce sync.Once
	githubClient     *github.Client
	githubClientOnce sync.Once
)

func GetGitLabClient() *gitlab.Client {

	gitlabClientOnce.Do(func() {
		git, err := gitlab.NewClient(config.Config().Token)
		if err != nil {
			log.Fatalf("Failed to create gitlab client: %v", err)
		}
		gitlabClient = git
	})
	return gitlabClient
}

func GetGitHubClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(ctx, ts)
	githubClientOnce.Do(func() {
		git := github.NewClient(tc)
		githubClient = git
	})
	return githubClient
}
