package states

import (
	"fmt"
	"strings"

	"github.com/neuralchecker/go-automata/interfaces"
)

type SymbolicState[T any] struct {
	name        string
	transitions []Pair[interfaces.Guard[T], State[T]]
	isFinal     bool
	hole        *SymbolicState[T]
}

func (s *SymbolicState[T]) GetHole() *SymbolicState[T] {
	return s.hole
}

func (s *SymbolicState[T]) SetHole(hole *SymbolicState[T]) {
	s.hole = hole
}

func (s *SymbolicState[T]) AddTransition(guard interfaces.Guard[T], next *SymbolicState[T]) {
	s.transitions = append(
		s.transitions,
		Pair[interfaces.Guard[T], State[T]]{
			Fst: guard,
			Snd: next,
		},
	)
}

// IsDeterministic always returns true. While may not always be true, symbolic state as it stands right
// now is always supposed to be deterministic.
func (s *SymbolicState[T]) IsDeterministic() bool {
	return true
}

func (s *SymbolicState[T]) IsFinal() bool {
	return s.isFinal
}

func (s *SymbolicState[T]) NextStateFor(symbol interfaces.Symbol[T]) (State[T], error) {
	for _, transition := range s.transitions {
		if transition.Fst.IsSatisfied(symbol) {
			return transition.Snd, nil
		}
	}
	if s.hole != nil {
		return s.hole, nil
	}
	return nil, &ErrNoTransition{symbol: symbol.GetValue()}

}

func (s *SymbolicState[T]) NextStatesFor(symbol interfaces.Symbol[T]) []State[T] {
	next, err := s.NextStateFor(symbol)
	//Since this method is not really supposed to be used, we can just return an empty slice
	if err != nil {
		return []State[T]{}
	}
	return []State[T]{next}
}

func (s *SymbolicState[T]) GetName() string {
	return s.name
}

func (s *SymbolicState[T]) GetTransitions() []Pair[fmt.Stringer, State[T]] {
	transitions := make([]Pair[fmt.Stringer, State[T]], len(s.transitions))
	for i, transition := range s.transitions {
		transitions[i] = Pair[fmt.Stringer, State[T]]{
			Fst: transition.Fst,
			Snd: transition.Snd,
		}
	}
	return transitions
}

func (s *SymbolicState[T]) String() string {
	return strings.ReplaceAll(s.GetName(), " ", "_")
}
