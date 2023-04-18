package api

type Config struct {
	Port string `toml: "port"`
}

func NewConfig() *Config {
	return &Config{
		Port: "8080",
	}
}
