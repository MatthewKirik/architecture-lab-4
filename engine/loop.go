package engine

import "fmt"

type EventLoop struct {
}

func (loop *EventLoop) Start() {
	fmt.Print("I am loop start")
}

func (loop *EventLoop) Push() {
	fmt.Print("I am loop push")
}

func (loop *EventLoop) AwaitFinish() {
	fmt.Print("I am loop finish")
}
