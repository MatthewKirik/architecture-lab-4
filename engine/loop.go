package engine

import (
	"sync"

	"github.com/MatthewKirik/architecture-lab-4/command"
)

type trustedHandler struct {
	loop *EventLoop
}

func (th *trustedHandler) Post(cmd command.Command) {
	th.loop.commands.push(cmd)
}

type EventLoop struct {
	commands *commandsQueue
	stopCond sync.Cond

	stopLocker    sync.Mutex
	stopRequested bool
	stopped       bool
	isRunning     bool
}

func (loop *EventLoop) init() {
	loop.commands = &commandsQueue{
		hasElements: *sync.NewCond(&sync.Mutex{}),
	}
	loop.stopCond = *sync.NewCond(&sync.Mutex{})
	loop.stopLocker = sync.Mutex{}
	loop.isRunning = true
	loop.stopRequested = false
	loop.stopped = false
}

func (loop *EventLoop) dispose() {
	loop.commands = nil
}

func (loop *EventLoop) stop() {
	loop.stopLocker.Lock()
	loop.stopped = true
	loop.stopRequested = false
	loop.isRunning = false
	loop.stopLocker.Unlock()
	loop.stopCond.Broadcast()
	loop.dispose()
}

func (loop *EventLoop) listen() {
	for !loop.isStopRequested() || !loop.commands.isEmpty() {
		cmd := loop.commands.pull()
		cmd.Execute(&trustedHandler{loop})
	}
	loop.stop()
}

func (loop *EventLoop) verifyRunning() {
	if !loop.isRunning {
		panic("Unable to perform an action. Loop was not started")
	}
}

func (loop *EventLoop) Start() {
	loop.init()
	go loop.listen()
}

func (loop *EventLoop) isStopRequested() bool {
	loop.stopLocker.Lock()
	defer loop.stopLocker.Unlock()
	return loop.stopRequested
}

func (loop *EventLoop) isStopped() bool {
	loop.stopLocker.Lock()
	defer loop.stopLocker.Unlock()
	return loop.stopped
}

func (loop *EventLoop) Post(cmd command.Command) {
	loop.verifyRunning()
	if loop.isStopRequested() || loop.isStopped() {
		return
	}
	loop.commands.push(cmd)
}

func (loop *EventLoop) AwaitFinish() {
	loop.verifyRunning()
	if loop.isStopped() {
		return
	}
	if !loop.isStopRequested() {
		loop.stopLocker.Lock()
		loop.stopRequested = true
		loop.stopLocker.Unlock()
	}
	loop.stopCond.L.Lock()
	loop.stopCond.Wait()
	loop.stopCond.L.Unlock()
}
