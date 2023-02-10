package automata

import "errors"

// Errors
var (
	ErrInitialStateNotSet = errors.New("initial state not set")
)

// Error types

// ErrInvalidStateType is returned when the state type is not valid.
type ErrInvalidStateType struct {
	// Expected is the expected state type.
	Expected string
	// Actual is the actual state type.
	Actual string
}

func (e *ErrInvalidStateType) Error() string {
	return "expected state type: " + e.Expected + ", but got: " + e.Actual
}
