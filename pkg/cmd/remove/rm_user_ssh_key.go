package add

import (
	"errors"

	"github.com/slim-ai/gitadm/pkg/scm"
	"github.com/urfave/cli/v2"
)

func removeUserSshKey(c *cli.Context) error {
	title := c.String("title")
	key, err := scm.FindSshKey(title)
	if errors.Is(err, scm.ErrSshKeyNotFound) {
		return nil
	} else if err != nil {
		return err
	}
	if err := scm.DeleteSshKey(key); err != nil {
		return err
	}
	return nil
}
