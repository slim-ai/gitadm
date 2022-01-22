package add

import (
	"github.com/slim-ai/gitadm/pkg/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	commands := &cli.Command{
		Name:  "add",
		Usage: "add commands",
		Subcommands: []*cli.Command{
			{
				Name:   "ssh-key",
				Usage:  "add an ssh key for the current user",
				Action: addSshKey,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "title",
						Usage:    "'TITLE' name for the new key",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "file",
						Usage:    "'FILE' path to the desired ssh public key file",
						Required: true,
					},
					&cli.BoolFlag{
						Name:        "overwrite",
						Usage:       "overwrite flag",
						Required:    false,
						DefaultText: "false",
						Value:       false,
					},
				},
			},
		},
	}
	cmd.AddCommands(commands)
}
