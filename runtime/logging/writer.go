package logging

import (
	"fmt"
	"io"
	"kdv/runtime/assert"
	"os"
	"time"
)

var logWriter io.Writer = os.Stdout

func SetOutput(w io.Writer) {
	logWriter = w
}

type logJob interface {
	do()
}

type item struct {
	ts    time.Time
	tag   string
	level int
	line  string
}

func (i item) do() {
	if logWriter == nil {
		return
	}

	year, month, day := i.ts.Date()
	hour, min, sec := i.ts.Clock()
	us := i.ts.Nanosecond() / 1000

	var ltag rune
	switch i.level {
	case DEBUG:
		ltag = 'D'
	case INFO:
		ltag = 'I'
	case WARNING:
		ltag = 'W'
	case ERROR:
		ltag = 'E'
	default:
		assert.Fail()
	}

	full := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d.%06d %s %c: %s\n", year, month, day, hour, min, sec, us, i.tag, ltag, i.line)

	logWriter.Write([]byte(full))
}

const logQueueSize = 1024

var channel chan logJob = make(chan logJob, logQueueSize)

func init() {
	go func() {
		for {
			i := <-channel
			i.do()
		}
	}()
}

func async_output(i item) {
	select {
	case channel <- i:
	default:
	}
}

type sentinel struct {
	done chan bool
}

func (s sentinel) do() {
	s.done <- true
}

func (s sentinel) wait() {
	<-s.done
}

func Drain() {
	// wait until the log queue is empty

	s := sentinel{make(chan bool)}
	channel <- s

	s.wait()
}
