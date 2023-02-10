package symbols

import "github.com/neuralchecker/go-automata/interfaces"

type StrSymbol struct {
	value string
}

// This is a type assertion to ensure that StrSymbol implements the Symbol[string] interface.
var _ interfaces.Symbol[string] = StrSymbol{}

func (s StrSymbol) AddToValue(value string) interfaces.Symbol[string] {
	return StrSymbol{value: s.value + value}
}

func (s StrSymbol) Equals(other interfaces.Symbol[string]) bool {
	return s.value == other.GetValue()
}

func (s StrSymbol) GetValue() string {
	return s.value
}

func (s StrSymbol) GreaterEqualThan(other interfaces.Symbol[string]) bool {
	return s.value >= other.GetValue()
}

func (s StrSymbol) GreaterThan(other interfaces.Symbol[string]) bool {
	return s.value > other.GetValue()
}

func (s StrSymbol) LesserEqualThan(other interfaces.Symbol[string]) bool {
	return s.value <= other.GetValue()
}

func (s StrSymbol) LesserThan(other interfaces.Symbol[string]) bool {
	return s.value < other.GetValue()
}

func (s StrSymbol) String() string {
	return s.value
}
