package logger

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	output  bytes.Buffer
	testStr string = "test"
)

func TestNewLogger(t *testing.T) {
	ioWriter := bufio.NewWriter(&output)
	logger := NewLogger(ioWriter)

	logger.Info(testStr)
	assert.Contains(t, output.String(), testStr)
	assert.Contains(t, output.String(), "info")
	output.Reset()

	logger.Err(testStr)
	assert.Contains(t, output.String(), testStr)
	assert.Contains(t, output.String(), "err")
	output.Reset()

	logger.Fatal(testStr)
	assert.Contains(t, output.String(), testStr)
	assert.Contains(t, output.String(), "fatal")
}
