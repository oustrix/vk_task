package handlers

import (
	"log"
	"vk_bot/internal/service"
	"vk_bot/pkg/botapi"
)

type Handler struct {
	services *service.Services
	updates  <-chan botapi.Update
}

func NewHandler(services *service.Services, updates <-chan botapi.Update) *Handler {
	return &Handler{
		services: services,
		updates:  updates,
	}
}

func (h *Handler) Init() {
	for update := range h.updates {
		go h.handleUpdate(update)
	}
}

func (h *Handler) handleUpdate(update botapi.Update) {
	switch update.Type {
	case botapi.MessageNew:
		h.handleMessageNew(update)
	default:
		log.Printf("Unsupported update type: %s\n", update.Type)
	}
}
