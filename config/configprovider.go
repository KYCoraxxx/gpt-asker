package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	APIKey   string `yaml:"APIKey"`
	UserName string `yaml:"UserName"`
}

type ConfigProvider struct {
	cfg *Config
}

var configProvider *ConfigProvider

func getConfigProvider() *ConfigProvider {
	if configProvider == nil {
		configProvider = &ConfigProvider{}
		configProvider.cfg = new(Config)
	}
	return configProvider
}

func (p *ConfigProvider) readConfig() *Config {

	content, err := ioutil.ReadFile("./application.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if err = yaml.Unmarshal(content, p.cfg); err != nil {
		log.Fatal(err)
	}

	return p.cfg
}

func GetConfig() *Config {
	provider := getConfigProvider()
	return provider.readConfig()
}
