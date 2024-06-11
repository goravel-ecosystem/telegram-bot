package commands

import (
	"fmt"

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
	return "list all available commands and their descriptions"
}

// Handle executes the command.
func (command *HelpCommand) Handle(ctx tele.Context) error {
	// notify the user that the bot is typing
	err := ctx.Notify(tele.Typing)
	if err != nil {
		return err
	}

	var text string
	commands, err := ctx.Bot().Commands()
	if err != nil {
		return err
	}

	text += "*Available Commands:*\n\n"

	for _, cmd := range commands {
		if cmd.Text == command.Name() {
			continue
		}
		text += fmt.Sprintf("`%s` - _%s_\n", cmd.Text, cmd.Description)
	}

	return ctx.Send(text, tele.ModeMarkdown)
}
