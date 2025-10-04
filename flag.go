package go_flags

import (
	"flag"
	"fmt"
	"strings"
)

type (
	// Flag is a custom flag type that restricts its value to a predefined set of allowed strings.
	Flag struct {
		defaultValue string
		value        *string
		allowed      []string
		name         string
		usage        string
	}
)

// NewFlag creates a new Flag with a default value and a list of allowed values.
//
// Parameters:
//
//	defaultValue - the default value for the flag.
//	value - the default value for the flag.
//	allowed - slice of allowed string values.
//	name - the name of the flag.
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
		defaultValue,
		value,
		allowed,
		name,
		usage,
	}
}

// Default returns the default value of the flag.
//
// Returns:
//
//	The default value of the flag as a string.
func (f *Flag) Default() string {
	if f == nil {
		return ""
	}
	return f.defaultValue
}

// String returns the string representation of the flag's current value.
//
// Returns:
//
//	The current value of the flag as a string.
func (f *Flag) String() string {
	if f == nil {
		return ""
	}
	if f.value != nil {
		return *f.value
	}
	return f.defaultValue
}

// Value returns the current value of the flag.
//
// Returns:
//
//	The current value of the flag as a string.
func (f *Flag) Value() string {
	if f == nil {
		return ""
	}
	return f.String()
}

// Allowed returns the list of allowed values for the flag.
//
// Returns:
//
//	The allowed values.
func (f *Flag) Allowed() []string {
	if f == nil {
		return nil
	}
	return f.allowed
}

// Set validates and sets the flag's value.
//
// Parameters:
//
//	value - the value to set.
//
// Returns:
//
//	An error if the value is not allowed, otherwise nil.
func (f *Flag) Set(value string) error {
	if f == nil {
		return ErrNilFlag
	}
	for _, v := range f.allowed {
		if value == v {
			f.value = &value
			return nil
		}
	}
	return fmt.Errorf(
		ErrInvalidValue, value,
		strings.Join(f.allowed, ", "),
	)
}

// SetFlag registers the custom flag with the standard flag package.
func (f *Flag) SetFlag() {
	SetFlag(f, f.name, f.usage)
}

// SetFlag registers the custom flag with the standard flag package.
//
// Parameters:
//
//	value - the flag.Value to register.
//	name - the flag name.
//	usage - the usage description for the flag.
func SetFlag(value flag.Value, name string, usage string) {
	flag.Var(
		value,
		name,
		usage,
	)
}
