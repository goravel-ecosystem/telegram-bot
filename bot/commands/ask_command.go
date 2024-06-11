package commands

import (
	tele "gopkg.in/telebot.v3"
)

type AskCommand struct {
}

func (command *AskCommand) Name() string {
	return "/ask"
}

func (command *AskCommand) Description() string {
	return "ask a question"
}

func (command *AskCommand) Handle(ctx tele.Context) error {
	err := ctx.Notify(tele.Typing)
	if err != nil {
		return err
	}

	question := ctx.Message().Payload

	message := "You asked: " + question
	return ctx.Reply(message, tele.ModeMarkdown)
}
