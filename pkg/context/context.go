package context

import (
	"strings"

	"github.com/slim-ai/gitadm/pkg/common"
	"github.com/slim-ai/gitadm/pkg/config"
	"github.com/urfave/cli/v2"
)

var globalContext *cli.Context

func GetOrDie() *cli.Context {
	if globalContext == nil {
		panic("global context not set...code depends on that as a global")
	}
	return globalContext
}

func Set(c *cli.Context) error {
	globalContext = c // set first
	token := GetOrDie().String(common.ParamToken)
	cfgPath := c.Path(common.ParamConfig)
	defCfg := &config.AppConfig{
		Token: common.DefaultToken,
	}
	if err := config.Load(cfgPath); err != nil {
		if token != "" {
			defCfg.Token = token
		}
		if err := config.Save(cfgPath, defCfg); err != nil {
			return err
		}
		config.Set(defCfg)
	}
	if token == "" {
		c.Set(common.ParamToken, config.Config().Token)
	}
	return nil
}

func GetStringParam(c *cli.Context, param string) string {
	return strings.Trim(c.String(param), " \t\n")
}
