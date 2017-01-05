package simplelog

import (
	"fmt"
	"io"
	"os"
	"time"
)

// These are logger levels prefixes.
const (
	debug = "[DEBUG]"
	info  = "[INFO]"
	err   = "[ERROR]"
)

var (
	red    = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	yellow = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	white  = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	green  = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

// SimpleLogger is a struct which holds the chosen output for the log.
type SimpleLogger struct {
	output io.Writer
}

// InitLogger returns logger with standarized log format to use throughout the app.
func InitLogger() *SimpleLogger {
	return &SimpleLogger{
		output: os.Stderr,
	}
}

func (l *SimpleLogger) print(msg, prefix, color string) {
	t := time.Now().Format("2006-02-01 15:04:05")
	fmt.Fprintf(l.output, "%s%7s%s [ %s ] %s\n", color, prefix, reset, t, msg)
}

// Debug prints log with "[DEBUG]" prefix
func (l *SimpleLogger) Debug(msg string) {
	l.print(msg, debug, yellow)
}

// Info prints log with "[INFO]" prefix
func (l *SimpleLogger) Info(msg string) {
	l.print(msg, info, green)
}

// Error prints log with "[ERROR]" prefix
func (l *SimpleLogger) Error(msg string) {
	l.print(msg, err, red)
}
