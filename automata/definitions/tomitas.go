package definitions

import (
	"github.com/neuralchecker/go-automata/automata"
	"github.com/neuralchecker/go-automata/base_types/alphabets"
	"github.com/neuralchecker/go-automata/base_types/states"
	"github.com/neuralchecker/go-automata/base_types/symbols"
	"github.com/neuralchecker/go-automata/comparators"
	"github.com/neuralchecker/go-automata/interfaces"
)

var (
	tomitasAlphabet = alphabets.New([]interfaces.Symbol[rune]{
		symbols.RuneSymbol('a'),
		symbols.RuneSymbol('b'),
	})
	a = symbols.RuneSymbol('a')
	b = symbols.RuneSymbol('b')
)

func GetAllTomitas() []*automata.DeterministicFiniteAutomaton[rune] {
	return []*automata.DeterministicFiniteAutomaton[rune]{
		GetTomitas1(),
		GetTomitas2(),
		GetTomitas3(),
		GetTomitas4(),
		GetTomitas5(),
		GetTomitas6(),
		GetTomitas7(),
	}
}

func GetTomitas1() *automata.DeterministicFiniteAutomaton[rune] {
	states := []*states.FAState[rune]{
		states.NewFAState[rune]("State A", true),
		states.NewFAState[rune]("State B", false),
	}
	states[0].AddTransition(b, states[0])
	states[0].AddTransition(a, states[1])
	states[1].AddTransition(b, states[1])
	states[1].AddTransition(a, states[1])

	return automata.NewDFA[rune](
		"tomita's 1", tomitasAlphabet, states[0], states, &comparators.RandomWalk[rune]{})
}

func GetTomitas2() *automata.DeterministicFiniteAutomaton[rune] {
	states := []*states.FAState[rune]{
		states.NewFAState[rune]("State A", true),
		states.NewFAState[rune]("State B", false),
		states.NewFAState[rune]("State C", false),
	}
	states[0].AddTransition(b, states[1])
	states[0].AddTransition(a, states[2])
	states[1].AddTransition(b, states[2])
	states[1].AddTransition(a, states[0])
	states[2].AddTransition(b, states[2])
	states[2].AddTransition(a, states[2])

	return automata.NewDFA[rune](
		"tomita's 2", tomitasAlphabet, states[0], states, &comparators.RandomWalk[rune]{})
}

func GetTomitas3() *automata.DeterministicFiniteAutomaton[rune] {
	states := []*states.FAState[rune]{
		states.NewFAState[rune]("State A", true),
		states.NewFAState[rune]("State B", true),
		states.NewFAState[rune]("State C", true),
		states.NewFAState[rune]("State D", false),
		states.NewFAState[rune]("State E", false),
	}
	states[0].AddTransition(b, states[1])
	states[0].AddTransition(a, states[0])
	states[1].AddTransition(b, states[0])
	states[1].AddTransition(a, states[3])
	states[2].AddTransition(b, states[1])
	states[2].AddTransition(a, states[3])
	states[3].AddTransition(b, states[4])
	states[3].AddTransition(a, states[2])
	states[4].AddTransition(b, states[4])
	states[4].AddTransition(a, states[4])

	return automata.NewDFA[rune](
		"tomita's 3", tomitasAlphabet, states[0], states, &comparators.RandomWalk[rune]{})
}

func GetTomitas4() *automata.DeterministicFiniteAutomaton[rune] {
	states := []*states.FAState[rune]{
		states.NewFAState[rune]("State A", true),
		states.NewFAState[rune]("State B", true),
		states.NewFAState[rune]("State C", true),
		states.NewFAState[rune]("State D", false),
	}
	states[0].AddTransition(b, states[0])
	states[0].AddTransition(a, states[1])
	states[1].AddTransition(b, states[0])
	states[1].AddTransition(a, states[2])
	states[2].AddTransition(b, states[0])
	states[2].AddTransition(a, states[3])
	states[3].AddTransition(b, states[3])
	states[3].AddTransition(a, states[3])

	return automata.NewDFA[rune](
		"tomita's 4", tomitasAlphabet, states[0], states, &comparators.RandomWalk[rune]{})
}

func GetTomitas5() *automata.DeterministicFiniteAutomaton[rune] {
	states := []*states.FAState[rune]{
		states.NewFAState[rune]("State A", true),
		states.NewFAState[rune]("State B", false),
		states.NewFAState[rune]("State C", false),
		states.NewFAState[rune]("State D", false),
	}
	states[0].AddTransition(b, states[1])
	states[0].AddTransition(a, states[2])
	states[1].AddTransition(b, states[0])
	states[1].AddTransition(a, states[3])
	states[2].AddTransition(b, states[3])
	states[2].AddTransition(a, states[0])
	states[3].AddTransition(b, states[2])
	states[3].AddTransition(a, states[1])

	return automata.NewDFA[rune](
		"tomita's 5", tomitasAlphabet, states[0], states, &comparators.RandomWalk[rune]{})
}

func GetTomitas6() *automata.DeterministicFiniteAutomaton[rune] {
	states := []*states.FAState[rune]{
		states.NewFAState[rune]("State A", true),
		states.NewFAState[rune]("State B", false),
		states.NewFAState[rune]("State C", false),
	}
	states[0].AddTransition(b, states[1])
	states[0].AddTransition(a, states[2])
	states[1].AddTransition(b, states[2])
	states[1].AddTransition(a, states[0])
	states[2].AddTransition(b, states[0])
	states[2].AddTransition(a, states[1])

	return automata.NewDFA[rune](
		"tomita's 6", tomitasAlphabet, states[0], states, &comparators.RandomWalk[rune]{})
}

func GetTomitas7() *automata.DeterministicFiniteAutomaton[rune] {
	states := []*states.FAState[rune]{
		states.NewFAState[rune]("State A", true),
		states.NewFAState[rune]("State B", true),
		states.NewFAState[rune]("State C", true),
		states.NewFAState[rune]("State D", true),
		states.NewFAState[rune]("State E", false),
	}
	states[0].AddTransition(b, states[1])
	states[0].AddTransition(a, states[0])
	states[1].AddTransition(b, states[1])
	states[1].AddTransition(a, states[2])
	states[2].AddTransition(b, states[3])
	states[2].AddTransition(a, states[2])
	states[3].AddTransition(b, states[3])
	states[3].AddTransition(a, states[4])
	states[4].AddTransition(b, states[4])
	states[4].AddTransition(a, states[4])

	return automata.NewDFA[rune](
		"tomita's 7", tomitasAlphabet, states[0], states, &comparators.RandomWalk[rune]{})
}
