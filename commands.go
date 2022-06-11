package main

import (
	"fmt"
	"strings"

	"github.com/MatthewKirik/architecture-lab-4/engine"
)

type printCommand struct {
	text string
}

func (cmd *printCommand) Execute(handler engine.Handler) {
	fmt.Println(cmd.text)
}

type splitCmd struct {
	text      string
	separator string
}

func (cmd *splitCmd) Execute(handler engine.Handler) {
	parts := strings.Split(cmd.text, cmd.separator)
	for _, part := range parts {
		handler.Post(&printCommand{text: part})
	}
}
