package sequence

import (
	"github.com/neuralchecker/go-automata/interfaces"
	"github.com/neuralchecker/go-automata/internal/iterator"
)

type Sequence[T any] struct {
	symbols []interfaces.Symbol[T]
}

var _ interfaces.Sequence[rune] = Sequence[rune]{}

func (s Sequence[T]) Append(symbols ...interfaces.Symbol[T]) interfaces.Sequence[T] {
	newSymbols := make([]interfaces.Symbol[T], len(s.symbols)+len(symbols))
	copy(newSymbols, s.symbols)
	copy(newSymbols[len(s.symbols):], symbols)
	return Sequence[T]{symbols: newSymbols}
}

func (s Sequence[T]) AsSlice() []interfaces.Symbol[T] {
	return s.symbols
}

func (s Sequence[T]) Equals(other interfaces.Sequence[T]) bool {
	sSlice := s.AsSlice()
	oSlice := other.AsSlice()
	if len(sSlice) != len(oSlice) {
		return false
	}
	for i := range sSlice {
		if !sSlice[i].Equals(oSlice[i]) {
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
	panic("unimplemented")
}

func (s Sequence[T]) Iterator() iterator.Iterator[interfaces.Symbol[T]] {
	panic("unimplemented")
}

func (s Sequence[T]) Length() int {
	panic("unimplemented")
}

func (s Sequence[T]) Prepend(symbols ...interfaces.Symbol[T]) interfaces.Sequence[T] {
	panic("unimplemented")
}

func (s Sequence[T]) String() string {
	panic("unimplemented")
}

func (s Sequence[T]) GetSymbolAt(index int) interfaces.Symbol[T] {
	panic("unimplemented")
}

func (s Sequence[T]) getSubsequence(start, end int) interfaces.Sequence[T] {
	return Sequence[T]{symbols: s.symbols[start:end]}
}
