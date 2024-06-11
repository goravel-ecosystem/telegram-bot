package command

import tele "gopkg.in/telebot.v3"

// Command is an interface that represents a console command.
type Command interface {
	// Name returns the name of the command.
	Name() string
	// Description returns the description of the command.
	Description() string
	// Handle executes the command.
	Handle(tele.Context) error
}
