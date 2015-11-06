package main

import (
	"io/ioutil"
	"log"
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

func loadConfig(configPath string) Config {
	var config Config

	_, err := os.Stat(configPath)
	if err != nil {
		log.Fatal("ERROR: ", err.Error())
	}

	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("ERROR: ", err.Error())
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		log.Fatal("ERROR: ", err.Error())
	}
	return config
}
