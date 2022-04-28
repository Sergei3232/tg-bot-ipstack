package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *Commanders) Help(inputMessage *tgbotapi.Message) {
	outputMsgText := "Это команда хелп"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Commander.Help: error sending reply message to chat - %v", err)
	}
}
