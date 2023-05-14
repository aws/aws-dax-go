package client

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota - 1
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
	LogLevelPanic
	LogLevelNoop
)

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Panic(format string, args ...interface{})
	Log(args ...interface{}) // deprecated: compatible for aws.Logger
}

// NewDefaultLogger returns a Logger which will write log messages to stdout, and
// use same formatting runes as the stdlib log.Logger
func NewDefaultLogger(logLevel ...LogLevel) Logger {
	lv := LogLevelInfo
	if len(logLevel) > 0 {
		lv = logLevel[0]
	}

	return &defaultLogger{
		level:  lv,
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// A defaultLogger provides a minimalistic logger satisfying the Logger interface.
type defaultLogger struct {
	level LogLevel
	*log.Logger
}

var _ Logger = (*defaultLogger)(nil)

func (l defaultLogger) Debug(format string, args ...interface{}) {
	if l.level <= LogLevelDebug {
		l.Printf(format, args...)
	}
}

func (l defaultLogger) Info(format string, args ...interface{}) {
	if l.level <= LogLevelInfo {
		l.Printf(format, args...)
	}
}

func (l defaultLogger) Warn(format string, args ...interface{}) {
	if l.level <= LogLevelWarn {
		l.Printf(format, args...)
	}
}

func (l defaultLogger) Error(format string, args ...interface{}) {
	if l.level >= LogLevelError {
		l.Printf(format, args...)
	}
}
func (l defaultLogger) Fatal(format string, args ...interface{}) {
	if l.level <= LogLevelFatal {
		l.Fatalf(format, args...)
	}
}
func (l defaultLogger) Panic(format string, args ...interface{}) {
	if l.level <= LogLevelPanic {
		l.Panicf(format, args...)
	}
}

func (l defaultLogger) Log(args ...interface{}) {
	l.Println(args...)
}
