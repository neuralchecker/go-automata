package alphabets

import (
	"strings"

	"github.com/neuralchecker/go-automata/interfaces"
	"github.com/neuralchecker/go-automata/internal/iterator"
)

type alphabet[T any] struct {
	symbols      map[interfaces.Symbol[T]]struct{}
	symbolsSlice []interfaces.Symbol[T]
}

func New[T any](symbols []interfaces.Symbol[T]) interfaces.Alphabet[T] {
	//struct{}{} is an empty struct, so it takes no space in memory.
	symbolsSet := make(map[interfaces.Symbol[T]]struct{}, len(symbols))
	for _, symbol := range symbols {
		symbolsSet[symbol] = struct{}{}
	}
	return alphabet[T]{
		symbols:      symbolsSet,
		symbolsSlice: symbols,
	}
}

func (a alphabet[T]) IsComplete() bool {
	return true
}

// Contains implements interfaces.Alphabet
func (a alphabet[T]) Contains(symbol interfaces.Symbol[T]) bool {
	_, ok := a.symbols[symbol]
	return ok
}

// Equals implements interfaces.Alphabet
func (a alphabet[T]) Equals(other interfaces.Alphabet[T]) bool {
	if a.Length() != other.Length() {
		return false
	}
	it := a.Iterator()
	for it.HasNext() {
		if !other.Contains(it.Next()) {
			return false
		}
	}
	return true
}

func (a alphabet[T]) GetSymbolAt(i int) interfaces.Symbol[T] {
	return a.symbolsSlice[i]
}

// GetSymbols implements interfaces.Alphabet
func (a alphabet[T]) GetSymbols() []interfaces.Symbol[T] {
	return a.symbolsSlice
}

// Iterator implements interfaces.Alphabet
func (a alphabet[T]) Iterator() iterator.Iterator[interfaces.Symbol[T]] {
	it := iterator.NewSliceIterator(a.symbolsSlice)

	return it
}

// Length implements interfaces.Alphabet
func (a alphabet[T]) Length() int {
	return len(a.symbolsSlice)
}

// String implements interfaces.Alphabet
func (a alphabet[T]) String() string {
	count := 0
	strBuilder := strings.Builder{}
	strBuilder.WriteString("Alphabet{")
	for symbol := range a.symbols {
		strBuilder.WriteString(symbol.String())
		if count < len(a.symbols)-1 {
			strBuilder.WriteString(", ")
		}
	}
	strBuilder.WriteString("}")
	return strBuilder.String()
}
