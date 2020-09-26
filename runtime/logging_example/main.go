package main

import (
	"flag"

	"github.com/kjdv/kgul/runtime/logging"
)

func init() {
	flag.Parse()
}

func main() {
	a := logging.New("logger.A")

	a.Debug("This should be ignored")
	a.Infof("Printing info %.3f", 3.1515)
	a.SetLevel(logging.DEBUG)
	a.Debug("Debug logging is now enabled")

	b := logging.New("logger.B")
	b.SetLevel(logging.ALL)
	b.Debug("debug", "log")
	b.Infof("blah %s", "foo")
	b.Info("info", " log")
	b.Warning("warning ", "warning")
	b.Warningf("something %s", "bad")
	b.Error("error ", "panic")
	b.Errorf("error %s", "panic")

	b.SetLevel(logging.NONE)
	b.Error("this is ignored")

	c := logging.New("logger.C")
	c.Debug("used")
	c.Info("level")
	c.Warning("is")
	c.Error("global")

	logging.Drain()
}
