package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const yamlFile = "./config.yaml"

type Config struct {
	DnsDB            string `yaml:"bot_db"`
	TokenTelegramBot string `yaml:"token_telegram"`
	AccessKey        string `yaml:"access_key"`
	HostNameIp       string `yaml:"host_name_ip"`
	BindAddr         string `yaml:"bind_addr"`
	LogLevel         string `yaml:"log_level"`
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
