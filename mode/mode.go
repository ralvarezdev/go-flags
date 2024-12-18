package mode

import (
	"flag"
	goflags "github.com/ralvarezdev/go-flags"
)

const (
	// Name is the name of the mode flag
	Name = "m"

	// Dev is the development mode
	Dev = "dev"

	// Prod is the production mode
	Prod = "prod"

	// Debug is the debug mode
	Debug = "debug"
)

// Flag is a custom flag type for mode
type Flag struct {
	goflags.Flag
}

// NewFlag creates a new Flag with allowed values
func NewFlag(defaultValue string, allowed []string) *Flag {
	return &Flag{
		Flag: *goflags.NewFlag(defaultValue, allowed),
	}
}

// IsDev returns true if it's the development mode
func (m *Flag) IsDev() bool {
	return m.Value() == Dev
}

// IsProd returns true if it's the production mode
func (m *Flag) IsProd() bool {
	return m.Value() == Prod
}

// IsDebug returns true if it's the debug mode
func (m *Flag) IsDebug() bool {
	return m.Value() == Debug
}

// Mode is the environment mode
var Mode = NewFlag(Dev, []string{Dev, Prod, Debug})

// SetFlag sets the mode flag
func SetFlag() {
	flag.Var(
		Mode,
		Name,
		"Specify mode. Allowed values are: dev, prod, debug. Default is the development mode",
	)
}
