package foundation

import (
	"github.com/goravel/framework/support/file"
	"github.com/philippgille/chromem-go"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"

	"github.com/goravel-ecosystem/telegram-bot/contracts/command"
	"github.com/goravel-ecosystem/telegram-bot/support"
)

var App *Application

type Application struct {
	bot        *tele.Bot
	config     *viper.Viper
	openai     *openai.Client
	collection *chromem.Collection
}

func init() {
	app := &Application{
		config: viper.New(),
	}

	// Enable automatic environment variables
	app.config.AutomaticEnv()

	envPath := support.EnvPath
	if file.Exists(envPath) {
		app.config.SetConfigType("env")
		app.config.SetConfigFile(envPath)

		if err := app.config.ReadInConfig(); err != nil {
			panic(err)
		}
	}

	App = app
}

// NewApplication creates a new application instance.
func NewApplication() *Application {
	return App
}

// Bot returns the telegram bot instance.
func (app *Application) Bot() *tele.Bot {
	return app.bot
}

// Config returns the viper instance.
func (app *Application) Config() *viper.Viper {
	return app.config
}

// SetBot sets the telegram bot instance.
func (app *Application) SetBot(bot *tele.Bot) {
	app.bot = bot
}

// SetOpenAIClient sets the openai client instance.
func (app *Application) SetOpenAIClient(client *openai.Client) {
	app.openai = client
}

// SetCollection sets the chromem collection instance.
func (app *Application) SetCollection(collection *chromem.Collection) {
	app.collection = collection
}

func (app *Application) RegisterCommands(commands []command.Command) {
	var cmds []tele.Command
	for _, cmd := range commands {
		cmds = append(cmds, tele.Command{
			Text:        cmd.Name(),
			Description: cmd.Description(),
		})

		// register the handler
		app.bot.Handle(cmd.Name(), cmd.Handle)
	}

	if err := app.bot.SetCommands(cmds); err != nil {
		panic(err)
	}
}

func Config() *viper.Viper {
	return App.Config()
}

func Bot() *tele.Bot {
	return App.Bot()
}

func OpenAIClient() *openai.Client {
	return App.openai
}

func Collection() *chromem.Collection {
	return App.collection
}
