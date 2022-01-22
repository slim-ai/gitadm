package scm

import (
	"github.com/xanzy/go-gitlab"
)

func GetUserOrganizationMemberships() ([]*gitlab.Group, error) {
	opt := true
	groups, _, err := Client().Groups.ListGroups(
		&gitlab.ListGroupsOptions{TopLevelOnly: &opt},
	)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
