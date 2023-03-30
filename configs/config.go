package configs

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig() (*Config, error) {
	var c Config
	var filePath string

	switch os.Getenv("APP_ENV") {
	case "production":
		filePath = "./files/config.production.yaml"
	default:
		filePath = "./files/config.development.yaml"
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &c, err
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return &c, err
	}

	return &c, err
}
