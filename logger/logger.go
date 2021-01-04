package logger

import (
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Err   *log.Logger
	Fatal *log.Logger
}

func NewLogger() *Logger {
	logger := new(Logger)

	logger.Info = log.New(os.Stdout, "info: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Err = log.New(os.Stdout, "err: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Fatal = log.New(os.Stdout, "fatal: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
