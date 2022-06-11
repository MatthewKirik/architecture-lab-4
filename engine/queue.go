package engine

import (
	"sync"
)

type commandsQueue struct {
	commands []Command
	cond     sync.Cond
}

func (queue *commandsQueue) isEmpty() bool {
	queue.cond.L.Lock()
	defer queue.cond.L.Unlock()
	return len(queue.commands) == 0
}

func (queue *commandsQueue) push(cmd Command) {
	queue.cond.L.Lock()
	queue.commands = append(queue.commands, cmd)
	queue.cond.L.Unlock()
	queue.cond.Broadcast()
}

func (queue *commandsQueue) pull() Command {
	for len(queue.commands) == 0 {
		queue.cond.Wait()
	}
	queue.cond.L.Lock()
	defer queue.cond.L.Unlock()
	cmd := queue.commands[0]
	queue.commands[0] = nil
	queue.commands = queue.commands[1:]
	return cmd
}
