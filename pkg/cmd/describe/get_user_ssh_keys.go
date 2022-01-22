package describe

import (
	"github.com/slim-ai/gitadm/pkg/scm"
	"github.com/slim-ai/gitadm/pkg/util"
	"github.com/urfave/cli/v2"
)

func getUserSshKeys(c *cli.Context) error {
	keys, err := scm.GetSshKeys()
	if err != nil {
		return err
	}
	if !c.Bool("short") {
		util.PrettyPrint(keys)
		return nil
	}
	titles := make([]string, len(keys))
	for i, key := range keys {
		titles[i] = key.Title
	}
	util.PrettyPrint(titles)
	return nil
}
