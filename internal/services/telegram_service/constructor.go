package telegram_service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Service interface {
	Handle() error
}

type service struct {
	updates tgbotapi.UpdatesChannel
	bot     *tgbotapi.BotAPI
}

func New(telegramBotToken string) (Service, error) {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		return nil, err
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}

	return &service{
		updates: updates,
		bot:     bot,
	}, nil
}
