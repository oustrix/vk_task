package service

import "vk_bot/pkg/botapi"

type GreetingService struct {
	bot *botapi.Bot
}

const (
	greetingMessage = "Здравствуйте! Надеюсь, я Вам понравился)"
)

func NewGreetingService(bot *botapi.Bot) *GreetingService {
	return &GreetingService{
		bot: bot,
	}
}

func (g *GreetingService) Greeting(userID int) error {
	msg := botapi.MessageConfig{
		UserID:   userID,
		RandomID: 0,
		Message:  greetingMessage,
	}

	err := g.bot.SendMessage(&msg)

	return err
}
