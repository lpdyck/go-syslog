// +build windows plan9 nacl

package gsyslog

import (
	"fmt"
	"log"
)

// NewLogger will return a syslog mock that will write to stderr with a default
// priority level and include the facility and tag. Used in unsupported systems
// windows, plan9, nacl
func NewLogger(p Priority, facility, tag string) (Syslogger, error) {
	return &unsupportedLogger{defaultLevel: p, facility: facility, tag: tag}, nil
}

// DialLogger is the same as NewLogger for unsupported systems
func DialLogger(network, raddr string, p Priority, facility, tag string) (Syslogger, error) {
	return nil, fmt.Errorf("Platform does not support syslog")
}

type unsupportedLogger struct {
	defaultLevel  Priority
	facility, tag string
}

func (l *unsupportedLogger) Write(b []byte) (int, error) {
	writeOutput(l.defaultLevel, l.facility, l.tag, b)
	return 0, nil
}

func (l *unsupportedLogger) WriteLevel(p Priority, b []byte) error {
	writeOutput(p, l.facility, l.tag, b)
	return nil
}

func (l *unsupportedLogger) Close() error {
	return nil
}

func writeOutput(p Priority, facility, tag string, b []byte) {
	s := getPriorityString(p)
	log.Printf("%s: %s", s, b)
}

func getPriorityString(p Priority) (s string) {
	switch p {
	case LOG_EMERG:
		s = "EMERGENCY:"
	case LOG_ALERT:
		s = "ALERT:"
	case LOG_CRIT:
		s = "CRIT:"
	case LOG_ERR:
		s = "ERROR:"
	case LOG_WARNING:
		s = "WARNING"
	case LOG_NOTICE:
		s = "NOTICE"
	case LOG_INFO:
		s = "INFO"
	case LOG_DEBUG:
		s = "DEBUG"
	default:
		s = ""
	}
	return s
}
