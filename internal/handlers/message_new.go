package handlers

import (
	"github.com/mitchellh/mapstructure"
	"log"
	"vk_task/pkg/botapi"
)

func (h *Handler) handleMessageNew(update botapi.Update) {
	var message botapi.NewMessage
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &message,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	err := decoder.Decode(update.Object)
	if err != nil {
		log.Println(err)
		return
	}

	userID := message.Message.FromID
	h.services.User.Greeting(userID)
}
