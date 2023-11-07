package storage

type Config struct {
	DBname   string `toml:"dbname" env:"dbname" env-default:"postgres"`
	Port     string `toml:"port" env:"port" env-default:"5432"`
	Host     string `toml:"host" env:"host" env-default:"localhost"`
	User     string `toml:"user" env:"user" env-default:"postgres"`
	Password string `toml:"password" env:"password" env-default:"postgres"`
	Sslmode  string `toml:"sslmode" env:"sslmode" env-default:"disable"`
}

func NewConfig() *Config {
	return &Config{}
}
