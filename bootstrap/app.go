package bootstrap

import (
	"time"

	"github.com/philippgille/chromem-go"
	"github.com/sashabaranov/go-openai"
	tele "gopkg.in/telebot.v3"

	"github.com/goravel-ecosystem/telegram-bot/bot"
	"github.com/goravel-ecosystem/telegram-bot/config"
	"github.com/goravel-ecosystem/telegram-bot/foundation"
	"github.com/goravel-ecosystem/telegram-bot/routes"
)

func Boot() {
	app := foundation.NewApplication()

	config.Boot()

	// Set the bot instance
	app.SetBot(createBot())

	// Set the openai client instance
	app.SetOpenAIClient(createOpenAIClient())

	// Set the collection instance
	app.SetCollection(createCollection())

	// Register the commands
	kernel := &bot.Kernel{}
	app.RegisterCommands(kernel.Commands())

	// Set Custom Handler
	routes.Bot()
}

func createBot() *tele.Bot {
	con := foundation.Config()

	b, err := tele.NewBot(tele.Settings{
		Token:  con.GetString("bot.token"),
		Poller: &tele.LongPoller{Timeout: time.Duration(con.GetInt("bot.poller.timeout")) * time.Second},
	})

	if err != nil {
		panic(err)
	}

	return b
}

func createOpenAIClient() *openai.Client {
	con := foundation.Config()

	client := openai.NewClient(con.GetString("openai.api_key"))

	return client
}

func createCollection() *chromem.Collection {
	con := foundation.Config()
	db, err := chromem.NewPersistentDB("", true)
	if err != nil {
		panic(err)
	}

	return db.GetCollection(con.GetString("database.collection"), chromem.NewEmbeddingFuncOpenAI(con.GetString("openai.api_key"), chromem.EmbeddingModelOpenAI3Small))
}
