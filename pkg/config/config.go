package config

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/slim-ai/gitadm/pkg/common"
	"gopkg.in/yaml.v2"
)

// global config
var cfg *AppConfig

type AppConfig struct {
	Token string `yaml:"token"`
}

func Config() *AppConfig {
	if cfg == nil {
		panic("no configuration file loaded - coding error")
	}
	return cfg
}

// NewConfig returns a new decoded Config struct
func Load(fullpath string) error {
	cfg = &AppConfig{}
	file, err := os.Open(fullpath)
	if err != nil {
		return err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(cfg); err != nil {
		return err
	}
	return nil
}

func Set(c *AppConfig) {
	cfg = c
}

func Save(fullpath string, cfg interface{}) error {
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(fullpath), 0700) // Create your file
	}
	// open a file
	f, err := os.Create(fullpath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}

func homeDir() string {
	currentUser, err := user.Current()
	if err != nil {
		return os.Getenv(common.EnvVarUserHomeDir)
	}
	return currentUser.HomeDir
}

func Directory(appname string) (configDir string) {
	if cfgHome := os.Getenv(common.EnvVarConfigHomeDir); cfgHome != "" {
		// $XDG_CONFIG_HOME/appName
		configDir = filepath.Join(cfgHome, appname)
	} else {
		configDir = filepath.Join(homeDir(), ".config", appname)
	}
	return
}

func FullFilePath(appname string) (filename string) {
	filename = filepath.Join(Directory(appname), "config")
	return
}
