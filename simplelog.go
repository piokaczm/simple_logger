package simplelog

import (
	"log"
	"os"
)

// These are logger levels prefixes.
const (
	debug = "[DEBUG] "
	info  = "[INFO]  "
	err   = "[ERROR] "
)

// SimpleLogger is a struct which holds the standard logger
// along with it's prefixes.
type SimpleLogger struct {
	logger *log.Logger
	debug  string
	info   string
	err    string
}

// InitLogger returns logger with standarized log format to use throughout the app.
func InitLogger() *SimpleLogger {
	return &SimpleLogger{
		logger: log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds),
		debug:  debug,
		info:   info,
		err:    err,
	}
}

// Debug prints log with "[DEBUG]" prefix
func (l *SimpleLogger) Debug(msg string) {
	l.logger.SetPrefix(l.debug)
	l.logger.Println(msg)
}

// Info prints log with "[INFO]" prefix
func (l *SimpleLogger) Info(msg string) {
	l.logger.SetPrefix(l.info)
	l.logger.Println(msg)
}

// Error prints log with "[ERROR]" prefix
func (l *SimpleLogger) Error(msg string) {
	l.logger.SetPrefix(l.err)
	l.logger.Println(msg)
}
