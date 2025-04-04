package config

type Configuration struct {
	ServerPort          int
	DBMaxConnection     int
	DBMaxIdleConnection int
	DBPath              string
	EventServer         string
	EventServerPort     int
}

func LoadConfiguration() *Configuration {

	cfg := &Configuration{
		ServerPort:          8081,
		DBMaxConnection:     10,
		DBMaxIdleConnection: 5,
		DBPath:              "C:\\Users\\Kokumo\\Documents\\Uala\\pajarit-feed-service\\pajarit.db",
		EventServer:         "nats://localhost",
		EventServerPort:     4222,
	}

	return cfg
}
