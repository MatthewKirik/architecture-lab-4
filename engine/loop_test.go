package engine

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/MatthewKirik/architecture-lab-4/command"
	"github.com/stretchr/testify/assert"
)

// func captureStdoutString() string {
// 	old := os.Stdout
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w

// 	// fmt.Println("ABOBA!!!")

// 	w.Close()
// 	os.Stdout = old
// 	// fmt.Println("ABOBA!!!")

// 	var buf bytes.Buffer
// 	io.Copy(&buf, r)
// 	return buf.String()
// }

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

func ExampleEventLoop() {
	loop := new(EventLoop)
	loop.Start()
	loop.Post(&command.PrintCmd{Text: "Hello, I am event loop!"})
	loop.AwaitFinish()

	// Output:
	// Hello, I am event loop!
}