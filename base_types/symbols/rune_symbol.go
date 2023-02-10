package symbols

import (
	"fmt"

	"github.com/neuralchecker/go-automata/interfaces"
)

type RuneSymbol rune

// This is a type assertion to ensure that RuneSymbol implements the Symbol[rune] interface.
var _ interfaces.Symbol[rune] = RuneSymbol(0)

func (r RuneSymbol) AddToValue(value rune) interfaces.Symbol[rune] {
	return RuneSymbol(rune(r) + value)
}

func (r RuneSymbol) Equals(other interfaces.Symbol[rune]) bool {
	return rune(r) == other.GetValue()
}

func (r RuneSymbol) GetValue() rune {
	return rune(r)
}

func (r RuneSymbol) GreaterEqualThan(other interfaces.Symbol[rune]) bool {
	return rune(r) >= other.GetValue()
}

func (r RuneSymbol) GreaterThan(other interfaces.Symbol[rune]) bool {
	return rune(r) > other.GetValue()
}

func (r RuneSymbol) LesserEqualThan(other interfaces.Symbol[rune]) bool {
	return rune(r) <= other.GetValue()
}

func (r RuneSymbol) LesserThan(other interfaces.Symbol[rune]) bool {
	return rune(r) < other.GetValue()
}

func (r RuneSymbol) String() string {
	return fmt.Sprintf("%c", rune(r))
}
