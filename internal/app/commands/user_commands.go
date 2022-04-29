package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	welcomeMessage    string = "Добро пожаловать! \nДля помощи введите команду help"
	listCommandToUser string = "Команды пользователя:\n" +
		"chekIp [ip] - Проверка сайта по ip\n" +
		"userHistory - История уникальных пользовательских запросов по его telegram id\n" +
		"Command format: /{command} {command argument}"
	listCommandToAdmin string = "Команды администратора:\n" +
		"addAdmin [telegram id]- Добавление пользователю роли администратора по его telegram id\n" +
		"deleteAdmin [telegram id]- Удаление у пользователя роли администратора по его telegram id\n" +
		"userHistoryAdm [telegram id] - Вывод всех айпи что проверял пользователь по его telegram id" +
		"messageUsers [telegram id] - Отправить сообщение всем пользователям бота"
)

func (c *Commanders) Help(inputMessage *tgbotapi.Message) {
	outputMsgText := listCommandToUser

	if ok, _ := c.bot.DB.HasAdministratorRols(inputMessage.From.ID); ok {
		outputMsgText += "\n\n" + listCommandToAdmin
	}

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
