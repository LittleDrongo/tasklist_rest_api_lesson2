package api

const (
	ConfigFilePath = "configs/api.toml"
)

type config struct {
	Host         string `toml:"host"`
	Port         string `toml:"port"`
	LoggerLevel  string `toml:"logger_level"`
	DataBasePath string `toml:"database_path"`
}

func DefaultConfig() *config {
	return &config{
		Host:         "http://localhost",
		Port:         ":8080",
		LoggerLevel:  "debug",
		DataBasePath: "data/db/tasklist.db",
	}
}
