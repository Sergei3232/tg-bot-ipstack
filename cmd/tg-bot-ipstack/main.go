package main

import (
	"fmt"
	"github.com/Sergei3232/tg-bot-ipstack/internal/app/db"
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

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
	fmt.Println(configs, bot, db)
}
