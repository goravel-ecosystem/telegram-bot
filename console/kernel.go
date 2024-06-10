package console

import (
	"github.com/goravel-ecosystem/telegram-bot/console/commands"
	"github.com/goravel-ecosystem/telegram-bot/contracts/console"
)

type Kernel struct {
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{
		&commands.HelpCommand{},
	}
}
