package commands

import (
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/bot"
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/ipstack"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander interface {
	Help(inputMessage *tgbotapi.Message)
	Start(inputMessage *tgbotapi.Message)
	AddAdmin(inputMessage *tgbotapi.Message)
	DeleteAdmin(inputMessage *tgbotapi.Message)
	MessageUsers(inputMessage *tgbotapi.Message)
	ChekIp(inputMessage *tgbotapi.Message)
	GetHistoryUserQuery(inputMessage *tgbotapi.Message)
	GetHistoryUserQueryAdmin(inputMessage *tgbotapi.Message)
}

type Commanders struct {
	bot      *bot.TgBot
	clietnIp ipstack.QueryIP
}

func NewDemoCommander(bot *bot.TgBot, clietnIp ipstack.QueryIP) Commander {
	return &Commanders{
		bot:      bot,
		clietnIp: clietnIp,
	}
}
