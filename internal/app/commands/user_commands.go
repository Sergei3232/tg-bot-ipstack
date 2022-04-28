package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	welcomeMessage    string = "Добро пожаловать! \nДля помощи введите команду help"
	listCommandToUser string = "Пользовательские команды:\n" +
		"chekIp - Проверка сайта по ip\n" +
		"user_history - История уникальных пользовательских запросов\n" +
		"Command format: /{command} {command argument}"
)

func (c *Commanders) Help(inputMessage *tgbotapi.Message) {
	outputMsgText := listCommandToUser
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Commander.Help: error sending reply message to chat - %v", err)
	}
}

func (c *Commanders) Start(inputMessage *tgbotapi.Message) {
	outputMsgText := welcomeMessage
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Commander.Help: error sending reply message to chat - %v", err)
	}
}
