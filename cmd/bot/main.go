package main

import (
	"log"
	"vk_bot/config"
	"vk_bot/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error while reading config: %v\n", err)
	}

	a := app.NewApp(cfg)
	a.Start()
}
