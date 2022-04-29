package commands

import (
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander interface {
	Help(inputMessage *tgbotapi.Message)
	Start(inputMessage *tgbotapi.Message)
	AddAdmin(inputMessage *tgbotapi.Message)
	DeleteAdmin(inputMessage *tgbotapi.Message)
}

type Commanders struct {
	bot *bot.TgBot
}

func NewDemoCommander(bot *bot.TgBot) Commander {
	return &Commanders{
		bot: bot,
	}
}
