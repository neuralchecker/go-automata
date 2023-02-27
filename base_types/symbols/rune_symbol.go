package symbols

import (
	"fmt"

	"github.com/neuralchecker/go-automata/interfaces"
)

type RuneSymbol rune

func (r RuneSymbol) AddToValue(value rune) interfaces.Symbol[rune] {
	return RuneSymbol(rune(r) + value)
}

func (r RuneSymbol) Equal(other interfaces.Symbol[rune]) bool {
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

func (r RuneSymbol) Hash() int {
	return int(r)
}
