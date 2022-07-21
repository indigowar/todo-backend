package main

import (
	"github.com/indigowar/todo-backend/internal/app"
	"github.com/indigowar/todo-backend/internal/config"
	"log"
)

func main() {
	cfg, err := config.Init("config")
	if err != nil {
		log.Fatal("Application is not configured")
	}

	app.Run(cfg)
}
