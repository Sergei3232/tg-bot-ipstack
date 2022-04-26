package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const yamlFile = "./config.yaml"

type Config struct {
	DnsDB            string `yaml:"bot_db"`
	TokenTelegramBot string `yaml:"token_telegram"`
}

func NenConfig() *Config {
	var c Config
	yamlFile, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &c
}
