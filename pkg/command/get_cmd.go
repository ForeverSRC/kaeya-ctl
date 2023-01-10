package command

import (
	"context"
	"fmt"

	"github.com/ForeverSRC/kaeya-ctl/pkg/client"
)

const (
	getCmd cmdText = "get"
)

type GetCommand struct {
	commandHead
	getter client.Getter
}

func NewGetCommand(getter client.Getter) *GetCommand {
	return &GetCommand{
		commandHead: commandHead{
			cmd:     getCmd,
			argNums: 1,
		},
		getter: getter,
	}
}

func (g *GetCommand) Run(ctx context.Context, args []string) (string, error) {
	if err := validateArgNums(g.cmd, args, g.argNums); err != nil {
		return "", err
	}
	key := args[0]
	kv, err := g.getter.Get(ctx, key)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s -> %s", kv.Key, kv.Value), nil
}

func (g *GetCommand) Help() string {
	return "get\nGet value of the key. Usage: get KEY"
}
