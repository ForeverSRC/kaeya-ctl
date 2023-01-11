package command

import (
	"context"

	"github.com/ForeverSRC/kaeya-ctl/pkg/client"
	"github.com/ForeverSRC/kaeya-ctl/pkg/domain"
)

const (
	setCmd cmdText = "set"
)

type SetCommand struct {
	commandHead
	setter client.Setter
}

func NewSetCommand(setter client.Setter) *SetCommand {
	return &SetCommand{
		commandHead: commandHead{
			cmd:     setCmd,
			argNums: 2,
		},
		setter: setter,
	}
}

func (s *SetCommand) Run(ctx context.Context, args []string) (string, error) {
	if err := validateArgNums(s.cmd, args, s.argNums); err != nil {
		return "", err
	}

	key := args[0]
	value := args[1]

	err := s.setter.Set(ctx, domain.KV{
		Key:   key,
		Value: value,
	})

	if err != nil {
		return "", err
	}

	return "success", nil
}

func (s *SetCommand) Help() string {
	return "set\nSet kv pair. Usage: set KEY VALUE"
}
