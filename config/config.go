package config

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

func DefaultLoad() (*ServerConfig, error) {
	return &ServerConfig{
		Env:  "local",
		Port: "9000",
		Host: "localhost",
		Database: Database{
			Port:         "8000",
			Host:         "localhost",
			DatabaseName: "message_server",
		},
	}, nil
}
