package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		Bot
	}

	Bot struct {
		Token string `env:"BOT_TOKEN" env-required:"true"`
	}
)

func NewConfig() (*Config, error) {
	var cfg Config

	// Read config from .yml file (basic settings)
	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		return nil, err
	}

	// Read security-sensitive data from .env file or environment variables and override settings from .yml file
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
