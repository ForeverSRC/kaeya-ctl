package command

import (
	"context"
)

const (
	exitCmd cmdText = "exit"
)

type ExitCommand struct {
	commandHead
}

func NewExitCommand() *ExitCommand {
	return &ExitCommand{
		commandHead: commandHead{
			cmd:     exitCmd,
			argNums: 0,
		},
	}
}

func (g *ExitCommand) Run(ctx context.Context, args []string) (string, error) {
	return "", ErrExit
}

func (g *ExitCommand) Help() string {
	return "exit\nExit interactive mode. Usage: exit"
}
