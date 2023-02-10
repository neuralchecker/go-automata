package states

import (
	"fmt"
)

type ErrNonDeterministicState struct {
	symbol any
}

func (e *ErrNonDeterministicState) Error() string {
	return fmt.Sprintf("non deterministic state, multiple transitions for symbol %v", e.symbol)
}

type ErrNoTransition struct {
	symbol any
}

func (e *ErrNoTransition) Error() string {
	return fmt.Sprintf("no transition for symbol %v", e.symbol)
}
