package config

import (
	"github.com/spf13/viper"
	"message-server/common/database"
)

type ServerConfig struct {
	Env           string            `json:"env" mapstructure:"env" yaml:"env"`
	Host          string            `json:"host" mapstructure:"host" yaml:"host"`
	HttpPort      string            `json:"httpPort" mapstructure:"httpPort" yaml:"httpPort"`
	GRPCPort      string            `json:"GRPCPort" mapstructure:"GRPCPort" yaml:"GRPCPort"`
	Database      database.DBConfig `json:"database" mapstructure:"database" yaml:"database"`
	MigrateFolder string            `json:"migrateFolder" mapstructure:"migreateFolder" yaml:"migrateFolder"`
}

func DefaultLoadConfig() *ServerConfig {
	return &ServerConfig{
		Env:           "local",
		HttpPort:      "9080",
		GRPCPort:      "8080",
		Host:          "localhost",
		Database:      database.MySQLDefaultConfig(),
		MigrateFolder: "file://sql/migration",
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
