package bot

import (
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/db"
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/ipstack"
	"github.com/Sergei3232/tg-bot-ipstack/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type TgBot struct {
	DB db.Repository
	*tgbotapi.BotAPI
	CliettIP ipstack.QueryIP
}

func NewBotTgClient(config *config.Config) *TgBot {

	dbClient, errDb := db.NewDbConnectClient(config.DnsDB)
	if errDb != nil {
		log.Panic(errDb)
	}

	bot, err := tgbotapi.NewBotAPI(config.TokenTelegramBot)
	if err != nil {
		log.Panic(err)
	}

	clientIp := ipstack.NewClientIP(config.HostNameIp, config.AccessKey)

	return &TgBot{dbClient, bot, clientIp}
}
