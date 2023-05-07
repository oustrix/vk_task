package app

import (
	"log"
	"vk_bot/config"
	"vk_bot/pkg/botapi"
)

type App struct {
	cfg *config.Config
	bot *botapi.Bot
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

	a.bot = botapi.NewBot(a.cfg.Bot.Token, a.cfg.Bot.GroupID)

	err := a.bot.InitSession()
	if err != nil {
		log.Println(err)
	}

	updates := a.bot.GetUpdatesChan()
	for update := range updates {
		log.Printf("%+v\n", update)
	}

}
