package engine

import (
	"sync"
)

type commandsQueue struct {
	commands    []Command
	hasElements sync.Cond
}

func (queue *commandsQueue) isEmpty() bool {
	queue.hasElements.L.Lock()
	defer queue.hasElements.L.Unlock()
	return len(queue.commands) == 0
}

func (queue *commandsQueue) push(cmd Command) {
	queue.hasElements.L.Lock()
	queue.commands = append(queue.commands, cmd)
	queue.hasElements.L.Unlock()
	queue.hasElements.Broadcast()
}

func (queue *commandsQueue) pull() Command {
	queue.hasElements.L.Lock()
	for len(queue.commands) == 0 {
		queue.hasElements.Wait()
	}
	defer queue.hasElements.L.Unlock()
	cmd := queue.commands[0]
	queue.commands[0] = nil
	queue.commands = queue.commands[1:]
	return cmd
}
