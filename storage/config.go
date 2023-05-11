package storage

type Config struct {
	DBname   string `toml:"dbname"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	SSLmode  string `toml:"sslmode"`
}

func NewConfig() *Config {
	return &Config{}
}
