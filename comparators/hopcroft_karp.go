package comparators

import (
	"errors"

	"github.com/neuralchecker/go-automata/automata"
	"github.com/neuralchecker/go-automata/interfaces"
)

type HopcroftKarp[T any] struct{}

var _ automata.Comparator[int] = &HopcroftKarp[int]{}

func (h HopcroftKarp[T]) AreEquivalent(automaton1 automata.Automaton[T], automaton2 automata.Automaton[T]) (bool, error) {
	_, err := h.GetCounterexampleBetween(automaton1, automaton2)
	if err != nil {
		if errors.Is(err, ErrNoCounterexampleFound) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (h HopcroftKarp[T]) GetCounterexampleBetween(automaton1 automata.Automaton[T], automaton2 automata.Automaton[T]) (interfaces.Sequence[T], error) {
	panic("unimplemented")
}
