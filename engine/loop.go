package engine

import (
	"sync"
)

type EventLoop struct {
	commands *commandsQueue
	locker   sync.Mutex
}

func (loop *EventLoop) Start() {
	loop.commands = &commandsQueue{
		hasElements: make(chan struct{}),
	}
	go loop.listen()
}

func (loop *EventLoop) listen() {
	for !loop.commands.isEmpty() {
		cmd := loop.commands.pull()
		cmd.Execute(loop)
	}
}

func (loop *EventLoop) Post(cmd Command) {
}

func (loop *EventLoop) AwaitFinish() {
}
