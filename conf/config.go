package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type MainConfig struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func DatabaseDSN() (string, error) {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return "", err
	}

	var config MainConfig
	err = yaml.Unmarshal(yamlFile, &config)
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
