package logging

import (
	"fmt"
	"time"
)

const (
	notSet = -1
	ALL    = iota
	DEBUG
	INFO
	WARNING
	ERROR
	NONE
)

var logLevel int = INFO

type Logger struct {
	tag   string
	level int
}

func New(tag string) Logger {
	return Logger{tag, notSet}
}

func (l *Logger) SetLevel(level int) {
	l.level = level
}

func (l *Logger) doLog(level int, line string) {
	if l.level == notSet && level < logLevel { // no per-logger override, below global level
		return
	}

	if level < l.level { // below per-logger override
		return
	}

	async_output(item{time.Now(), l.tag, level, line})
}

func (l *Logger) Debug(items ...interface{}) {
	l.doLog(DEBUG, fmt.Sprint(items...))
}

func (l *Logger) Debugf(format string, items ...interface{}) {
	l.doLog(DEBUG, fmt.Sprintf(format, items...))
}

func (l *Logger) Info(items ...interface{}) {
	l.doLog(INFO, fmt.Sprint(items...))
}

func (l *Logger) Infof(format string, items ...interface{}) {
	l.doLog(INFO, fmt.Sprintf(format, items...))
}

func (l *Logger) Warning(items ...interface{}) {
	l.doLog(WARNING, fmt.Sprint(items...))
}

func (l *Logger) Warningf(format string, items ...interface{}) {
	l.doLog(WARNING, fmt.Sprintf(format, items...))
}

func (l *Logger) Error(items ...interface{}) {
	l.doLog(ERROR, fmt.Sprint(items...))
}

func (l *Logger) Errorf(format string, items ...interface{}) {
	l.doLog(ERROR, fmt.Sprintf(format, items...))
}
