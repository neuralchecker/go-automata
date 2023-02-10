package state

import "github.com/neuralchecker/go-automata/interfaces"

// State represents a state in a finite automaton. While an empty state is
// possible, it is still recommended to use the New function even to create
// an empty state.
type FAState[T any] struct {
	name            string
	transitions     map[interfaces.Symbol[T]][]*FAState[T]
	isFinal         bool
	isDeterministic bool
	hole            *FAState[T]
}

func NewFAState[T any](name string, isFinal bool) *FAState[T] {
	return &FAState[T]{
		name:            name,
		transitions:     make(map[interfaces.Symbol[T]][]*FAState[T]),
		isFinal:         isFinal,
		isDeterministic: true,
		hole:            nil,
	}
}

// AddHoleTransition adds a transition to the hole state. If the state already has a transition for the given symbol, the state will
// be overridden.
func (s *FAState[T]) AddHoleTransition(hole *FAState[T]) {
	s.hole = hole
}

// AddTransition adds a transition to the state. If the state already has a transition for the given symbol, the state will becom
// non-deterministic.
func (s *FAState[T]) AddTransition(symbol interfaces.Symbol[T], next ...*FAState[T]) {
	s.transitions[symbol] = append(s.transitions[symbol], next...)
	if len(s.transitions[symbol]) > 1 {
		s.isDeterministic = false
	}
}

func (s *FAState[T]) Equals(other *FAState[T]) bool {
	return s.name == other.GetName()
}

// IsDeterministic returns true if the state has no more than one transition for each symbol in the alphabet.
func (s *FAState[T]) IsDeterministic() bool {
	return s.isDeterministic
}

// IsFinal returns true if the state is final.
func (s *FAState[T]) IsFinal() bool {
	return s.isFinal
}

// NextState returns the next state for the given symbol. If the state is non-deterministic, an error will be returned along with one of
// the next states, chosen arbitrarily.
func (s *FAState[T]) NextStateFor(symbol interfaces.Symbol[T]) (State[T], error) {
	nextStates := s.NextStatesFor(symbol)
	if len(nextStates) > 1 {
		return nextStates[0], &ErrNonDeterministicState{symbol: symbol.GetValue()}
	}
	return nextStates[0], nil
}

// NextStates returns the next states for the given symbol.
func (s *FAState[T]) NextStatesFor(symbol interfaces.Symbol[T]) []State[T] {
	nextStates := make([]State[T], len(s.transitions[symbol]))
	transitions := s.transitions[symbol]
	for i := range transitions {
		nextStates[i] = s.transitions[symbol][i]
	}

	if len(nextStates) == 0 {
		return []State[T]{s.hole}
	}

	return nextStates
}

// String returns the name of the state. Same as GetName().
func (s *FAState[T]) String() string {
	return s.GetName()
}

// GetName returns the name of the state. Same as String().
func (s *FAState[T]) GetName() string {
	return s.name
}
