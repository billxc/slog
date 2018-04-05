package slog

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	TRACE = iota
	DEBUG
	INFO
	ERROR
)

var defaultLogger = Logger{ERROR, os.Stdout}

type Logger struct {
	Level int
	Out   io.Writer
}

func SetDefault(d Logger) {
	defaultLogger = d
}

func SetDefaultLevel(level int) {
	defaultLogger.Level = level
}

func GetLogger(level int, out io.Writer) Logger {
	if out == nil {
		out = os.Stdout
	}
	return Logger{level, out}
}

func FileLogger(level int, fileName string) Logger {
	out, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		out = os.Stdout
	}
	return Logger{level, out}
}

func Info(format string, inputs ...interface{}) {
	defaultLogger.Info(format, inputs...)
}

func (l *Logger) Info(format string, inputs ...interface{}) {
	l.log(INFO, "[INFO] ", format, inputs...)
}

func Error(format string, inputs ...interface{}) {
	defaultLogger.Error(format, inputs...)
}

func (l *Logger) Error(format string, inputs ...interface{}) {
	l.log(ERROR, "[ERROR] ", format, inputs...)
}

func Debug(format string, inputs ...interface{}) {
	defaultLogger.Debug(format, inputs...)
}

func (l *Logger) Debug(format string, inputs ...interface{}) {
	l.log(DEBUG, "[DEBUG] ", format, inputs...)
}

func Trace(format string, inputs ...interface{}) {
	defaultLogger.Trace(format, inputs...)
}

func (l *Logger) Trace(format string, inputs ...interface{}) {
	l.log(TRACE, "[TRACE] ", format, inputs...)
}

func (l *Logger) log(level int, prefix, format string, inputs ...interface{}) {
	if level >= l.Level {
		fmt.Fprintf(l.Out, time.Now().Format("2006-01-02 15:04:05.000 ")+prefix+format+"\n", inputs...)
	}
}
