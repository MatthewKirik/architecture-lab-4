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
	// TODO: implements

	return nil, nil
}

var commandsArr = map[string]cmdProcessor{
	"print": processPrintCmd,
	"split": processSplitCmd,
}

// var commandsArr = []string{
// 	"split",
// 	"print",
// }

/*
{
	args: 2 int
	command: "print" string
	type: &printCommand ???
}

*/

// JS:
// const commandsArr = {
//	split: "split",
//	print: "print",
// }

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

// func handleSplit(split *[]string) error {
// 	if len(*split) != 3 {
// 		return fmt.Errorf("Wrong count of arguments, expected 3  - got: %d", len(*split))
// 	}
// 	if len((*split)[2]) != 1 {
// 		return errors.New("Separator should be 1 character long!")
// 	}
// 	return nil
// }

// func handlePrint(split *[]string) error {
// 	if len(*split) != 2 {
// 		return fmt.Errorf("Wrong count of arguments, expected 2  - got: %d", len(*split))
// 	}
// 	return nil
// }
