package engine

import (
	"sync"
)

type commandsQueue struct {
	commands    []Command
	locker      sync.Mutex
	hasElements chan struct{}
}

func (queue *commandsQueue) isEmpty() bool {
	return len(queue.commands) == 0
}

func (queue *commandsQueue) push(cmd Command) {
	queue.locker.Lock()
	defer queue.locker.Unlock()
	wasEmpty := queue.isEmpty()
	queue.commands = append(queue.commands, cmd)
	if wasEmpty {
		queue.hasElements <- struct{}{}
	}
}

func (queue *commandsQueue) pull() Command {
	queue.locker.Lock()
	defer queue.locker.Unlock()
	if queue.isEmpty() {
		queue.locker.Unlock()
		<-queue.hasElements
		queue.locker.Lock()
	}
	cmd := queue.commands[0]
	queue.commands[0] = nil
	queue.commands = queue.commands[1:]
	return cmd
}
