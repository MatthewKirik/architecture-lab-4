package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var errPrefix = "SYNTAX ERROR:"

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
	inputStr := "split error:::String :::"

	cmd := Parse(inputStr)

	assert.Contains(cmd.(*PrintCmd).Text, errPrefix)
}

func TestErrorSplitCmdCountArgs(t *testing.T) {
	assert := assert.New(t)
	inputStr := "split too many arguments"

	cmd := Parse(inputStr)

	assert.Contains(cmd.(*PrintCmd).Text, errPrefix)
}

func TestErrorPrintCmdCountArgs(t *testing.T) {
	assert := assert.New(t)
	inputStr := "print too many arguments"

	cmd := Parse(inputStr)

	assert.Contains(cmd.(*PrintCmd).Text, errPrefix)
}

func TestErrorUnknownCommand(t *testing.T) {
	assert := assert.New(t)
	inputStr := "perkele 1337"
	// errPrefix := "SYNTAX ERROR:"

	cmd := Parse(inputStr)

	assert.Contains(cmd.(*PrintCmd).Text, errPrefix)
}

func ExampleParse() {
	inputPrint := "print your-string-here"
	inputSplit := "split"

	cmdPrint := Parse(inputPrint)
	cmdSplit := Parse(inputSplit)

	var handler Handler
	handler.Post(cmdPrint)
	handler.Post(cmdSplit)
}
