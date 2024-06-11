package bot

import (
	"github.com/goravel-ecosystem/telegram-bot/bot/commands"
	commandcontract "github.com/goravel-ecosystem/telegram-bot/contracts/command"
)

type Kernel struct {
}

func (kernel *Kernel) Commands() []commandcontract.Command {
	return []commandcontract.Command{
		&commands.ListCommand{},
		&commands.HelpCommand{},
		&commands.PingCommand{},
	}
}
