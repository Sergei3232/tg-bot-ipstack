package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const yamlFile = "./config.yaml"

type Config struct {
	DnsDB            string `yaml:"bot_db"`
	TokenTelegramBot string `yaml:"token_telegram"`
}

func NenConfig() (*Config, error) {
	var c Config
	yamlFile, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
