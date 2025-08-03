package mode

import (
	"fmt"
	"strings"

	goflags "github.com/ralvarezdev/go-flags"
)

type (
	// Mode represents the environment mode as a string.
	Mode string

	// Flag is a custom flag type for mode, embedding goflags.Flag.
	Flag struct {
		goflags.Flag
	}
)

var (
	// Dev is the development mode.
	Dev Mode = "dev"

	// Prod is the production mode.
	Prod Mode = "prod"

	// Debug is the debug mode.
	Debug Mode = "debug"

	// Migrate is the migration mode.
	Migrate Mode = "migrate"

	// DefaultMode is the default mode (development).
	DefaultMode = Dev

	// AllowedModes is the list of allowed modes.
	AllowedModes = []string{
		string(Dev), string(Prod), string(Debug), string(Migrate),
	}

	// ModeFlag is the environment mode flag instance.
	ModeFlag = NewFlag(
		string(DefaultMode),
		nil,
		AllowedModes,
		FlagName,
		fmt.Sprintf(
			FlagUsage,
			strings.Join(AllowedModes, ", "),
			string(Dev),
		),
	)
)

// NewFlag creates a new Flag with allowed values.
//
// Parameters:
//
//	defaultValue - the default value for the flag.
//	value - pointer to a string to store the flag value (can be nil).
//	allowed - slice of allowed string values.
//	name - the flag name.
//	usage - the usage description for the flag.
//
// Returns:
//
//	A pointer to the created Flag.
func NewFlag(
	defaultValue string,
	value *string,
	allowed []string,
	name string,
	usage string,
) *Flag {
	return &Flag{
		Flag: *goflags.NewFlag(defaultValue, value, allowed, name, usage),
	}
}

// Default returns the default value of the flag.
//
// Returns:
//
//	The default value.
func (f *Flag) Default() string {
	return f.Default()
}

// IsDev returns true if the current mode is development.
//
// Returns:
//
//	True if development mode, false otherwise.
func (f *Flag) IsDev() bool {
	return f.Value() == string(Dev)
}

// IsProd returns true if the current mode is production.
//
// Returns:
//
//	True if production mode, false otherwise.
func (f *Flag) IsProd() bool {
	return f.Value() == string(Prod)
}

// IsDebug returns true if the current mode is debug.
//
// Returns:
//
//	True if debug mode, false otherwise.
func (f *Flag) IsDebug() bool {
	return f.Value() == string(Debug)
}

// IsMigrate returns true if the current mode is migration.
//
// Returns:
//
//	True if migration mode, false otherwise.
func (f *Flag) IsMigrate() bool {
	return f.Value() == string(Migrate)
}

// init initializes the mode flag.
func init() {
	ModeFlag.SetFlag()
}
