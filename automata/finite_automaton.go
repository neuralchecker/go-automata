package automata

import (
	"github.com/neuralchecker/go-automata/base_types/states"
	"github.com/neuralchecker/go-automata/interfaces"
)

// FiniteAutomaton is an interface for finite automata.
type FiniteAutomaton[T any] interface {
	Automaton[T]
	// SetName sets the name of the automaton.
	SetName(name string)
	// GetComparator returns the comparator of the automaton.
	GetComparator() Comparator[T]
	// SetComparator sets the comparator of the automaton.
	SetComparator(comparator Comparator[T])
	// SetAlphabet sets the alphabet of the automaton.
	SetAlphabet(alphabet interfaces.Alphabet[T])
	// GetHoleAsFAState returns the hole of the automaton as a FAState.
	GetHoleAsFAState() *states.FAState[T]
	// SetHole sets the hole on every state of the automaton.
	SetHole(hole *states.FAState[T])
	// GetInitialStates returns the initial states of the automaton.
	GetInitialStates() []states.State[T]
	// GetInitialStatesAsFAStates returns the initial states of the automaton as a slice of *states.FAState[T].
	GetInitialStatesAsFAStates() []*states.FAState[T]
	// GetStates returns the states of the automaton.
	GetStates() []states.State[T]
	// GetStatesAsFAStates returns the states of the automaton as a slice of *states.FAState[T].
	GetStatesAsFAStates() []*states.FAState[T]
	// GetInitialStateAsFAState returns the initial state of the automaton as a FAState. If the automaton has more than one initial state, an error will be
	// returned.
	GetInitialStateAsFAState() (*states.FAState[T], error)
	// IsDeterministic returns true if the automaton is deterministic.
	IsDeterministic() bool
	// GetExporters returns the exporters of the automaton.
	GetExporters() []Exporter[T]
	// AddExporter adds an exporter to the automaton.
	AddExporter(exporter ...Exporter[T])
	// RemoveExporter removes an exporter from the automaton.
	RemoveExporter(exporter ...Exporter[T])
	// RemoveAllExporters removes all exporters from the automaton.
	RemoveAllExporters()
	// Export exports the automaton to the given path. The path should be a directory, and the name of the file will be the
	// name of the automaton + the corresponding extension. If the automaton has no name, a default name will be generated depending on the type
	// of the automaton.
	Export(path string) error
	Equal(other Automaton[T]) bool
}

type Automaton[T any] interface {
	// GetName returns the name of the automaton.
	GetName() string
	// GetAlphabet returns the alphabet of the automaton.
	GetAlphabet() interfaces.Alphabet[T]
	// GetInitialStates returns the initial states of the automaton.
	GetInitialStates() []states.State[T]
	// GetHole returns the hole of the automaton.
	GetHole() states.State[T]
	// IsDeterministic returns true if the automaton is deterministic.
	IsDeterministic() bool
	// HasFullAlphabet returns true if the alphabet in the automaton contains all the symbols in the actual alphabet.
	HasFullAlphabet() bool
	// Step performs a step in the automaton with the given symbol. If the automaton is deterministic, the step will be performed
	Step(symbol interfaces.Symbol[T]) (bool, error)
	// Accepts returns true if the automaton accepts the given sequence, false otherwise.
	Accepts(sequence interfaces.Sequence[T]) (bool, error)
	// Reset resets the automaton.
	Reset() error
	// GetStates returns the states of the automaton.
	GetStates() []states.State[T]
}

type Comparator[T any] interface {
	// AreEquivalent returns true if the two automata are equivalent, false otherwise.
	AreEquivalent(automaton1 Automaton[T], automaton2 Automaton[T]) (bool, error)
	// GetCounterexampleBetween returns a counterexample between the two automata.
	GetCounterexampleBetween(automaton1 Automaton[T], automaton2 Automaton[T]) (interfaces.Sequence[T], error)
	// HasEquivalentOutput returns true if the two automata have equivalent output, false otherwise.
}
