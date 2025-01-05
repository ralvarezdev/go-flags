package go_mode_flag

import (
	"flag"
	"fmt"
	"strings"
)

// Flag is a custom flag type
type Flag struct {
	value   string
	allowed []string
}

// NewFlag creates a new Flag
func NewFlag(value string, allowed []string) *Flag {
	return &Flag{
		value:   value,
		allowed: allowed,
	}
}

// String returns the string representation of the flag value
func (f *Flag) String() string {
	return f.value
}

// Value returns the flag value
func (f *Flag) Value() string {
	return f.value
}

// Allowed returns the allowed values
func (f *Flag) Allowed() []string {
	return f.allowed
}

// Set validates and sets the flag value
func (f *Flag) Set(value string) error {
	for _, v := range f.allowed {
		if value == v {
			f.value = value
			return nil
		}
	}
	return fmt.Errorf(
		"invalid value %q, allowed values are: %s", value,
		strings.Join(f.allowed, ", "),
	)
}

// SetFlag sets the mode flag
func SetFlag(value flag.Value, name string, usage string) {
	flag.Var(
		value,
		name,
		usage,
	)
}
