package describe

import (
	"github.com/slim-ai/gitadm/pkg/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	commands := &cli.Command{
		Name:  "describe",
		Usage: "describe command set",
		Subcommands: []*cli.Command{
			{
				Name:   "user",
				Usage:  "get the users details",
				Action: getUserDetails,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "username",
						Usage:    "gitlab username",
						Required: true,
					},
				},
			},
			{
				Name:   "orgs",
				Usage:  "get current user organization memberships",
				Action: getUserOrgMemberships,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "short",
						Usage:       "short format results in a comma separated list of orgs spaces",
						Aliases:     []string{"s"},
						Required:    false,
						DefaultText: "false",
						Value:       false,
					},
				},
			},
			{
				Name:   "ssh-keys",
				Usage:  "get a list of the ssh keys for the current user",
				Action: getUserSshKeys,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "short",
						Usage:       "short outputs a list of key titles only",
						Aliases:     []string{"s"},
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
