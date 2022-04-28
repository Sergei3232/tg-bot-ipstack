package main

import (
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/bot"
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

	tgClient := bot.NewBotTgClient(configs)
	tgClient.Bot.Debug = true

	log.Printf("Authorized on account %s", tgClient.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	routerHandler := routerTg.NewRouter(tgClient.Bot)

	updates, err := tgClient.Bot.GetUpdatesChan(u)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
