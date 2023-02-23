package guards

import "github.com/neuralchecker/go-automata/interfaces"

type Equality[T any] struct {
	S interfaces.Symbol[T]
}

// IsSatisfied implements interfaces.Guard
func (g Equality[T]) IsSatisfied(symbol interfaces.Symbol[T]) bool {
	if g.S == nil {
		return false
	}
	return g.S.Equals(symbol)
}

// String implements interfaces.Guard
func (g Equality[T]) String() string {
	if g.S == nil {
		return "⊥"
	}
	return g.S.String()
}

func NewEquality[T any](s interfaces.Symbol[T]) interfaces.Guard[T] {
	return Equality[T]{
		S: s,
	}
}
