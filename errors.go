package errors

import (
	"fmt"
	"time"
)

func New(message string) error {
	return &errorUp{
		msg:      message,
		cause:    nil,
		pcs:      callers(),
		occurred: timeNow(),
	}
}

func Errorf(format string, args ...interface{}) error {
	return &errorUp{
		msg:      fmt.Sprintf(format, args...),
		cause:    nil,
		pcs:      callers(),
		occurred: timeNow(),
	}
}

func NewWithDepth(depth int, skip int, message string) error {
	return &errorUp{
		msg:      message,
		cause:    nil,
		pcs:      callersByAssigned(depth, skip),
		occurred: timeNow(),
	}
}

func ErrorWithDepthf(depth int, skip int, format string, args ...interface{}) error {
	return &errorUp{
		msg:      fmt.Sprintf(format, args...),
		cause:    nil,
		pcs:      callersByAssigned(depth, skip),
		occurred: timeNow(),
	}
}

func With(cause error, message string) error {
	if cause == nil {
		return nil
	}
	return &errorUp{
		msg:      message,
		cause:    cause,
		pcs:      callers(),
		occurred: timeNow(),
	}
}

func Withf(cause error, format string, args ...interface{}) error {
	if cause == nil {
		return nil
	}
	return &errorUp{
		msg:      fmt.Sprintf(format, args...),
		cause:    cause,
		pcs:      callers(),
		occurred: timeNow(),
	}
}

func WithDepth(depth int, skip int, cause error, message string) error {
	if cause == nil {
		return nil
	}
	return &errorUp{
		msg:      message,
		cause:    cause,
		pcs:      callersByAssigned(depth, skip),
		occurred: timeNow(),
	}
}

func WithDepthf(depth int, skip int, cause error, format string, args ...interface{}) error {
	if cause == nil {
		return nil
	}
	return &errorUp{
		msg:      fmt.Sprintf(format, args...),
		cause:    cause,
		pcs:      callersByAssigned(depth, skip),
		occurred: timeNow(),
	}
}

func Wrap(e error) error {
	return &errorUp{
		msg:      e.Error(),
		cause:    nil,
		pcs:      callers(),
		occurred: timeNow(),
	}
}

type errorUp struct {
	msg      string
	cause    error
	pcs      []uintptr
	occurred time.Time
}

func (e *errorUp) Error() string {
	return e.msg
}

func (e *errorUp) OccurTime() time.Time {
	return e.occurred
}

func (e *errorUp) PCS() []uintptr {
	return e.pcs
}

func (e *errorUp) Format(s fmt.State, verb rune) {
	Configure().formatFn(s, verb, e)
}

func (e *errorUp) Cause() error {
	if e.cause != nil {
		return e.cause
	}
	return nil
}

func (e *errorUp) Contains(cause error) bool {
	if e.cause == nil {
		return false
	}
	err := e.cause
	if err != nil {
		if err == cause {
			return true
		}
		hasCause, ok := err.(Errors)
		if !ok {
			return false
		}
		return hasCause.Contains(cause)
	}
	return false
}

func Contains(a error, b error) bool {
	if a == nil || b == nil {
		return false
	}
	e, ok := a.(Errors)
	if !ok {
		return false
	}
	return e.Contains(b)
}

type Errors interface {
	Error() string
	Cause() error
	OccurTime() time.Time
	PCS() []uintptr
	Contains(error) bool
	Format(fmt.State, rune)
}
