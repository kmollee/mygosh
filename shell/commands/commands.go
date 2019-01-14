package commands

import (
	"errors"
)

type CommandFn func(...string) error

var commandMap = map[string]CommandFn{
	"cd":   Cd,
	"exit": Exit,
}

func Run(argv ...string) error {
	command := argv[0]

	if cmd, exist := commandMap[command]; !exist {
		return errors.New(argv[0] + " is not a builtin command")
	} else {
		return cmd(argv...)
	}
}
