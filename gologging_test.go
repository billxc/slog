package slog

import (
	"fmt"
	"testing"
)

func TestLevel(t *testing.T) {
	fmt.Println("Trace")
	SetDefaultLevel(TRACE)
	logall()
}

func TestFile(t *testing.T) {
	fmt.Println("Trace")
	SetDefault(FileLogger(TRACE, "out.log"))
	logall()
}

func TestInfo(t *testing.T) {
	fmt.Println("INFO:")
	SetDefault(GetLogger(INFO, nil))
	logall()
	fmt.Println("DEUBG:")
	SetDefault(GetLogger(DEBUG, nil))
	logall()
	fmt.Println("TRACE:")
	SetDefault(GetLogger(TRACE, nil))
	logall()
	fmt.Println("Error:")
	SetDefault(GetLogger(ERROR, nil))
	logall()
}

func logall() {
	Info("Test%s", "info")
	Error("Test%s", "error")
	Trace("Test%s", "Trace")
	Debug("Test%s", "debug")
}
