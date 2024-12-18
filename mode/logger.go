package mode

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

// Logger is the logger for the flag
type Logger struct {
	logger gologger.Logger
}

// NewLogger is the logger for the flag
func NewLogger(logger gologger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.NilLoggerError
	}

	return &Logger{logger: logger}, nil
}

// ModeFlagSet is the flag set for the mode
func (l *Logger) ModeFlagSet(mode *Flag) {
	// Check if the mode flag is nil
	if mode == nil {
		return
	}

	l.logger.LogMessage(gologger.NewLogMessage("mode flag set", gologgerstatus.StatusDebug, nil, mode.String()))
}
