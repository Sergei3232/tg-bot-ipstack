package main

import (
	"fmt"
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/db"
	routerTg "github.com/Sergei3232/tg-bot-ipstack/internal/app/router"
	"github.com/Sergei3232/tg-bot-ipstack/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	configs, err := config.NenConfig()
	if err != nil {
		log.Panicln(err)
	}

	db, errDb := db.NewDbConnectClient(configs.DnsDB)
	if errDb != nil {
		log.Panic(errDb)
	}

	bot, err := tgbotapi.NewBotAPI(configs.TokenTelegramBot)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	routerHandler := routerTg.NewRouter(bot)

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
	fmt.Println(configs, bot, db)
}
