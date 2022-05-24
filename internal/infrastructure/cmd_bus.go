package infrastructure

import (
	"fmt"

	"github.com/MAZEN-Kenjrawi/pwd/internal/app/command"
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/handler"
)

type CmdBus struct {
	deleteLoginHandler  handler.DeleteLoginHandler
	replaceLoginHandler handler.ReplaceLoginHandler
	secureLoginHandler  handler.SecureLoginHandler
}

func (b *CmdBus) Handle(c interface{}) error {
	cmdType := getType(c)
	switch cmdType {
	case "command.DeleteLogin":
		return b.deleteLoginHandler.Handle(c.(command.DeleteLogin))
	case "command.ReplaceLogin":
		return b.replaceLoginHandler.Handle(c.(command.ReplaceLogin))
	case "command.SecureLogin":
		return b.secureLoginHandler.Handle(c.(command.SecureLogin))
	}

	return fmt.Errorf("no handler found for command %s", cmdType)
}
