// package config

// type Configuration struct {
// 	ServerPort          int
// 	DBMaxConnection     int
// 	DBMaxIdleConnection int
// 	DBPath              string
// 	EventServer         string
// 	EventServerPort     int
// }

// func LoadConfiguration() *Configuration {

// 	cfg := &Configuration{
// 		ServerPort:          8081,
// 		DBMaxConnection:     10,
// 		DBMaxIdleConnection: 5,
// 		DBPath:              "C:\\Users\\Kokumo\\Documents\\Uala\\pajarit-feed-service\\pajarit.db",
// 		EventServer:         "nats://localhost",
// 		EventServerPort:     4222,
// 	}

// 	return cfg
// }

package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Port int `yaml:"port"`
}

type Database struct {
	Path              string `yaml:"path"`
	MaxConnection     int    `yaml:"maxConnection"`
	MaxIdleConnection int    `yaml:"maxIdleConnection"`
}

type Event struct {
	Server string `yaml:"serverUrl"`
	Port   int    `yaml:"port"`
}

type Configuration struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Event    Event    `yaml:"event"`
}

func LoadConfiguration() (*Configuration, error) {
	cfg, err := loadYaml()
	if err != nil {
		return nil, err
	}

	if value := os.Getenv("EVENT_SERVER_URL"); value != "" {
		cfg.Event.Server = value
	}

	if value := os.Getenv("DB_PATH"); value != "" {
		cfg.Database.Path = value
	}

	return cfg, nil
}

func loadYaml() (*Configuration, error) {
	cfg := &Configuration{}

	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
