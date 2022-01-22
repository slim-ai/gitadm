package scm

import (
	"errors"

	"github.com/xanzy/go-gitlab"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

func GetUserDetails(username string) (*gitlab.User, error) {
	users, _, err := Client().Users.ListUsers(
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
