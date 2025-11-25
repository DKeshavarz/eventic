package telegram

type Config struct {
	APIKey string `env:"TELEGRAM_API_KEY"`
}

func DefaultConfig() *Config {
	return &Config{}
}