package main

import (
	"github.com/MeguMan/AmiBot/internal/services/telegram_service"
	"log"
	"os"
)

func main() {
	telegramService, err := telegram_service.New(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	if err = telegramService.Handle(); err != nil {
		log.Fatal(err)
	}
}
