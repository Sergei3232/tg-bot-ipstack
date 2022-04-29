package router

import (
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/bot"
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"runtime/debug"
)

type Router interface {
	HandleUpdate(update tgbotapi.Update)
	handleMessage(msg *tgbotapi.Message)
	showCommandFormat(inputMessage *tgbotapi.Message)
}

type ClientRouter struct {
	bot       *bot.TgBot
	commander commands.Commander
}

func NewRouter(bot *bot.TgBot) Router {
	return &ClientRouter{
		bot,
		commands.NewDemoCommander(bot),
	}

}

func (c *ClientRouter) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	switch {
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *ClientRouter) handleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(msg)
		return
	}

	switch msg.Command() {
	case "start":
		c.commander.Start(msg)
	case "help":
		c.commander.Help(msg)
	case "addAdmin":
		c.commander.AddAdmin(msg)
	case "deleteAdmin":
		c.commander.DeleteAdmin(msg)
	case "messageToUsers":
		c.commander.MessageUsers(msg)
	default:
		log.Printf("Unknown command - %s", msg.Command())
	}
}

func (c *ClientRouter) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command} {command argument}")

	_, err := c.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.showCommandFormat: error sending reply message to chat - %v", err)
	}
}
