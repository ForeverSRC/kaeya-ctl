package service

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/ForeverSRC/kaeya-ctl/pkg/client"
	"github.com/ForeverSRC/kaeya-ctl/pkg/command"
)

const (
	welcome = "Kaeya-ctl interaction mode."
	bye     = "Kaeya-ctl interaction mode exit."
	prompt  = ">> "
)

type DBInteraction interface {
	Interactive(ctx context.Context, addr string) error
}

type DefaultDBInteraction struct {
	reader     io.Reader
	writer     io.Writer
	dispatcher command.Dispatcher
}

func NewDefaultDBInteraction(reader io.Reader, writer io.Writer) *DefaultDBInteraction {
	return &DefaultDBInteraction{
		reader: reader,
		writer: writer,
	}
}

func (d *DefaultDBInteraction) Interactive(ctx context.Context, addr string) error {
	kc := client.NewDefaultKaeyaClient(addr)
	d.dispatcher = command.NewDefaultDispatcher(kc)

	_, _ = fmt.Fprintln(d.writer, welcome)
	defer func() {
		_, _ = fmt.Fprintln(d.writer, bye)
	}()

	_, _ = fmt.Fprintln(d.writer, "Available Commands: ")
	_, _ = fmt.Fprintln(d.writer, d.dispatcher.Helps())

	_, _ = fmt.Fprint(d.writer, prompt)

	scanner := bufio.NewScanner(d.reader)
	for scanner.Scan() {
		line := scanner.Text()

		output, err := d.dispatcher.Dispatch(ctx, line)
		if err == nil {
			_, _ = fmt.Fprintln(d.writer, output)
		} else {
			if errors.Is(err, command.ErrExit) {
				break
			}

			_, _ = fmt.Fprintln(d.writer, err.Error())
		}

		_, _ = fmt.Fprint(d.writer, prompt)
	}

	return nil

}
