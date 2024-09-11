package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var mp = map[string]tgbotapi.InlineKeyboardMarkup{
	"start": tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Find repeatable signers on different accounts",
				"Write 2 or more accounts and their transactions",
			),
		),
	),
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text != "/start" {
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			switch update.Message.Text {
			case "/start":
				msg.ReplyMarkup = mp["start"]
			}

			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.AnswerCallbackQuery(callback); err != nil {
				panic(err)
			}

			msg := getMessage(update.CallbackQuery)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}

func getMessage(callback *tgbotapi.CallbackQuery) tgbotapi.Chattable {
	newMarkup, ok := mp[callback.Data]
	if ok {
		return tgbotapi.NewEditMessageReplyMarkup(callback.Message.Chat.ID, callback.Message.MessageID, newMarkup)
	}

	return tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
}
