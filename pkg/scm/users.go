package scm

import (
	"context"
	"errors"
	"log"

	"github.com/google/go-github/v42/github"
	"github.com/xanzy/go-gitlab"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

func GetUserDetails(username string) (*gitlab.User, error) {
	users, _, err := GetGitLabClient().Users.ListUsers(
		&gitlab.ListUsersOptions{
			Username: &username,
		},
	)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, ErrUserNotFound
	}
	return users[0], nil
}

func GetUserDetailsforGitHub(username string) (*github.User, error) {

	ctx := context.Background()
	user, _, err := GetGitHubClient().Users.Get(ctx, username)
	if err != nil {
		log.Fatalf("error listing users: %v", err)
	}
	return user, nil
}
