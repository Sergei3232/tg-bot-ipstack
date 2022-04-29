package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *Commanders) AddAdmin(inputMessage *tgbotapi.Message) {
	var outputMsgText string
	arguments := inputMessage.CommandArguments()

	argUserTelegram, err := strconv.Atoi(arguments)
	if err != nil {
		outputMsgText = "Ошибка передачи параметра!"
	} else {
		if err := c.bot.DB.AddAdmin(argUserTelegram, inputMessage.From.ID); err != nil {
			outputMsgText = err.Error()
		} else {
			outputMsgText = "Команда успешно выполнена!"
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, errSend := c.bot.Send(msg)
	if errSend != nil {
		log.Printf("Commander.Help: error sending reply message to chat - %v", err)
	}
}
