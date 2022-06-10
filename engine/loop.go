package engine

import "fmt"

type EventLoop struct {
	commands *commandsQueue
}

func (loop *EventLoop) Start() {
	fmt.Print("I am loop start")
	loop.commands = new(commandsQueue)
	go loop.listen()
}

func (loop *EventLoop) listen() {
	for !loop.commands.isEmpty() {
		cmd := loop.commands.pull()
		cmd.Execute(loop)
	}
}

func (loop *EventLoop) Post(cmd Command) {
	fmt.Print("I am loop post")
}

func (loop *EventLoop) AwaitFinish() {
	fmt.Print("I am loop finish")
}
