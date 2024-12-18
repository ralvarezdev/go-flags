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

// IsDev returns true if the mode is development
func (m *Flag) IsDev() bool {
	return m.Value() == Dev
}

// IsProd returns true if the mode is production
func (m *Flag) IsProd() bool {
	return m.Value() == Prod
}

// Mode is the environment mode
var Mode = NewFlag(Dev, []string{Dev, Prod})

// SetFlag sets the mode flag
func SetFlag() {
	flag.Var(
		Mode,
		Name,
		"Specify mode. Allowed values are: dev, prod. Default is the development mode",
	)
}
