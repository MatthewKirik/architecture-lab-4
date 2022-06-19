package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MatthewKirik/architecture-lab-4/engine"
)

type cmdProcessor func(args []string) (engine.Command, error)

func processPrintCmd(args []string) (engine.Command, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Wrong number of arguments."+
			"Expected 1, got %d instead", len(args))
	}

	return &printCommand{args[0]}, nil
}

func processSplitCmd(args []string) (engine.Command, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Wrong number of arguments."+
			"Expected 2, got %d instead", len(args))
	}

	if len(args[1]) != 1 {
		return nil, fmt.Errorf("Separator's length should be 1 character long")
	}

	return &splitCmd{args[0], args[1]}, nil
}

var commandsArr = map[string]cmdProcessor{
	"print": processPrintCmd,
	"split": processSplitCmd,
}

func findCommand(commandStr string) (cmdProcessor, error) {
	for cmd, fn := range commandsArr {
		if cmd == commandStr {
			return fn, nil
		}
	}

	return nil, errors.New("unknown command")
}

func parse(text string) engine.Command {
	values := strings.Split(text, " ")
	cmdFunc, err := findCommand(values[0])
	if err != nil {
		return &printCommand{
			text: err.Error(),
		}
	}

	command, err := cmdFunc(values[1:])
	if err != nil {
		return &printCommand{
			text: err.Error(),
		}
	}

	return command
}
