package handlers

import "vk_bot/pkg/botapi"

func (h *Handler) handleMessageNew(message botapi.NewMessage) {
	userID := message.Message.FromID
	h.services.User.Greeting(userID)
}
