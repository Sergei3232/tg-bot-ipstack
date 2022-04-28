package bot

import (
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/db"
	"github.com/Sergei3232/tg-bot-ipstack/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type TgBot struct {
	DB db.Repository
	*tgbotapi.BotAPI
}

func NewBotTgClient(config *config.Config) *TgBot {

	db, errDb := db.NewDbConnectClient(config.DnsDB)
	if errDb != nil {
		log.Panic(errDb)
	}

	bot, err := tgbotapi.NewBotAPI(config.TokenTelegramBot)
	if err != nil {
		log.Panic(err)
	}

	return &TgBot{db, bot}
}
