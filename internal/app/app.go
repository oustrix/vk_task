package app

import (
	"log"
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
	if len(a.cfg.Bot.Token) == 0 {
		log.Fatalf("bot token is empty")
	}

	a.bot = bot.NewBot(a.cfg.Bot.Token, a.cfg.Bot.GroupID)

	err := a.bot.Polling()
	if err != nil {
		log.Println(err)
	}

}
