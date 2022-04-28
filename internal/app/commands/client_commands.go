package commands

import (
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander interface {
	Help(inputMessage *tgbotapi.Message)
}

type Commanders struct {
	bot *bot.BotTg
}

func NewDemoCommander(bot *bot.BotTg) Commander {
	return &Commanders{
		bot: bot,
	}
}
