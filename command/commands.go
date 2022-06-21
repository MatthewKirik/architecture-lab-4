package command

import (
	"fmt"
	"strings"
)

type PrintCmd struct {
	Text string
}

func (pc *PrintCmd) Execute(handler Handler) {
	fmt.Println(pc.Text)
}

type SplitCmd struct {
	Text      string
	Separator string
}

func (sc *SplitCmd) Execute(handler Handler) {
	parts := strings.Split(sc.Text, sc.Separator)
	for _, part := range parts {
		handler.Post(&PrintCmd{Text: part})
	}
}
