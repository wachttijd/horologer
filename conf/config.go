package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type MainConfig struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func RetrieveConfig() (MainConfig, error) {
	yamlFile, err := os.ReadFile("config.yml")
	if err != nil {
		return MainConfig{}, err
	}

	var config MainConfig

	err = yaml.Unmarshal(yamlFile, &config)
	
	return config, err
}

func DatabaseDSN() (string, error) {
	config, err := RetrieveConfig()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s@(%s:%d)/%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	), nil
}
