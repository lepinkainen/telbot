package main

import (
	"gopkg.in/telegram-bot-api.v4"
)

// Reply to a message
func (h Handler) Reply(message string) {
	msg := tgbotapi.NewMessage(h.update.Message.Chat.ID, message)
	msg.ReplyToMessageID = h.update.Message.MessageID
	h.bot.Send(msg)
}

// Send message to chat
func (h Handler) Say(message string) {
	msg := tgbotapi.NewMessage(h.update.Message.Chat.ID, message)
	h.bot.Send(msg)
}
