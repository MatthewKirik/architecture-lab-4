package main

import (
	"fmt"

	"github.com/MatthewKirik/architecture-lab-4/engine"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	eventLoop.Post(&printCommand{"Hello"})
	eventLoop.Post(&printCommand{"World"})
	eventLoop.Post(&printCommand{"!"})
	eventLoop.Post(&splitCmd{"split me, please", ","})
	for i := 0; i < 1000000; i++ {
		csvText := fmt.Sprintf("i am,a csv,file,29-11-29292,%v", i)
		eventLoop.Post(&splitCmd{csvText, ","})
	}
	eventLoop.AwaitFinish()
}
