package commands

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"regexp"
	"time"
)

const (
	welcomeMessage    string = "Добро пожаловать! \nДля помощи введите команду help"
	listCommandToUser string = "Команды пользователя:\n" +
		"chekIp [0.0.0.0] - Проверка сайта по ip\n" +
		"userHistory - История уникальных пользовательских запросов по его telegram id\n" +
		"Command format: /{command} {command argument}"
	listCommandToAdmin string = "Команды администратора:\n" +
		"addAdmin [telegram id]- Добавление пользователю роли администратора по его telegram id\n" +
		"deleteAdmin [telegram id]- Удаление у пользователя роли администратора по его telegram id\n" +
		"userHistoryAdm [telegram id] - Вывод всех айпи что проверял пользователь по его telegram id" +
		"messageToUsers [telegram id] - Отправить сообщение всем пользователям бота"
	ipRegExp string = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
)

func (c *Commanders) Start(inputMessage *tgbotapi.Message) {
	outputMsgText := welcomeMessage

	if err := c.bot.DB.AddNewUserBot(inputMessage.From.ID, inputMessage.From.UserName); err != nil {
		log.Println(err.Error())
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Commander.Help: error sending reply message to chat - %v", err)
	}
}

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

func (c *Commanders) ChekIp(inputMessage *tgbotapi.Message) {
	var outputMsgText string
	ip := inputMessage.CommandArguments()

	ok, err := regexp.MatchString(ipRegExp, ip)
	if err != nil {
		log.Panicln(err.Error())
	}

	if !ok {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
			"Ошибка передачи параметра! IP должен быть вида [000.000.000.000]")

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Commander.ChekIp: error sending reply message to chat - %v", err)
		}
		log.Panicln(errors.New("error. IP is not valid"))
		return
	}

	user, err := c.bot.DB.GetUserTelegram(inputMessage.From.ID)
	if err != nil {
		log.Println(err)
	}
	textResult := "Test"
	timeQuery := time.Now()

	c.bot.DB.AddUserHistoryQuery(user.Id, ip, textResult, timeQuery)

	log.Println(outputMsgText)
}

func (c *Commanders) validationIpAdress(ip string) bool {
	re, _ := regexp.Compile(ipRegExp)
	if re.MatchString(ip) {
		return true
	} else {
		return false
	}
}
