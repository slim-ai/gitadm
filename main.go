package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/slim-ai/gitadm/pkg/build"
	"github.com/slim-ai/gitadm/pkg/cmd"
	"github.com/slim-ai/gitadm/pkg/common"
	"github.com/slim-ai/gitadm/pkg/config"
	"github.com/slim-ai/gitadm/pkg/context"
	"github.com/urfave/cli/v2"

	// Auto-Import Command Sets
	_ "github.com/slim-ai/gitadm/pkg/cmd/add"
	_ "github.com/slim-ai/gitadm/pkg/cmd/describe"
	_ "github.com/slim-ai/gitadm/pkg/cmd/remove"
)

const banner = `Gitlab client`

func main() {
	ctrlCHandler()

	app := &cli.App{
		Usage:                build.ApplicationDescription,
		Description:          build.ApplicationDescription,
		Version:              fmt.Sprintf("%s.%s %s %s-%s", build.Tag, build.Rev, build.Time, runtime.GOOS, runtime.GOARCH),
		Compiled:             time.Now(),
		Copyright:            build.Copyright,
		Commands:             cmd.GetCommands(),
		EnableBashCompletion: true,
		Before: func(c *cli.Context) error {
			if err := context.Set(c); err != nil { // set global context and load configuration file
				return err
			}
			t := context.GetOrDie().String(common.ParamToken)
			if (t == "" || t == common.DefaultToken) && !(strings.ToLower(c.Command.Name) == "h" || strings.ToLower(c.Command.Name) == "help") {
				cli.ShowAppHelp(c)
				return fmt.Errorf("\"token\" parameter not set. Set the %s environment variable, add your token to the configuration file located at %s, or supply it as a commandline argument",
					common.EnvVarAPIToken,
					config.FullFilePath(build.ApplicationName))
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    common.ParamConfig,
				EnvVars: []string{common.EnvVarConfigFilePath},
				Aliases: []string{"c"},
				Usage:   fmt.Sprintf("Configuration file location. Defaults to %s", config.Directory(build.ApplicationName)),
				Value:   config.FullFilePath(build.ApplicationName),
			},
			&cli.StringFlag{
				Name:    common.ParamToken,
				EnvVars: []string{common.EnvVarAPIToken},
				Aliases: []string{"t"},
				Usage:   "Gitlab API token",
			},
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("\n\nERROR:\n%s\n\n", err.Error())
	}
}

////////////

func ctrlCHandler() {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)
	go func() {
		<-stopCh
		os.Exit(0)
	}()
}
