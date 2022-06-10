package engine

import (
	"sync"
)

type commandsQueue struct {
	commands []Command
	locker   sync.Mutex
}

func (queue *commandsQueue) isEmpty() bool {
	return len(queue.commands) == 0
}

func (queue *commandsQueue) push(cmd Command) {
	queue.locker.Lock()
	queue.commands = append(queue.commands, cmd)
	queue.locker.Unlock()
}

func (queue *commandsQueue) pull() Command {
	queue.locker.Lock()
	defer queue.locker.Unlock()
	if queue.isEmpty() {
		return nil
	}
	cmd := queue.commands[0]
	queue.commands[0] = nil
	queue.commands = queue.commands[1:]
	return cmd
}
