package main

import (
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Target       string
	Organization map[string]string
}

func (self *Config) mergeCliArgs(args cli.Args) {
	self.Target = args.First()
}

func loadConfig(configPath string) (*Config, error) {
	var config Config

	_, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
