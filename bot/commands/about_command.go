package commands

import tele "gopkg.in/telebot.v3"

type AboutCommand struct {
}

func (command *AboutCommand) Name() string {
	return "/about"
}

func (command *AboutCommand) Description() string {
	return "about the bot"
}

func (command *AboutCommand) Handle(ctx tele.Context) error {
	err := ctx.Notify(tele.Typing)
	if err != nil {
		return err
	}

	message := "about the bot"
	return ctx.Reply(message, tele.ModeMarkdown)
}
