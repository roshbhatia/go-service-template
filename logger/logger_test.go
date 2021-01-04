package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	infoOutput, errOutput, fatalOutput bytes.Buffer
	testStr                            string = "test"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()

	logger.Info.SetOutput(&infoOutput)
	logger.Err.SetOutput(&errOutput)
	logger.Fatal.SetOutput(&fatalOutput)

	logger.Info.Println(testStr)
	logger.Err.Println(testStr)
	logger.Fatal.Println(testStr)

	assert.Contains(t, infoOutput.String(), testStr)
	assert.Contains(t, infoOutput.String(), "info")
	assert.Contains(t, errOutput.String(), testStr)
	assert.Contains(t, errOutput.String(), "err")
	assert.Contains(t, fatalOutput.String(), testStr)
	assert.Contains(t, fatalOutput.String(), "fatal")
}
