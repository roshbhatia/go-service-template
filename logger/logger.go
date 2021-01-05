package logger

import (
	"io"
	"log"
	"sync"
)

type Logger struct {
	info  *log.Logger
	err   *log.Logger
	fatal *log.Logger
	mux   sync.Mutex
	out   io.Writer
}

func NewLogger(out io.Writer) *Logger {
	logger := new(Logger)

	logger.info = log.New(out, "info: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.err = log.New(out, "err: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.fatal = log.New(out, "fatal: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

// As we have three seperate loggers, need a shared mux between the three as we're writing to the same output
// TODO: Add Infof,Errf,FatalF functions so we don't have to do many fmt.Sprintf's in our log lines
// TODO: Could just use one logger, and sprintf within each of these. That way, no mux :)
func (l *Logger) Info(logStr string) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.info.Println(logStr)
}

func (l *Logger) Err(logStr string) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.err.Println(logStr)
}

func (l *Logger) Fatal(logStr string) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.fatal.Fatal(logStr)
}
