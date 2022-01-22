package scm

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/xanzy/go-gitlab"
)

var (
	ErrSshKeyAlreadyExists = errors.New("ssh key already exists")
	ErrSshKeyNotFound      = errors.New("ssh key not found")
)

func AddSshKeys(title string, filename string, overwrite bool) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return err
	}
	if key, err := FindSshKey(title); !errors.Is(err, ErrSshKeyNotFound) {
		return err
	} else if key != nil {
		if !overwrite {
			return ErrSshKeyAlreadyExists
		}
		if err := DeleteSshKey(key); err != nil {
			return err
		}
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	keyMaterial := string(b)
	// Setup new key
	if _, _, err := Client().Users.AddSSHKey(
		&gitlab.AddSSHKeyOptions{
			Title: &title,
			Key:   &keyMaterial,
		},
	); err != nil {
		return err
	}
	return nil
}

func DeleteSshKey(key *gitlab.SSHKey) error {
	if _, err := Client().Users.DeleteSSHKey(key.ID); err != nil {
		return err
	}
	return nil
}

func FindSshKey(title string) (*gitlab.SSHKey, error) {
	keys, _, err := Client().Users.ListSSHKeys()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		if key.Title == title {
			return key, nil
		}
	}
	return nil, ErrSshKeyNotFound
}

func GetSshKeys() ([]*gitlab.SSHKey, error) {
	keys, _, err := Client().Users.ListSSHKeys()
	if err != nil {
		return nil, err
	}
	return keys, nil
}
