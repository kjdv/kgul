package logging

import (
	"bytes"
	"testing"

	"github.com/kjdv/kgul/testing/expects"
)

type testBuffer struct {
	channel chan string
}

func NewTestBuffer() testBuffer {
	return testBuffer{make(chan string, 100)}
}

func (tb *testBuffer) Write(buf []byte) (int, error) {
	tb.channel <- string(buf)

	return len(buf), nil
}

func (tb *testBuffer) Get() string {
	return <-tb.channel
}

func TestDistinctLevels(t *testing.T) {
	previous := notSet
	for _, current := range []int{ALL, DEBUG, INFO, WARNING, ERROR, NONE} {

		if previous >= current {
			t.Fail()
		}

		previous = current
	}
}

func lastLine(b *bytes.Buffer) string {
	s := b.String()
	b.Reset()
	return s
}

const tsRegex = "[0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{6}"

func regex(tag string, level string, tail string) string {
	return "^" + tsRegex + " " + tag + " " + level + ": " + tail + "\n$"
}

func TestPrintLog(t *testing.T) {
	expect := expects.New(t)

	output := NewTestBuffer()
	SetOutput(&output)

	logger := New("test")

	logger.Info("some ", 2, " messages")
	expect.Regex(regex("test", "I", "some 2 messages"), output.Get())

	logger.Infof("format %d string", 3)
	expect.Regex(regex("test", "I", "format 3 string"), output.Get())

	logger.Warning("warn")
	expect.Regex(regex("test", "W", "warn"), output.Get())

	logger.Warningf("%s", "warn")
	expect.Regex(regex("test", "W", "warn"), output.Get())

	logger.Error("error")
	expect.Regex(regex("test", "E", "error"), output.Get())

	logger.Errorf("%s", "error")
	expect.Regex(regex("test", "E", "error"), output.Get())

	logger.SetLevel(DEBUG)
	logger.Debug("debug")
	expect.Regex(regex("test", "D", "debug"), output.Get())

	logger.Debugf("%s", "debug")
	expect.Regex(regex("test", "D", "debug"), output.Get())
}

func TestSurpressLog(t *testing.T) {
	expect := expects.New(t)

	output := NewTestBuffer()
	SetOutput(&output)

	logger := New("test")

	shouldLog := func() {
		expect.Regex("^"+tsRegex+" test ", output.Get())
	}
	shouldNotLog := func() {
		output.Write([]byte("sentinel"))
		expect.Equals("sentinel", output.Get())
	}

	logger.SetLevel(ALL)
	logger.Error("e")
	shouldLog()
	logger.Warning("w")
	shouldLog()
	logger.Info("i")
	shouldLog()
	logger.Debug("d")
	shouldLog()

	logger.SetLevel(DEBUG)
	logger.Error("e")
	shouldLog()
	logger.Warning("w")
	shouldLog()
	logger.Info("i")
	shouldLog()
	logger.Debug("d")
	shouldLog()

	logger.SetLevel(INFO)
	logger.Error("e")
	shouldLog()
	logger.Warning("w")
	shouldLog()
	logger.Info("i")
	shouldLog()
	logger.Debug("d")
	shouldNotLog()

	logger.SetLevel(WARNING)
	logger.Error("e")
	shouldLog()
	logger.Warning("w")
	shouldLog()
	logger.Info("i")
	shouldNotLog()
	logger.Debug("d")
	shouldNotLog()

	logger.SetLevel(ERROR)
	logger.Error("e")
	shouldLog()
	logger.Warning("w")
	shouldNotLog()
	logger.Info("i")
	shouldNotLog()
	logger.Debug("d")
	shouldNotLog()

	logger.SetLevel(NONE)
	logger.Error("e")
	shouldNotLog()
	logger.Warning("w")
	shouldNotLog()
	logger.Info("i")
	shouldNotLog()
	logger.Debug("d")
	shouldNotLog()
}

func TestGlobalVersusLocal(t *testing.T) {
	expect := expects.New(t)

	output := NewTestBuffer()
	SetOutput(&output)

	logger := New("test")

	shouldLog := func() {
		expect.Regex("^"+tsRegex+" test", output.Get())
	}
	shouldNotLog := func() {
		output.Write([]byte("sentinel"))
		expect.Equals("sentinel", output.Get())
	}

	expect.Equals(INFO, logLevel, "self-test")

	logger.Debug("should not be logged")
	shouldNotLog()
	logger.Info("should be logged")
	shouldLog()

	logger.SetLevel(DEBUG)

	logger.Debug("should now be enabled")
	shouldLog()
	logger.Info("should still be enabled")
	shouldLog()

	// new logger should follow old rules
	logger = New("test-recreated")
	logger.Debug("should not logged")
	shouldNotLog()
	logger.Info("should be logged")
	shouldLog()

	// manipulate global level
	(*logLevelT)(&logLevel).Set("all")

	logger.Debug("should now be enabled")
	shouldLog()
	logger.Info("should still be enabled")
	shouldLog()
}
