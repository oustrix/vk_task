package app

import (
	"vk_bot/config"
	"vk_bot/internal/bot"
)

type App struct {
	cfg *config.Config
	bot *bot.Bot
}

func NewApp(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Start() {

	a.bot = bot.NewBot(a.cfg.Bot.Token)

}
