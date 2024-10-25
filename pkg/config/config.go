package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	User     string `yaml:"container" env:"USER" env-default:"container"`
	Password string `yaml:"password" env:"PASSWORD" env-default:"1234"`
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Name     string `yaml:"name" env:"NAME" env-default:"postgres"`
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &cfg
}
