package automata

import (
	"github.com/neuralchecker/go-automata/base_types/state"
	"github.com/neuralchecker/go-automata/interfaces"
)

// FiniteAutomaton is an interface for finite automata.
type FiniteAutomaton[T any] interface {
	// GetName returns the name of the automaton.
	GetName() string
	// SetName sets the name of the automaton.
	SetName(name string)
	// GetComparator returns the comparator of the automaton.
	GetComparator() AutomataComparator[T]
	// SetComparator sets the comparator of the automaton.
	SetComparator(comparator AutomataComparator[T])
	// GetAlphabet returns the alphabet of the automaton.
	GetAlphabet() interfaces.Alphabet[T]
	// SetAlphabet sets the alphabet of the automaton.
	SetAlphabet(alphabet interfaces.Alphabet[T])
	// GetHole returns the hole of the automaton.
	GetHole() *state.FAState[T]
	// SetHole sets the hole on every state of the automaton.
	SetHole(hole *state.FAState[T])
	// GetInitialStates returns the initial states of the automaton.
	GetInitialStates() []*state.FAState[T]
	// GetInitialState returns the initial state of the automaton. If the automaton has more than one initial state, an error will be
	// returned.
	GetInitialState() (*state.FAState[T], error)
	// HasFullAlphabet returns true if the alphabet in the automaton contains all the symbols in the actual alphabet.
	HasFullAlphabet() bool
	// IsDeterministic returns true if the automaton is deterministic.
	IsDeterministic() bool
	// GetExporters returns the exporters of the automaton.
	GetExporters() []Exporter[T]
	// AddExporter adds an exporter to the automaton.
	AddExporter(exporter Exporter[T])
	// RemoveExporter removes an exporter from the automaton.
	RemoveExporter(exporter Exporter[T])
	// RemoveAllExporters removes all exporters from the automaton.
	RemoveAllExporters()
	// Export exports the automaton to the given path. The path should be a directory, and the name of the file will be the
	// name of the automaton + the corresponding extension. If the automaton has no name, a default name will be generated depending on the type
	// of the automaton.
	Export(path string)
	Equals(other FiniteAutomaton[T]) bool
}
