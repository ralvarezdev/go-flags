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

	// Setup is the setup mode.
	Setup Mode = "setup"

	// DefaultMode is the default mode (development).
	DefaultMode = Dev

	// AllowedModes is the list of allowed modes.
	AllowedModes = []Mode{
		Dev,
		Prod,
		Debug,
		Migrate,
		Setup,
	}
)

// NewFlag creates a new Flag with allowed values.
//
// Parameters:
//
//	defaultMode - the default mode.
//	allowed - slice of allowed string values.
//
// Returns:
//
//	A pointer to the created Flag.
func NewFlag(
	defaultMode Mode,
	allowed []Mode,
) *Flag {
	// Convert allowed modes to a slice of strings
	allowedModesStr := make([]string, len(allowed))
	for i, mode := range allowed {
		allowedModesStr[i] = string(mode)
	}

	// Convert the default mode to string
	defaultModeStr := string(defaultMode)

	return &Flag{
		Flag: *goflags.NewFlag(
			&defaultModeStr,
			allowedModesStr,
			FlagName,
			fmt.Sprintf(
				FlagUsage,
				strings.Join(allowedModesStr, ", "),
				defaultModeStr,
			),
		),
	}
}

// Default returns the default value of the flag.
//
// Returns:
//
//	The default value.
func (f *Flag) Default() string {
	if f == nil {
		return ""
	}
	return f.Default()
}

// IsDev returns true if the current mode is development.
//
// Returns:
//
//	True if development mode, false otherwise.
func (f *Flag) IsDev() bool {
	if f == nil {
		return false
	}
	return f.Value() == string(Dev)
}

// IsProd returns true if the current mode is production.
//
// Returns:
//
//	True if production mode, false otherwise.
func (f *Flag) IsProd() bool {
	if f == nil {
		return false
	}
	return f.Value() == string(Prod)
}

// IsDebug returns true if the current mode is debug.
//
// Returns:
//
//	True if debug mode, false otherwise.
func (f *Flag) IsDebug() bool {
	if f == nil {
		return false
	}
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

// IsSetup returns true if the current mode is setup.
//
// Returns:
//
//	True if setup mode, false otherwise.
func (f *Flag) IsSetup() bool {
	if f == nil {
		return false
	}
	return f.Value() == string(Setup)
}

// SetFlag initializes the mode flag
func SetFlag(flag *Flag) {
	if flag != nil {
		flag.SetFlag()
	}
}
