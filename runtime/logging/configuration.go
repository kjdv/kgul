package logging

import (
	"errors"
	"flag"
	"os"
)

type logLevelT int

const validLogLevels = "'all', 'debug', 'info', 'warning', 'error' or 'none'"

func (l *logLevelT) String() string {
	switch *l {
	case ALL:
		return "all"
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case ERROR:
		return "error"
	case NONE:
		return "none"
	default:
		return ""
	}
}

func (l *logLevelT) Set(value string) error {
	switch value {
	case "all":
		*l = ALL
	case "debug":
		*l = DEBUG
	case "info":
		*l = INFO
	case "warning":
		*l = WARNING
	case "error":
		*l = ERROR
	case "none":
		*l = NONE
	default:
		return errors.New("should be " + validLogLevels)
	}

	return nil
}

type logFile string

func (f *logFile) String() string {
	return string(*f)
}

func (f *logFile) Set(value string) error {
	w, err := os.Create(value)
	if err == nil {
		SetOutput(w)
	}
	return err
}

func init() {
	var dummyFile logFile = "stdout"

	flag.Var((*logLevelT)(&logLevel), "logLevel", "global log level")
	flag.Var(&dummyFile, "logFile", "file to write logs to")

	// when we are running go test, surpress logging by default
	if flag.Lookup("test.v") != nil {
		logWriter = nil
	}
}
