package commands

import (
	"errors"
	"fmt"
	"github.com/gedex/ginsta/clients"
)

type Runner struct {
	Args   []string
	Client *clients.Client
}

func (r *Runner) Execute() error {
	if len(r.Args) == 0 {
		usage()
	}
	c := r.Args[0]

	for _, cmd := range All() {
		if cmd.Name() == c && cmd.Runnable() {
			cmdArgs := r.Args[1:]

			cmd.Flag.Usage = func() {
				cmd.PrintUsage()
			}
			if err := cmd.Flag.Parse(cmdArgs); err != nil {
				return err
			}
			cmdArgs = cmd.Flag.Args()
			cmd.Callback(r, cmd, cmdArgs)
		}
	}

	return errors.New(fmt.Sprintf("Unknown command '%s'. Please run '%s help'", c, clients.Name))
}
