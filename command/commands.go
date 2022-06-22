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
	resultStr := strings.Join(parts, "\n")
	handler.Post(&PrintCmd{Text: resultStr})
}
