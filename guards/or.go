package guards

import (
	"strings"

	"github.com/neuralchecker/go-automata/interfaces"
)

// Or is a guard that is satisfied if any of its guards are satisfied.
type Or[T any] struct {
	guards []interfaces.Guard[T]
}

// IsSatisfied implements interfaces.Guard
func (g Or[T]) IsSatisfied(symbol interfaces.Symbol[T]) bool {
	for _, guard := range g.guards {
		if guard.IsSatisfied(symbol) {
			return true
		}
	}
	return false
}

// String implements interfaces.Guard
func (g Or[T]) String() string {
	strBuilder := strings.Builder{}
	for i, guard := range g.guards {
		strBuilder.WriteString(guard.String())
		if i < len(g.guards)-1 {
			strBuilder.WriteString(" ∨ ")
		}
	}
	return strBuilder.String()
}

func NewOrGuard[T any](gs ...interfaces.Guard[T]) interfaces.Guard[T] {
	return Or[T]{
		guards: gs,
	}
}
