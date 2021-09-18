package config

type Config struct {
	Port string
}

func GetConfig() Config {
	return Config{Port: ":5000"}
}
