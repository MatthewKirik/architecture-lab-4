package engine

import (
	"sync"
)

type EventLoop struct {
	commands      *commandsQueue
	locker        sync.Mutex
	stopRequested bool
}

func (loop *EventLoop) Start() {
	loop.commands = &commandsQueue{
		hasElements: make(chan struct{}),
	}
	go loop.listen()
}

func (loop *EventLoop) listen() {
	for !loop.stopRequested || !loop.commands.isEmpty() {
		cmd := loop.commands.pull()
		cmd.Execute(loop)
	}
}

func (loop *EventLoop) Post(cmd Command) {
	if loop.stopRequested {
		return
	}
	loop.commands.push(cmd)
}

func (loop *EventLoop) AwaitFinish() {
	loop.stopRequested = true
}
