package engine

type EventLoop struct {
	commands      *commandsQueue
	stopRequested bool
	stopped       chan struct{}
}

func (loop *EventLoop) Start() {
	loop.commands = new(commandsQueue)
	go loop.listen()
}

func (loop *EventLoop) listen() {
	for !loop.stopRequested || !loop.commands.isEmpty() {
		cmd := loop.commands.pull()
		cmd.Execute(loop)
	}
	loop.stopped <- struct{}{}
}

func (loop *EventLoop) Post(cmd Command) {
	if loop.stopRequested {
		return
	}
	loop.commands.push(cmd)
}

func (loop *EventLoop) AwaitFinish() {
	loop.Post(FuncCommand(func(handler Handler) {
		loop.stopRequested = true
	}))
	<-loop.stopped
}
