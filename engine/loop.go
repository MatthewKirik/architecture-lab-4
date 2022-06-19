package engine

import "sync"

type EventLoop struct {
	commands      *commandsQueue
	stopRequested bool
	stopped       bool
	stopLocker    sync.Mutex
	stopCond      sync.Cond
}

func (loop *EventLoop) Start() {
	loop.commands = &commandsQueue{
		cond: *sync.NewCond(&sync.Mutex{}),
	}
	loop.stopCond = *sync.NewCond(&sync.Mutex{})
	go loop.listen()
}

func (loop *EventLoop) listen() {
	for !loop.isStopRequested() || !loop.commands.isEmpty() {
		cmd := loop.commands.pull()
		cmd.Execute(&trustedHandler{loop})
	}
	loop.stopLocker.Lock()
	loop.stopped = true
	loop.stopLocker.Unlock()
	loop.stopCond.Broadcast()
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

func (loop *EventLoop) Post(cmd Command) {
	if loop.isStopRequested() {
		return
	}
	loop.commands.push(cmd)
}

func (loop *EventLoop) AwaitFinish() {
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
}

type trustedHandler struct {
	loop *EventLoop
}

func (th *trustedHandler) Post(cmd Command) {
	th.loop.commands.push(cmd)
}
