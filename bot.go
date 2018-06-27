package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
)

// https://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api#Message

func main() {
	bot, err := tgbotapi.NewBotAPI("lol:plz")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)


		switch update.Message.Text {
			case "hello":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "well, hello")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
		}

	}
}
