package command

import (
	"fmt"
	"strings"
)

type cmdProcessor func(args []string) (Command, error)

func processPrintCmd(args []string) (Command, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("wrong number of arguments."+
			"Expected 1, got %d instead", len(args))
	}

	return &PrintCmd{args[0]}, nil
}

func processSplitCmd(args []string) (Command, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("wrong number of arguments."+
			"Expected 2, got %d instead", len(args))
	}

	if len(args[1]) != 1 {
		return nil, fmt.Errorf("separator's length should be 1 character long")
	}

	return &SplitCmd{args[0], args[1]}, nil
}

var commandsMap = map[string]cmdProcessor{
	"print": processPrintCmd,
	"split": processSplitCmd,
}

func findCommand(commandStr string) (cmdProcessor, error) {
	for cmd, fn := range commandsMap {
		if cmd == commandStr {
			return fn, nil
		}
	}

	return nil, fmt.Errorf("unknown command")
}

func Parse(text string) Command {
	values := strings.Split(text, " ")
	errPrefix := "SYNTAX ERROR: "
	cmdFunc, err := findCommand(values[0])
	if err != nil {
		return &PrintCmd{
			Text: errPrefix + err.Error(),
		}
	}

	command, err := cmdFunc(values[1:])
	if err != nil {
		return &PrintCmd{
			Text: errPrefix + err.Error(),
		}
	}

	return command
}
