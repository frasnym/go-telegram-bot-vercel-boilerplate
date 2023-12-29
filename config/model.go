package config

// Configuration struct to hold environment variables
type Config struct {
	VercelUrl        string `env:"VERCEL_URL"`
	Port             string `env:"PORT"`
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN"`
}
