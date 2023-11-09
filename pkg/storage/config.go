package storage

type Config struct {
	DBname   string `env:"dbname" env-default:"postgres"`
	Port     string `env:"port" env-default:"5432"`
	Host     string `env:"host" env-default:"localhost"`
	User     string `env:"user" env-default:"postgres"`
	Password string `env:"password" env-default:"postgres"`
	Sslmode  string `env:"sslmode" env-default:"disable"`
}

func NewConfig() *Config {
	return &Config{}
}
