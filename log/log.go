package log

import (
	"log"
	"fmt"
)

const (
	LevelError =iota
	LevelWarning
	LevelInfo
	Leveldebug
)

type Logger struct {
	level int
	l *log.Logger
}

func (ll * Logger) Error(format string,v...interface{}){
	if ll.level > LevelError {
		return
	}
	msg := fmt.Sprintf("[E]"+format,v...)
	ll.l.Print(msg)
}