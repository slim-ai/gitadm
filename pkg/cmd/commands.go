package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var commands []*cli.Command

func GetCommands() []*cli.Command {
	return commands
}

func AddCommands(cmd *cli.Command) {
	commands = append(commands, cmd)
}

func NotYetImplemented(c *cli.Context) error {
	return fmt.Errorf("%s is not implemented yet - please contact support.\n", c.Command.Name)
}
