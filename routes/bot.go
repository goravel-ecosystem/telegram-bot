package routes

import (
	tele "gopkg.in/telebot.v3"

	"github.com/goravel-ecosystem/telegram-bot/bot/handlers"
	"github.com/goravel-ecosystem/telegram-bot/foundation"
)

func Bot() {
	bot := foundation.Bot()

	bot.Handle(tele.OnAddedToGroup, handlers.NewAddedToGroupHandler().Handle)
}
