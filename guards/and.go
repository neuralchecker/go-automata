package guards

import (
	"strings"

	"github.com/neuralchecker/go-automata/interfaces"
)

// And is a guard that is satisfied if all of its guards are satisfied.
type And[T any] struct {
	guards []interfaces.Guard[T]
}

// IsSatisfied implements interfaces.Guard
func (g And[T]) IsSatisfied(symbol interfaces.Symbol[T]) bool {
	for _, guard := range g.guards {
		if !guard.IsSatisfied(symbol) {
			return false
		}
	}
	return true
}

// String implements interfaces.Guard
func (g And[T]) String() string {
	strBuilder := strings.Builder{}
	for i, guard := range g.guards {
		strBuilder.WriteString(guard.String())
		if i < len(g.guards)-1 {
			strBuilder.WriteString(" ∧ ")
		}
	}
	return strBuilder.String()
}

func NewAnd[T any](gs ...interfaces.Guard[T]) interfaces.Guard[T] {
	return And[T]{
		guards: gs,
	}
}
