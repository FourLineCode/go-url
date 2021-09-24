package config

type Config struct {
	Port          string
	DBUrl         string
	OriginAllowed string
}

func GetConfig() Config {
	return Config{Port: ":5000", DBUrl: "test.sqlite", OriginAllowed: "http://localhost:3000"}
}
