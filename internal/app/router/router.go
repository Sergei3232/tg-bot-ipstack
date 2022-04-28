package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"runtime/debug"
)

type Router struct {
	bot *tgbotapi.BotAPI
}

func NewRouter(bot *tgbotapi.BotAPI) *Router {
	return &Router{bot}
}

func (c *Router) HandleUpdate(update tgbotapi.Update) {
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

func (c *Router) handleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		c.showCommandFormat(msg)

		return
	}

	switch msg.Command() {
	case "start":

	default:
		log.Printf("Unknown command - %s", msg.Command())
	}
}

func (c *Router) showCommandFormat(inputMessage *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command format: /{command} {command argument}")

	_, err := c.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.showCommandFormat: error sending reply message to chat - %v", err)
	}
}
