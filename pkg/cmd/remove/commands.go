package add

import (
	"github.com/slim-ai/gitadm/pkg/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	commands := &cli.Command{
		Name:  "rm",
		Usage: "remove command set",
		Subcommands: []*cli.Command{
			{
				Name:   "ssh-key",
				Usage:  "deletes an ssh key for the current user",
				Action: removeUserSshKey,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "title",
						Usage:    "'TITLE' name for the new key",
						Required: true,
					},
				},
			},
		},
	}
	cmd.AddCommands(commands)
}
