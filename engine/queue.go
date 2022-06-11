package engine

import (
	"sync"
)

type commandsQueue struct {
	commands    []Command
	locker      sync.Mutex
	emptyLocker sync.Mutex
}

func (queue *commandsQueue) isEmpty() bool {
	queue.locker.Lock()
	defer queue.locker.Unlock()
	return len(queue.commands) == 0
}

func (queue *commandsQueue) push(cmd Command) {
	queue.locker.Lock()
	queue.commands = append(queue.commands, cmd)
	queue.locker.Unlock()
	queue.emptyLocker.TryLock()
	queue.emptyLocker.Unlock()
}

func (queue *commandsQueue) pull() Command {
	queue.emptyLocker.Lock()
	queue.locker.Lock()
	defer queue.locker.Unlock()
	cmd := queue.commands[0]
	queue.commands[0] = nil
	queue.commands = queue.commands[1:]
	if len(queue.commands) > 0 {
		queue.emptyLocker.Unlock()
	}
	return cmd
}
