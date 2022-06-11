package engine

import (
	"sync"
)

type commandsQueue struct {
	commands    []Command
	locker      sync.Mutex
	hasElements chan struct{}
}

func newCommandsQueue() *commandsQueue {
	return &commandsQueue{
		hasElements: make(chan struct{}, 1),
	}
}

func (queue *commandsQueue) isEmpty() bool {
	return len(queue.commands) == 0
}

func (queue *commandsQueue) push(cmd Command) {
	queue.locker.Lock()
	queue.commands = append(queue.commands, cmd)
	queue.locker.Unlock()

	if len(queue.hasElements) == 0 {
		queue.hasElements <- struct{}{}
	}
}

func (queue *commandsQueue) pull() Command {
	<-queue.hasElements
	queue.locker.Lock()
	defer queue.locker.Unlock()
	cmd := queue.commands[0]
	queue.commands[0] = nil
	queue.commands = queue.commands[1:]
	return cmd
}
