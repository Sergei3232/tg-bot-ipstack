package main

import (
	"fmt"
	"github.com/Sergei3232/tg-bot-ipstack/internal/config"
)

func main() {
	configs := config.NenConfig()
	fmt.Println(configs.DnsDB)
	fmt.Println(configs.TokenTelegramBot)
}
