package scm

import (
	"log"
	"sync"

	"github.com/slim-ai/gitadm/pkg/config"
	"github.com/xanzy/go-gitlab"
)

var (
	gitlabClient     *gitlab.Client
	gitlabClientOnce sync.Once
)

func Client() *gitlab.Client {
	gitlabClientOnce.Do(func() {
		git, err := gitlab.NewClient(config.Config().Token)
		if err != nil {
			log.Fatalf("Failed to create gitlab client: %v", err)
		}
		gitlabClient = git
	})
	return gitlabClient
}
