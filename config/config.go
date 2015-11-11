package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Organization
}

type Organization struct {
	Name string
}

func Load(configPath string) (*Config, error) {
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
