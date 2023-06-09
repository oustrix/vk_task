package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"strings"
)

type (
	Config struct {
		Bot `yaml:"bot"`
	}

	Bot struct {
		Token   string `env:"BOT_TOKEN" env-required:"true"`
		GroupID string `env:"GROUP_ID" yaml:"group_id" env-required:"true"`
	}
)

func NewConfig() (*Config, error) {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Read config from .yml file (basic settings)
	err = cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		return nil, err
	}

	// Read security-sensitive data from .env file or environment variables and override settings from .yml file
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
