package automata

import (
	"github.com/neuralchecker/go-automata/base_types/state"
	"github.com/neuralchecker/go-automata/interfaces"
)

type Automaton[T any] interface {
	// GetName returns the name of the automaton.
	GetName() string
	// GetAlphabet returns the alphabet of the automaton.
	GetAlphabet() interfaces.Alphabet[T]
	GetInitialStates() []state.State[T]
	GetHole() state.State[T]
	IsDeterministic() bool
	HasFullAlphabet() bool
}

type AutomataComparator[T any] interface {
	// AreEquivalent returns true if the two automata are equivalent, false otherwise.
	AreEquivalent(automaton1 Automaton[T], automaton2 Automaton[T]) bool
	// GetCounterexampleBetween returns a counterexample between the two automata.
	GetCounterexampleBetween(automaton1 Automaton[T], automaton2 Automaton[T]) interfaces.Sequence[T]
	// HasEquivalentOutput returns true if the two automata have equivalent output, false otherwise.
	HasEquivalentOutput(automaton1 Automaton[T], automaton2 Automaton[T]) bool
}
