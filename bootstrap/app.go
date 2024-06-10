package bootstrap

import (
	"time"

	tele "gopkg.in/telebot.v3"

	"github.com/goravel-ecosystem/telegram-bot/config"
	"github.com/goravel-ecosystem/telegram-bot/console"
	"github.com/goravel-ecosystem/telegram-bot/foundation"
	"github.com/goravel-ecosystem/telegram-bot/routes"
)

func Boot() {
	app := foundation.NewApplication()

	config.Boot()

	// Set the bot instance
	app.SetBot(createBot())

	// Register the commands
	kernel := &console.Kernel{}
	app.RegisterCommands(kernel.Commands())

	// Set Custom Handler
	routes.Bot()
}

func createBot() *tele.Bot {
	con := foundation.Config()

	bot, err := tele.NewBot(tele.Settings{
		Token:  con.GetString("bot.token"),
		Poller: &tele.LongPoller{Timeout: time.Duration(con.GetInt("bot.poller.timeout")) * time.Second},
	})

	if err != nil {
		panic(err)
	}

	return bot
}
