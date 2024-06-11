package commands

import (
	tele "gopkg.in/telebot.v3"
)

// HelpCommand is a console command that shows help.
type HelpCommand struct {
}

// Name returns the name of the command.
func (command *HelpCommand) Name() string {
	return "/help"
}

// Description returns the description of the command.
func (command *HelpCommand) Description() string {
	return "help to guide user to use the bot"
}

// Handle executes the command.
func (command *HelpCommand) Handle(ctx tele.Context) error {
	return ctx.Send("This is a help message", tele.ModeMarkdown)
}
