package engine

import "sync"

type trustedHandler struct {
	loop *EventLoop
}

func (th *trustedHandler) Post(cmd Command) {
	th.loop.commands.push(cmd)
}

type EventLoop struct {
	commands *commandsQueue
	stopCond sync.Cond

	stopLocker    sync.Mutex
	stopRequested bool
	stopped       bool
}

func (loop *EventLoop) init() {
	loop.commands = &commandsQueue{
		hasElements: *sync.NewCond(&sync.Mutex{}),
	}
	loop.stopCond = *sync.NewCond(&sync.Mutex{})
}

func (loop *EventLoop) dispose() {
	loop.commands = nil
	loop.stopCond = sync.Cond{}
	loop.stopLocker = sync.Mutex{}
	loop.stopRequested = false
	loop.stopped = false
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

func (loop *EventLoop) Post(cmd Command) {
	if loop.isStopRequested() || loop.isStopped() {
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
	loop.dispose()
}
