package command

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ForeverSRC/kaeya-ctl/pkg/client"
)

const (
	helpCmd cmdText = "help"
)

var (
	errEmptyInput = errors.New("empty input")
	ErrExit       = errors.New("exit")
)

type Dispatcher interface {
	Dispatch(ctx context.Context, input string) (string, error)
	Helps() string
}

type DefaultDispatcher struct {
	cmdMapping map[cmdText]Command
}

func NewDefaultDispatcher(client client.KaeyaClient) *DefaultDispatcher {
	mapping := newCommandMappings(client)

	return &DefaultDispatcher{
		cmdMapping: mapping,
	}
}

func (d *DefaultDispatcher) Dispatch(ctx context.Context, input string) (string, error) {
	if input == "" {
		return "", errEmptyInput
	}

	cmdTxt, args := d.parseCmdText(input)

	switch cmdTxt {
	case helpCmd:
		return d.doHelp(args)
	default:
		cmd, ok := d.cmdMapping[cmdTxt]
		if !ok {
			return "", fmt.Errorf("invalid command: %s", cmdTxt)
		}

		return cmd.Run(ctx, args)
	}

}

func (d *DefaultDispatcher) parseCmdText(input string) (cmdText, []string) {
	strs := strings.Split(input, delim)
	return cmdText(strs[0]), strs[1:]
}

func (d *DefaultDispatcher) doHelp(args []string) (string, error) {
	if err := validateArgNums(helpCmd, args, 1); err != nil {
		return "", err
	}

	needHelpCmd := args[0]

	cmd, ok := d.cmdMapping[cmdText(needHelpCmd)]
	if !ok {
		return "", fmt.Errorf("invalid command: %s", needHelpCmd)
	}

	return cmd.Help(), nil

}

func (d *DefaultDispatcher) Helps() string {
	buf := strings.Builder{}
	for _, v := range d.cmdMapping {
		buf.WriteString(v.Help() + "\n")
	}

	return buf.String()

}
