package main

import (
	"fmt"
	"github.com/Sergei3232/tg-bot-ipstack/internal/config"
	"log"
)

func main() {
	configs, err := config.NenConfig()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(configs)
}
