package web

type Config struct {
	Port string `env:"port"`
}

func DefaultConfig() *Config {
	return &Config{
		Port: "8080",
	}
}