package guards

import "github.com/neuralchecker/go-automata/interfaces"

type Negation[T any] struct {
	G interfaces.Guard[T]
}

func (g Negation[T]) IsSatisfied(symbol interfaces.Symbol[T]) bool {
	if g.G == nil {
		return true
	}
	return !g.G.IsSatisfied(symbol)
}

func (g Negation[T]) String() string {
	if g.G == nil {
		return "¬⊥"
	}
	return "¬" + g.G.String()
}

func NewNegation[T any](g interfaces.Guard[T]) interfaces.Guard[T] {
	return Negation[T]{
		G: g,
	}
}
