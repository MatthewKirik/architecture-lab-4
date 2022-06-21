package engine

import (
	"sync"

	"github.com/MatthewKirik/architecture-lab-4/command"
)

type commandsQueue struct {
	commands    []command.Command
	hasElements sync.Cond
}

func (queue *commandsQueue) isEmpty() bool {
	queue.hasElements.L.Lock()
	defer queue.hasElements.L.Unlock()
	return len(queue.commands) == 0
}

func (queue *commandsQueue) push(cmd command.Command) {
	queue.hasElements.L.Lock()
	queue.commands = append(queue.commands, cmd)
	queue.hasElements.L.Unlock()
	queue.hasElements.Broadcast()
}

func (queue *commandsQueue) pull() command.Command {
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
