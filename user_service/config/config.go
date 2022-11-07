package config

import (
	"github.com/spf13/viper"
)

type Database struct {
	Port         string
	Host         string
	DatabaseName string
}

type ServerConfig struct {
	Env      string
	Port     string
	Host     string
	Database Database
}

func DefaultLoadConfig() *ServerConfig {
	return &ServerConfig{
		Env:  "local",
		Port: "9000",
		Host: "localhost",
		Database: Database{
			Port:         "8000",
			Host:         "localhost",
			DatabaseName: "message_server",
		},
	}
}

func LoadConfig() (*ServerConfig, error) {
	cfg := DefaultLoadConfig()
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
