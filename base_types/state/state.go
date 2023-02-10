package state

import "github.com/neuralchecker/go-automata/interfaces"

type State[T any] interface {
	IsDeterministic() bool
	IsFinal() bool
	NextStateFor(symbol interfaces.Symbol[T]) (State[T], error)
	NextStatesFor(symbol interfaces.Symbol[T]) []State[T]
	GetName() string
}
