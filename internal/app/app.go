package app

import (
	"log"
	"vk_bot/config"
	"vk_bot/internal/handlers"
	"vk_bot/internal/service"
	"vk_bot/pkg/botapi"
)

type App struct {
	cfg      *config.Config
	bot      *botapi.Bot
	services *service.Services
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

	userService := service.NewUserService(a.bot)

	a.services = &service.Services{
		User: userService,
	}

	handler := handlers.NewHandler(a.services, updates)
	handler.Init()
}
