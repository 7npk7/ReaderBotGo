package main

import (
	"ReaderBotGo/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

}

func mustToken() string {
	token := flag.String("token", "", "token-tg-bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}