package sequences

import (
	"strings"

	"github.com/neuralchecker/go-automata/interfaces"
	"github.com/neuralchecker/go-automata/internal/iterator"
)

type Sequence[T any] struct {
	symbols []interfaces.Symbol[T]
}

func New[T any](symbols ...interfaces.Symbol[T]) interfaces.Sequence[T] {
	return Sequence[T]{symbols: symbols}
}

func FromSlice[T any](symbols []interfaces.Symbol[T]) interfaces.Sequence[T] {
	return New(symbols...)
}

func (s Sequence[T]) Append(symbols ...interfaces.Symbol[T]) interfaces.Sequence[T] {
	newSymbols := make([]interfaces.Symbol[T], len(s.symbols)+len(symbols))
	copy(newSymbols, s.symbols)
	copy(newSymbols[len(s.symbols):], symbols)
	return Sequence[T]{symbols: newSymbols}
}

func (s Sequence[T]) AsSlice() []interfaces.Symbol[T] {
	return s.symbols
}

func (s Sequence[T]) Equal(other interfaces.Sequence[T]) bool {
	sSlice := s.AsSlice()
	oSlice := other.AsSlice()
	if len(sSlice) != len(oSlice) {
		return false
	}
	for i := range sSlice {
		if !sSlice[i].Equal(oSlice[i]) {
			return false
		}
	}
	return true
}

func (s Sequence[T]) GetPreffixes() []interfaces.Sequence[T] {
	prefixes := make([]interfaces.Sequence[T], 0, s.Length()+1)
	for i := 0; i <= s.Length(); i++ {
		prefixes = append(prefixes, s.getSubsequence(0, i))
	}
	return prefixes
}

func (s Sequence[T]) GetSuffixes() []interfaces.Sequence[T] {
	suffixes := make([]interfaces.Sequence[T], 0, s.Length()+1)
	for i := 0; i <= s.Length(); i++ {
		suffixes = append(suffixes, s.getSubsequence(i, s.Length()))
	}
	return suffixes
}

func (s Sequence[T]) IsEmpty() bool {
	return len(s.symbols) == 0
}

func (s Sequence[T]) Iterator() iterator.Iterator[interfaces.Symbol[T]] {
	return iterator.NewSliceIterator(s.symbols)
}

func (s Sequence[T]) Length() int {
	return len(s.symbols)
}

func (s Sequence[T]) Prepend(symbols ...interfaces.Symbol[T]) interfaces.Sequence[T] {
	newSymbols := make([]interfaces.Symbol[T], len(s.symbols)+len(symbols))
	copy(newSymbols, s.symbols)
	copy(newSymbols[len(s.symbols):], symbols)
	return Sequence[T]{symbols: newSymbols}
}

func (s Sequence[T]) String() string {
	strBuilder := strings.Builder{}
	for _, symbol := range s.symbols {
		strBuilder.WriteString(symbol.String())
	}
	if str := strBuilder.String(); str == "" {
		return "ε"
	} else {
		return str
	}
}

func (s Sequence[T]) GetSymbolAt(index int) interfaces.Symbol[T] {
	return s.symbols[index]
}

func (s Sequence[T]) getSubsequence(start, end int) interfaces.Sequence[T] {
	return Sequence[T]{symbols: s.symbols[start:end]}
}

func (s Sequence[T]) Hash() int {
	hash := 0
	mult := 1
	for _, symbol := range s.symbols {
		hash += symbol.Hash() * mult
		mult *= 31
	}
	return hash
}
