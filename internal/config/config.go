package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port string `yaml:"port" env:"PORT" env-default:"8080"`
}

func NewConfig(configPath string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		logrus.Error("cannot read the config")
		return nil, err
	}

	return &cfg, nil
}
