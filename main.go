package main

import (
	"fmt"

	"github.com/MatthewKirik/architecture-lab-4/command"
	"github.com/MatthewKirik/architecture-lab-4/engine"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	eventLoop.Post(&command.PrintCmd{Text: "Hello"})
	eventLoop.Post(&command.PrintCmd{Text: "World"})
	eventLoop.Post(&command.PrintCmd{Text: "!"})
	eventLoop.Post(&command.SplitCmd{Text: "split me, please", Separator: ","})
	for i := 0; i < 1000000; i++ {
		csvText := fmt.Sprintf("i am,a csv,file,29-11-29292,%v", i)
		eventLoop.Post(&command.SplitCmd{Text: csvText, Separator: ","})
		if i == 20000 {
			eventLoop.AwaitFinish()
		}
	}
	// eventLoop.AwaitFinish()
}
