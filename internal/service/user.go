package service

import (
	"log"
	"vk_bot/internal/keyboards"
	"vk_bot/pkg/botapi"
)

type UserService struct {
	bot *botapi.Bot
}

const (
	greetingMessage = "Здравствуйте! Надеюсь, я Вам понравился)"
)

func NewUserService(bot *botapi.Bot) *UserService {
	return &UserService{
		bot: bot,
	}
}

func (u *UserService) Greeting(userID int) {
	msg := botapi.MessageConfig{
		UserID:   userID,
		RandomID: 0,
		Message:  greetingMessage,
		Keyboard: keyboards.NewGreetingKeyboard(),
	}

	err := u.bot.SendMessage(&msg)

	if err != nil {
		log.Println(err)
	}
}
