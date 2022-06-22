package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsPath = "../tests/parset-tests.txt"

func TestParsePrintCmd(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("", "")
}

func TestParseSplitCmd(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("", "")
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
