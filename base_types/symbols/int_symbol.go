package symbols

import (
	"fmt"

	"github.com/neuralchecker/go-automata/interfaces"
)

type IntSymbol struct {
	value int
}

func (s IntSymbol) AddToValue(value int) interfaces.Symbol[int] {
	return IntSymbol{value: s.value + value}
}

func (s IntSymbol) Equal(other interfaces.Symbol[int]) bool {
	return s.value == other.GetValue()
}

func (s IntSymbol) GetValue() int {
	return s.value
}

func (s IntSymbol) GreaterEqualThan(other interfaces.Symbol[int]) bool {
	return s.value >= other.GetValue()
}

func (s IntSymbol) GreaterThan(other interfaces.Symbol[int]) bool {
	return s.value > other.GetValue()
}

func (s IntSymbol) LesserEqualThan(other interfaces.Symbol[int]) bool {
	return s.value <= other.GetValue()
}

func (s IntSymbol) LesserThan(other interfaces.Symbol[int]) bool {
	return s.value < other.GetValue()
}

func (s IntSymbol) String() string {
	return fmt.Sprint(s.value)
}

func (s IntSymbol) Hash() int {
	return s.value
}
