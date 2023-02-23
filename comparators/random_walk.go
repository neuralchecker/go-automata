package comparators

import (
	"errors"
	"math/rand"

	"github.com/hashicorp/go-multierror"
	"github.com/neuralchecker/go-automata/automata"
	"github.com/neuralchecker/go-automata/base_types/alphabets"
	"github.com/neuralchecker/go-automata/base_types/sequences"
	"github.com/neuralchecker/go-automata/interfaces"
)

type RandomWalk[T any] struct {
	// MaxSteps is the maximum number of steps to take in the random walk.
	// If MaxSteps is 0, then the random walk will continue until a counterexample is found.
	MaxSteps int
	// ResetProbability is the probability of resetting the random walk to the initial state.
	// If ResetProbability is 0, then the random walk will never reset.
	ResetProbability float64
	// Rand is the random number generator to use. If Rand is nil, then the default random number
	// generator will be used, keep in mind that the default random number generator is deterministic.
	Random *rand.Rand
}

func NewRandomWalk[T any](maxSteps int, resetProbability float64, random *rand.Rand) *RandomWalk[T] {
	return &RandomWalk[T]{
		MaxSteps:         maxSteps,
		ResetProbability: resetProbability,
		Random:           random,
	}
}

func (r *RandomWalk[T]) AreEquivalent(automaton1 automata.Automaton[T], automaton2 automata.Automaton[T]) (bool, error) {
	_, err := r.GetCounterexampleBetween(automaton1, automaton2)
	if err != nil {
		if errors.Is(err, ErrNoCounterexampleFound) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (r *RandomWalk[T]) GetCounterexampleBetween(a1 automata.Automaton[T], a2 automata.Automaton[T]) (interfaces.Sequence[T], error) {
	if r.Random == nil {
		// This magic number is irrelevant, it just needs to be deterministic.
		r.Random = rand.New(rand.NewSource(42))
	}
	alphabet, err := r.getAlphabet(a1, a2)
	if err != nil {
		return nil, err
	}
	steps := 0
	counterExample := sequences.New[T]()
	for steps < r.MaxSteps || r.MaxSteps == 0 {
		if r.Random.Float64() < r.ResetProbability {
			counterExample = sequences.New[T]()
		}
		if ok, err := r.equivalentOutput(a1, a2, counterExample); !ok || err != nil {
			if err != nil {
				return nil, err
			}
			return counterExample, nil
		}
		pos := r.Random.Int31n(int32(alphabet.Length()))
		symbol := alphabet.GetSymbolAt(int(pos))
		counterExample = counterExample.Append(symbol)
		steps++
	}

	return nil, ErrNoCounterexampleFound
}

func (r *RandomWalk[T]) getAlphabet(a1 automata.Automaton[T], a2 automata.Automaton[T]) (interfaces.Alphabet[T], error) {
	if a1.HasFullAlphabet() && a2.HasFullAlphabet() {
		if !a1.GetAlphabet().Equals(a2.GetAlphabet()) {
			return nil, ErrAlphabetsDiffer
		}
		return a1.GetAlphabet(), nil
	}
	if !a1.HasFullAlphabet() && a2.HasFullAlphabet() {
		if !a1.GetAlphabet().Equals(a2.GetAlphabet()) {
			return nil, ErrAlphabetsDiffer
		}
		return a2.GetAlphabet(), nil
	}
	if a1.HasFullAlphabet() && !a2.HasFullAlphabet() {
		if !a2.GetAlphabet().Equals(a1.GetAlphabet()) {
			return nil, ErrAlphabetsDiffer
		}
		return a1.GetAlphabet(), nil
	}
	// If both automata have partial alphabets, then we need to merge them.
	alphabet := alphabets.Join(a1.GetAlphabet(), a2.GetAlphabet())
	return alphabet, nil
}

func (r *RandomWalk[T]) equivalentOutput(a1 automata.Automaton[T], a2 automata.Automaton[T], sequence interfaces.Sequence[T]) (bool, error) {
	errResult := &multierror.Error{}

	a1Out, err1 := a1.Accepts(sequence)
	a2Out, err2 := a2.Accepts(sequence)
	errResult = multierror.Append(errResult, err1, err2)

	if errResult.ErrorOrNil() != nil {
		return false, errResult
	}
	return a1Out == a2Out, nil
}
