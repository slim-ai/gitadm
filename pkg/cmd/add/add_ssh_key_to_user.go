package add

import (
	"github.com/slim-ai/gitadm/pkg/scm"
	"github.com/urfave/cli/v2"
)

func addSshKey(c *cli.Context) error {
	title := c.String("title")
	filename := c.String("file")
	overwrite := c.Bool("overwrite")
	if err := scm.AddSshKeys(title, filename, overwrite); err != nil {
		return err
	}
	return nil
}
