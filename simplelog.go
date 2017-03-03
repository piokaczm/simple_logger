package simplelog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

// Logger levels prefixes and standard time format.
const (
	debug  = "[DEBUG]"
	info   = "[INFO]"
	err    = "[ERROR]"
	format = "2006-02-01 15:04:05"
)

var (
	red    = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	yellow = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	green  = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

// SimpleLogger is a struct which holds the chosen output for the log
// and timestamp format.
type SimpleLogger struct {
	Output io.Writer
	Format string
}

// InitLogger returns logger with standarized log format to use throughout the app.
func InitLogger() *SimpleLogger {
	return &SimpleLogger{
		Output: os.Stderr,
		Format: format,
	}
}

func (l *SimpleLogger) print(msg, prefix, color string) {
	t := time.Now().Format(l.Format)
	fmt.Fprintf(l.Output, "%s%7s%s [ %s ] %s\n", color, prefix, reset, t, msg)
}

// Debug prints log with "[DEBUG]" prefix
func (l *SimpleLogger) Debug(msg string) {
	l.print(fmt.Sprintf("%s -> %s", funcCaller(), msg), debug, yellow)
}

// Info prints log with "[INFO]" prefix
func (l *SimpleLogger) Info(msg string) {
	l.print(msg, info, green)
}

// Error prints log with "[ERROR]" prefix
func (l *SimpleLogger) Error(msg string) {
	l.print(fmt.Sprintf("%s -> %s", funcCaller(), msg), err, red) // print file and line
}

// funcCaller returns file and line of last call.
func funcCaller() string {
	_, f, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s line %d", f, line)
}
