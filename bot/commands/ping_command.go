package commands

import tele "gopkg.in/telebot.v3"

type PingCommand struct {
}

func (command *PingCommand) Name() string {
	return "/ping"
}

func (command *PingCommand) Description() string {
	return "pong!"
}

func (command *PingCommand) Handle(ctx tele.Context) error {
	err := ctx.Notify(tele.Typing)
	if err != nil {
		return err
	}

	return ctx.Reply("pong!")
}
