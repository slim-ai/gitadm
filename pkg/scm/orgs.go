package scm

import (
	"context"

	"github.com/google/go-github/v42/github"
	"github.com/xanzy/go-gitlab"
)

func GetUserOrganizationMembershipsForGitlab() ([]*gitlab.Group, error) {
	opt := true
	groups, _, err := GetGitLabClient().Groups.ListGroups(
		&gitlab.ListGroupsOptions{TopLevelOnly: &opt},
	)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func GetUserOrganizationMembershipsForGitHub(username string) ([]*github.Organization, error) {
	opts := &github.ListOptions{
		Page: 1,
	}
	ctx := context.Background()
	groups, _, err := GetGitHubClient().Organizations.List(ctx, username, opts)

	if err != nil {
		return nil, err
	}
	return groups, nil
}
