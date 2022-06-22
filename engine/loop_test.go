package engine

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/MatthewKirik/architecture-lab-4/command"
	"github.com/stretchr/testify/assert"
)

func captureStdoutString() func() (string, error) {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	done := make(chan error, 1)

	save := os.Stdout
	os.Stdout = w

	var buf strings.Builder

	go func() {
		_, err := io.Copy(&buf, r)
		r.Close()
		done <- err
	}()

	return func() (string, error) {
		os.Stdout = save
		w.Close()
		err := <-done
		return buf.String(), err
	}
}

func TestExecutionPrintCommand(t *testing.T) {
	assert := assert.New(t)
	loop := new(EventLoop)
	input := "this is testing message"
	expected := "this is testing message\n"
	cmd := command.PrintCmd{Text: input}

	loop.Start()
	getStr := captureStdoutString()
	loop.Post(&cmd)
	loop.AwaitFinish()

	capturedOutput, err := getStr()
	if err != nil {
		panic(err)
	}

	assert.Equal(capturedOutput, expected)
}

func TestExecutionSplitCommand(t *testing.T) {
	assert := assert.New(t)
	loop := new(EventLoop)
	inputStr := "split,me,please"
	inputSep := ","
	expected := "split\nme\nplease\n"
	cmd := command.SplitCmd{
		Text:      inputStr,
		Separator: inputSep,
	}

	loop.Start()
	getStr := captureStdoutString()
	loop.Post(&cmd)
	loop.AwaitFinish()

	capturedOutput, err := getStr()
	if err != nil {
		panic(err)
	}

	assert.Equal(capturedOutput, expected)
}

func TestStopEmptyLoop(t *testing.T) {
	loop := new(EventLoop)
	loop.Start()
	loop.AwaitFinish()
}

func TestMultipleAwaits(t *testing.T) {
	loop := new(EventLoop)
	loop.Start()
	loop.Post(&command.PrintCmd{Text: "something here! ?@ s ddd!"})
	loop.Post(&command.PrintCmd{Text: "Who are you?!"})
	loop.Post(&command.SplitCmd{
		Text:      "How do u like it?",
		Separator: " ",
	})
	go loop.AwaitFinish()
	go loop.AwaitFinish()
	go loop.AwaitFinish()
	loop.AwaitFinish()
}

func TestPostCommandAfterLoopWasStopped(t *testing.T) {
	loop := new(EventLoop)
	loop.Start()
	loop.Post(&command.PrintCmd{Text: "Yeah I am going forward into event loop :)"})
	loop.AwaitFinish()
	loop.Post(&command.PrintCmd{Text: "Oh no, I will be executed in stopped loop!"})
	loop.Post(&command.PrintCmd{Text: "Me too, bro :("})
	loop.Post(&command.SplitCmd{
		Text:      "Yall stupid go home!",
		Separator: " ",
	})
}

func ExampleEventLoop() {
	loop := new(EventLoop)
	loop.Start()
	loop.Post(&command.PrintCmd{Text: "Hello, I am event loop!"})
	loop.AwaitFinish()

	// Output:
	// Hello, I am event loop!
}
