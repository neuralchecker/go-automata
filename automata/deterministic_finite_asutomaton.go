package automata

import (
	"errors"
	"reflect"

	"github.com/neuralchecker/go-automata/base_types/states"
	"github.com/neuralchecker/go-automata/interfaces"
)

type DeterministicFiniteAutomaton[T any] struct {
	States              []*states.FAState[T]
	Alphabet            interfaces.Alphabet[T]
	InitialState        *states.FAState[T]
	Hole                *states.FAState[T]
	ComparingStrategy   Comparator[T]
	ActualState         *states.FAState[T]
	ExportingStrategies []Exporter[T]
	Name                string
}

var _ FiniteAutomaton[int] = &DeterministicFiniteAutomaton[int]{}

func NewDeterministicFiniteAutomaton[T any](name string, alphabet interfaces.Alphabet[T], initialState *states.FAState[T], states []*states.FAState[T],
	comparator Comparator[T],
) *DeterministicFiniteAutomaton[T] {
	return &DeterministicFiniteAutomaton[T]{
		States:              states,
		Alphabet:            alphabet,
		InitialState:        initialState,
		ComparingStrategy:   comparator,
		ActualState:         initialState,
		ExportingStrategies: make([]Exporter[T], 0),
		Name:                name,
	}
}

// Shorthand for NewDeterministicFiniteAutomaton
func NewDFA[T any](name string, alphabet interfaces.Alphabet[T], initialState *states.FAState[T], states []*states.FAState[T],
	comparator Comparator[T],
) *DeterministicFiniteAutomaton[T] {
	return NewDeterministicFiniteAutomaton(name, alphabet, initialState, states, comparator)
}

func NewDeterministicFiniteAutomatonWithHole[T any](name string, alphabet interfaces.Alphabet[T], initialState *states.FAState[T], states []*states.FAState[T],
	comparator Comparator[T], hole *states.FAState[T],
) *DeterministicFiniteAutomaton[T] {
	dfa := NewDeterministicFiniteAutomaton(name, alphabet, initialState, states, comparator)
	dfa.SetHole(hole)
	return dfa
}

// Shorthand for NewDeterministicFiniteAutomatonWithHole
func NewDFAWithHole[T any](name string, alphabet interfaces.Alphabet[T], initialState *states.FAState[T], states []*states.FAState[T],
	comparator Comparator[T], hole *states.FAState[T],
) *DeterministicFiniteAutomaton[T] {
	return NewDeterministicFiniteAutomatonWithHole(name, alphabet, initialState, states, comparator, hole)
}

func (dfa *DeterministicFiniteAutomaton[T]) GetName() string {
	return dfa.Name
}

func (dfa *DeterministicFiniteAutomaton[T]) SetName(name string) {
	dfa.Name = name
}

func (dfa *DeterministicFiniteAutomaton[T]) GetComparator() Comparator[T] {
	return dfa.ComparingStrategy
}

func (dfa *DeterministicFiniteAutomaton[T]) SetComparator(comparator Comparator[T]) {
	dfa.ComparingStrategy = comparator
}

func (dfa *DeterministicFiniteAutomaton[T]) GetAlphabet() interfaces.Alphabet[T] {
	return dfa.Alphabet
}

func (dfa *DeterministicFiniteAutomaton[T]) SetAlphabet(alphabet interfaces.Alphabet[T]) {
	dfa.Alphabet = alphabet
}

func (dfa *DeterministicFiniteAutomaton[T]) HasFullAlphabet() bool {
	return dfa.Alphabet.IsComplete()
}

// GetStates returns the states of the automaton as a slice of state.State[T]. If you
// want to get the states as a slice of *states.FAState[T], use GetStatesAsFAStates.
func (dfa *DeterministicFiniteAutomaton[T]) GetStates() []states.State[T] {
	states := make([]states.State[T], len(dfa.States))
	for i := range dfa.States {
		states[i] = dfa.States[i]
	}
	return states
}

// GetStatesAsFAStates returns the states of the automaton as a slice of *states.FAState[T].
func (dfa *DeterministicFiniteAutomaton[T]) GetStatesAsFAStates() []*states.FAState[T] {
	return dfa.States
}

func (dfa *DeterministicFiniteAutomaton[T]) SetStates(states []*states.FAState[T]) {
	dfa.States = states
}

func (dfa *DeterministicFiniteAutomaton[T]) GetInitialState() (states.State[T], error) {
	return dfa.GetInitialStateAsFAState()
}

func (dfa *DeterministicFiniteAutomaton[T]) GetInitialStateAsFAState() (*states.FAState[T], error) {
	if dfa.InitialState == nil {
		return nil, ErrInitialStateNotSet
	}
	return dfa.InitialState, nil
}

func (dfa *DeterministicFiniteAutomaton[T]) GetInitialStates() []states.State[T] {
	return []states.State[T]{dfa.InitialState}
}

func (dfa *DeterministicFiniteAutomaton[T]) GetInitialStatesAsFAStates() []*states.FAState[T] {
	return []*states.FAState[T]{dfa.InitialState}
}

func (dfa *DeterministicFiniteAutomaton[T]) SetInitialState(initialState *states.FAState[T]) {
	dfa.InitialState = initialState
}

func (dfa *DeterministicFiniteAutomaton[T]) GetHole() states.State[T] {
	return dfa.Hole
}

func (dfa *DeterministicFiniteAutomaton[T]) GetHoleAsFAState() *states.FAState[T] {
	return dfa.Hole
}

func (dfa *DeterministicFiniteAutomaton[T]) SetHole(hole *states.FAState[T]) {
	dfa.Hole = hole
	for _, state := range dfa.States {
		state.AddHoleTransition(hole)
	}
}

func (dfa *DeterministicFiniteAutomaton[T]) Step(symbol interfaces.Symbol[T]) (bool, error) {
	if err := dfa.Reset(); err != nil {
		return false, err
	}

	actualState, err := dfa.step(dfa.ActualState, symbol)
	if err != nil {
		return false, err
	}
	dfa.ActualState = actualState
	return true, nil
}

func (dfa *DeterministicFiniteAutomaton[T]) Reset() error {
	if dfa.InitialState == nil {
		return ErrInitialStateNotSet
	}
	dfa.ActualState = dfa.InitialState
	return nil
}

func (dfa *DeterministicFiniteAutomaton[T]) Accepts(sequence interfaces.Sequence[T]) (bool, error) {
	var err error
	currentState := dfa.InitialState
	if err = dfa.Reset(); err != nil {
		return false, err
	}
	it := sequence.Iterator()
	for it.HasNext() {
		symbol := it.Next()
		currentState, err = dfa.step(currentState, symbol)
		if err != nil {
			return false, err
		}
	}
	return currentState.IsFinal(), nil
}

func (dfa *DeterministicFiniteAutomaton[T]) step(actualState *states.FAState[T], symbol interfaces.Symbol[T]) (*states.FAState[T], error) {
	newState, err := actualState.NextStateFor(symbol)
	if err != nil {
		return nil, err
	}
	newFAState, ok := newState.(*states.FAState[T])
	// This should never happen, because each state in the automaton is a *states.FAState[T]
	if !ok {
		return nil, &ErrInvalidStateType{
			Expected: reflect.TypeOf(newState).String(),
			Actual:   reflect.TypeOf(&states.FAState[T]{}).String(),
		}
	}
	return newFAState, nil
}

func (dfa *DeterministicFiniteAutomaton[T]) AddExporter(exporter ...Exporter[T]) {
	dfa.ExportingStrategies = append(dfa.ExportingStrategies, exporter...)
}

func (dfa *DeterministicFiniteAutomaton[T]) RemoveExporter(exporters ...Exporter[T]) {
	newExporters := make([]Exporter[T], 0, len(dfa.ExportingStrategies)-len(exporters))
	for _, exporter := range exporters {
		for i, dfaExporter := range dfa.ExportingStrategies {
			if reflect.TypeOf(dfaExporter) != reflect.TypeOf(exporter) {
				newExporters = append(newExporters, dfa.ExportingStrategies[i])
			}
		}
	}
	dfa.ExportingStrategies = newExporters
}

func (dfa *DeterministicFiniteAutomaton[T]) RemoveAllExporters() {
	dfa.ExportingStrategies = make([]Exporter[T], 0)
}

func (dfa *DeterministicFiniteAutomaton[T]) GetExporters() []Exporter[T] {
	return dfa.ExportingStrategies
}

func (dfa *DeterministicFiniteAutomaton[T]) Export(pathStr string) error {
	pathStr = ResolvePath[T](dfa, pathStr)
	var errResult error
	for _, exporter := range dfa.ExportingStrategies {
		err := exporter.Export(dfa, pathStr)
		if err != nil {
			errResult = errors.Join(errResult, err)
		}
	}
	return errResult
}

func (dfa *DeterministicFiniteAutomaton[T]) Equal(other Automaton[T]) bool {
	ret, _ := dfa.ComparingStrategy.AreEquivalent(dfa, other)
	return ret
}

func (dfa *DeterministicFiniteAutomaton[T]) IsDeterministic() bool {
	return true
}
