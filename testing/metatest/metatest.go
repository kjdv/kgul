package metatest

import (
	"fmt"
	"strings"
)

type Tester interface {
	// Interface equiavalent of testing.T, incomplete as it takes a yagni approach
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

type Metatest struct {
	errors []string
	logs   []string
}

func New() *Metatest {
	return &Metatest{}
}

func (mt *Metatest) String() string {
	if !mt.HasErrors() {
		return "no errors"
	}
	return fmt.Sprint(strings.Join(mt.errors, ",\n"))
}

func (mt *Metatest) Error(args ...interface{}) {
	msg := fmt.Sprint(args...)
	mt.errors = append(mt.errors, msg)
}

func (mt *Metatest) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	mt.errors = append(mt.errors, msg)
}

func (mt *Metatest) Log(args ...interface{}) {
	msg := fmt.Sprint(args...)
	mt.logs = append(mt.logs, msg)
}

func (mt *Metatest) Logf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	mt.logs = append(mt.logs, msg)
}

func (mt *Metatest) HasErrors() bool {
	return len(mt.errors) > 0
}

func (mt *Metatest) NumErrors() int {
	return len(mt.errors)
}

func (mt *Metatest) GetError(idx int) string {
	return mt.errors[idx]
}

func (mt *Metatest) LastError() (string, error) {
	if len(mt.errors) == 0 {
		return "", fmt.Errorf("No errors")
	}
	return mt.errors[len(mt.errors)-1], nil
}

func (mt *Metatest) Clear() {
	mt.errors = []string{}
	mt.logs = []string{}
}
