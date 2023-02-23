package states

import (
	"fmt"

	"github.com/neuralchecker/go-automata/interfaces"
)

type State[T any] interface {
	IsDeterministic() bool
	IsFinal() bool
	NextStateFor(symbol interfaces.Symbol[T]) (State[T], error)
	NextStatesFor(symbol interfaces.Symbol[T]) []State[T]
	GetName() string
	GetTransitions() []Pair[fmt.Stringer, State[T]]
	String() string
}

type Pair[T1, T2 any] struct {
	Fst T1
	Snd T2
}
