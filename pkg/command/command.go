package command

import (
	"context"
	"fmt"

	"github.com/ForeverSRC/kaeya-ctl/pkg/client"
)

const (
	delim = " "
)

type cmdText string

type Command interface {
	Run(ctx context.Context, args []string) (string, error)
	Help() string
}

type commandHead struct {
	cmd     cmdText
	argNums int
}

func newCommandMappings(client client.KaeyaClient) map[cmdText]Command {
	mapping := map[cmdText]Command{
		setCmd:  NewSetCommand(client),
		getCmd:  NewGetCommand(client),
		exitCmd: NewExitCommand(),
	}

	return mapping
}

func validateArgNums(cmd cmdText, args []string, expectedNum int) error {
	if len(args) != expectedNum {
		return fmt.Errorf("%s command needs %d args, got %d", cmd, expectedNum, len(args))
	}

	return nil
}
