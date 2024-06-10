package main

import (
	"github.com/goravel-ecosystem/telegram-bot/bootstrap"
	"github.com/goravel-ecosystem/telegram-bot/foundation"
)

func main() {
	bootstrap.Boot()

	bot := foundation.Bot()

	bot.Start()
}
