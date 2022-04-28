package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Commander struct {
	bot *tgbotapi.BotAPI
}

func NewDemoCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot: bot,
	}
}

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	outputMsgText := "Это команда хелп"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Commander.Help: error sending reply message to chat - %v", err)
	}
}
