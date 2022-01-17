package common

import (
	"log"
	"os"
)

const (
	DebugLevel = iota
	InfoLevel
	ErrorLevel
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
}

type sbLogger struct {
	logLevel int
	logger   *log.Logger
}

func (l *sbLogger) Debug(args ...interface{}) {
	if l.logLevel <= DebugLevel {
		l.logger.Print(args...)
	}
}

func (l *sbLogger) Debugf(format string, args ...interface{}) {
	if l.logLevel <= DebugLevel {
		l.logger.Printf(format, args...)
	}
}

func (l *sbLogger) Debugln(args ...interface{}) {
	if l.logLevel <= DebugLevel {
		l.logger.Println(args...)
	}
}

func (l *sbLogger) Info(args ...interface{}) {
	if l.logLevel <= InfoLevel {
		l.logger.Print(args...)
	}
}

func (l *sbLogger) Infof(format string, args ...interface{}) {
	if l.logLevel <= InfoLevel {
		l.logger.Printf(format, args...)
	}
}

func (l *sbLogger) Infoln(args ...interface{}) {
	if l.logLevel <= InfoLevel {
		l.logger.Println(args...)
	}
}

func (l *sbLogger) Error(args ...interface{}) {
	if l.logLevel <= ErrorLevel {
		l.logger.Print(args...)
	}
}

func (l *sbLogger) Errorf(format string, args ...interface{}) {
	if l.logLevel <= ErrorLevel {
		l.logger.Printf(format, args...)
	}
}

func (l *sbLogger) Errorln(args ...interface{}) {
	if l.logLevel <= ErrorLevel {
		l.logger.Println(args...)
	}
}

func InitializeLogger(_ *SbContext, logLevel int) Logger {
	return &sbLogger{logLevel, log.New(os.Stdout, "serverbox: ", 0)}
}
