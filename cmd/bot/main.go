package main

import (
	"log"
	"vk_task/config"
	"vk_task/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error while reading config: %v\n", err)
	}

	a := app.NewApp(cfg)
	a.Start()
}
