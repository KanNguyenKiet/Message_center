package database

import (
	"fmt"
	"net/url"
)

type DBConfig interface {
	String() string
	DSN() string
}

type Config struct {
	Host     string `json:"host" mapstructure:"host" yaml:"host"`
	Database string `json:"database" mapstructure:"database" yaml:"database"`
	Port     int32  `json:"port" mapstructure:"port" yaml:"port"`
	Username string `json:"username" mapstructure:"username" yaml:"username"`
	Password string `json:"password" mapstucture:"password" yaml:"password"`
	Option   string `json:"option" mapstructure:"option" yaml:"option"`
}

// DSN return Data source name for database
func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		c.Username,
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.Database,
		c.Option,
	)
}

// MySQLConfig to config for mysql database
type MySQLConfig struct {
	Config `mapstructure:",squash"`
}

func (c MySQLConfig) String() string {
	return fmt.Sprintf("mysql://%s", c.DSN())
}

func MySQLDefaultConfig() MySQLConfig {
	return MySQLConfig{Config{
		Host:     "localhost",
		Port:     3306,
		Database: "test",
		Username: "root",
		Password: "secret",
		Option:   "?parseTime=true",
	}}
}
