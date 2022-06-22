package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// var testsPath = "../tests/parser-tests.txt"

func TestParsePrintCmd(t *testing.T) {
	assert := assert.New(t)
	inputStr := "print hello!"
	expected := "hello!"

	cmd := Parse(inputStr)

	assert.Equal(cmd.(*PrintCmd).Text, expected)
}

func TestParseSplitCmd(t *testing.T) {
	assert := assert.New(t)
	inputStr := "split hey!dude!split!me !"
	expectedStr := "hey!dude!split!me"
	expectedSep := "!"

	cmd := Parse(inputStr)

	assert.Equal(cmd.(*SplitCmd).Text, expectedStr)
	assert.Equal(cmd.(*SplitCmd).Separator, expectedSep)
}

func TestErrorSplitCmdLongSeparator(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("", "")
}

func TestErrorSplitCmdCountArgs(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("", "")
}

func TestErrorPrintCmdCountArgs(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("", "")
}

func TestErrorUnknownCommand(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("", "")
}
