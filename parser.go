package main

import (
	"errors"
	"strings"

	"github.com/MatthewKirik/architecture-lab-4/engine"
)

// var commandsArr = map[string]int{
// 	"split": 3,
// 	"print": 2,
// }

var commandsArr = []string{
	"split",
	"print",
}

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

func findCommand(command string) (string, error) {
	for _, cmd := range commandsArr {
		if cmd == command {
			return cmd, nil
		}
	}

	return "", errors.New("unknown command")
}

func buildCommand(cmdStr string, args []string) (engine.Command, error) {
	// TODO: implements
	switch cmdStr {
	case "split":
		// ddd
		break
	case "print":
	}
	return nil, nil
}

func parse(text string) engine.Command {
	split := strings.Split(text, " ")
	cmdStr, err := findCommand(split[0])
	if err != nil {
		return &printCommand{
			text: err.Error(),
		}
	}

	command, err := buildCommand(cmdStr, split[0:])
	if err != nil {
		return &printCommand{
			text: err.Error(),
		}
	}

	return command

	// if split[0] == "split" {
	// 	handleSplit(&split)
	// 	return &splitCmd{split[1], split[2]}
	// }
	// if split[0] == "print" {
	// 	handlePrint(&split)
	// 	return &printCommand{split[1]}
	// }
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
