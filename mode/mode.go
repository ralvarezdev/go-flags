package mode

import (
	goflags "github.com/ralvarezdev/go-flags"
)

const (
	// FlagName is the name of the mode flag
	FlagName = "m"

	// FlagUsage is the usage of the mode flag
	FlagUsage = "Specify mode. Allowed values are: dev, prod, debug. Default is the dev mode"
)

type (
	// Mode is the environment mode
	Mode string

	// Flag is a custom flag type for mode
	Flag struct {
		goflags.Flag
	}
)

var (
	// Dev is the development mode
	Dev Mode = "dev"

	// Prod is the production mode
	Prod Mode = "prod"

	// Debug is the debug mode
	Debug Mode = "debug"
)

// NewFlag creates a new Flag with allowed values
func NewFlag(defaultValue string, allowed []string) *Flag {
	return &Flag{
		Flag: *goflags.NewFlag(defaultValue, allowed),
	}
}

// Mode returns the mode
func (f *Flag) Mode() Mode {
	if f.IsProd() {
		return Prod
	}
	if f.IsDev() {
		return Dev
	}
	return Debug
}

// IsDev returns true if it's the development mode
func (f *Flag) IsDev() bool {
	return f.Value() == string(Dev)
}

// IsProd returns true if it's the production mode
func (f *Flag) IsProd() bool {
	return f.Value() == string(Prod)
}

// IsDebug returns true if it's the debug mode
func (f *Flag) IsDebug() bool {
	return f.Value() == string(Debug)
}

// ModeFlag is the environment mode
var ModeFlag = NewFlag(
	string(Dev),
	[]string{string(Dev), string(Prod), string(Debug)},
)

// init initializes the mode flag
func init() {
	goflags.SetFlag(ModeFlag, FlagName, FlagUsage)
}
