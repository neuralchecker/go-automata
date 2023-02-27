package symbols

import (
	"github.com/neuralchecker/go-automata/interfaces"
)

type StrSymbol struct {
	value string
}

func (s StrSymbol) AddToValue(value string) interfaces.Symbol[string] {
	return StrSymbol{value: s.value + value}
}

func (s StrSymbol) Equal(other interfaces.Symbol[string]) bool {
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

func (s StrSymbol) Hash() int {
	hash := 0
	mult := 1
	for _, c := range s.value {
		hash += int(c) * mult
		mult *= 31
	}
	return hash
}
