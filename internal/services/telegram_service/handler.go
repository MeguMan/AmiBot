package telegram_service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

func (s *service) Handle() error {
	for update := range s.updates {
		if update.Message != nil {
			switch {
			case update.Message.Text == "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Choose action")
				msg.ReplyMarkup = mp["start"]
				if _, err := s.bot.Send(msg); err != nil {
					fmt.Println(err)
					continue
				}
			case update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text == "Write 2 or more accounts and their transactions":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "AHHAHA")
				if _, err := s.bot.Send(msg); err != nil {
					fmt.Println(err)
					continue
				}
			default:
				continue
			}
		} else if update.CallbackQuery != nil {
			msg := getMessage(update.CallbackQuery)
			if _, err := s.bot.Send(msg); err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	return nil
}

func getMessage(callback *tgbotapi.CallbackQuery) tgbotapi.Chattable {
	newMarkup, ok := mp[callback.Data]
	if ok {
		return tgbotapi.NewEditMessageReplyMarkup(callback.Message.Chat.ID, callback.Message.MessageID, newMarkup)
	}

	return tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
}
