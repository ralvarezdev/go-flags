package mode

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

var (
	// LogModeMap is the map of the log mode
	LogModeMap = map[Mode]map[gologgerstatus.Status]bool{
		Debug: {
			gologgerstatus.Info:     true,
			gologgerstatus.Error:    true,
			gologgerstatus.Warning:  true,
			gologgerstatus.Debug:    true,
			gologgerstatus.Critical: true,
		},
		Dev: {
			gologgerstatus.Info:     true,
			gologgerstatus.Error:    true,
			gologgerstatus.Warning:  true,
			gologgerstatus.Debug:    false,
			gologgerstatus.Critical: true,
		},
		Prod: {
			gologgerstatus.Info:     true,
			gologgerstatus.Error:    true,
			gologgerstatus.Warning:  true,
			gologgerstatus.Debug:    false,
			gologgerstatus.Critical: true,
		},
	}
)

type (
	// Logger interface for the mode logger
	Logger interface {
		ShouldLog(status gologgerstatus.Status) bool
		RunIfShouldLog(status gologgerstatus.Status, fn func())
		Log(message *gologger.Message)
		Info(subheader string, details ...string)
		Error(subheader string, errors ...error)
		Debug(subheader string, details ...string)
		Critical(subheader string, details ...string)
		Warning(subheader string, details ...string)
	}

	// DefaultLogger is the default mode logger
	DefaultLogger struct {
		header string
		logger gologger.Logger
	}
)

// NewDefaultLogger creates a new default mode logger
func NewDefaultLogger(header string, logger gologger.Logger) (
	*DefaultLogger,
	error,
) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.ErrNilLogger
	}

	return &DefaultLogger{header: header, logger: logger}, nil
}

// Log logs a message
func (d *DefaultLogger) Log(message *gologger.Message) {
	// Check if the message is nil
	if message == nil {
		return
	}

	d.RunIfShouldLog(
		message.Status(), func() {
			d.logger.Log(message)
		},
	)
}

// ShouldLog checks if the log should be logged
func (d *DefaultLogger) ShouldLog(status gologgerstatus.Status) bool {
	return LogModeMap[ModeFlag.Mode()][status]
}

// RunIfShouldLog runs the function if the log should be logged
func (d *DefaultLogger) RunIfShouldLog(
	status gologgerstatus.Status,
	fn func(),
) {
	if d.ShouldLog(status) {
		fn()
	}
}

// Info logs an info message
func (d *DefaultLogger) Info(subheader string, details ...string) {
	d.RunIfShouldLog(
		gologgerstatus.Info, func() {
			d.logger.Info(
				d.header,
				subheader,
				&details,
			)
		},
	)
}

// Error logs an error message
func (d *DefaultLogger) Error(subheader string, errors ...error) {
	d.RunIfShouldLog(
		gologgerstatus.Error, func() {
			d.logger.Error(
				d.header,
				subheader,
				&errors,
			)
		},
	)
}

// Debug logs a debug message
func (d *DefaultLogger) Debug(subheader string, details ...string) {
	d.RunIfShouldLog(
		gologgerstatus.Debug, func() {
			d.logger.Debug(
				d.header,
				subheader,
				&details,
			)
		},
	)
}

// Critical logs a critical message
func (d *DefaultLogger) Critical(subheader string, details ...string) {
	d.RunIfShouldLog(
		gologgerstatus.Critical, func() {
			d.logger.Critical(
				d.header,
				subheader,
				&details,
			)
		},
	)
}

// Warning logs a warning message
func (d *DefaultLogger) Warning(subheader string, details ...string) {
	d.RunIfShouldLog(
		gologgerstatus.Warning, func() {
			d.logger.Warning(
				d.header,
				subheader,
				&details,
			)
		},
	)
}
