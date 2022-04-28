package bot

import (
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/db"
	"github.com/Sergei3232/tg-bot-ipstack/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type BotTg struct {
	Bot *tgbotapi.BotAPI
	DB  db.Repository
}

func NewBotTgClient(config *config.Config) *BotTg {

	db, errDb := db.NewDbConnectClient(config.DnsDB)
	if errDb != nil {
		log.Panic(errDb)
	}

	bot, err := tgbotapi.NewBotAPI(config.TokenTelegramBot)
	if err != nil {
		log.Panic(err)
	}

	return &BotTg{bot, db}
}
