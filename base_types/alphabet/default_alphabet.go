package alphabet

import (
	"strings"

	"github.com/neuralchecker/go-automata/interfaces"
	"github.com/neuralchecker/go-automata/internal/iterator"
)

type alphabet[T any] struct {
	symbols map[interfaces.Symbol[T]]struct{}
}

func New[T any](symbols []interfaces.Symbol[T]) interfaces.Alphabet[T] {
	//struct{}{} is an empty struct, so it takes no space in memory.
	symbolsSet := make(map[interfaces.Symbol[T]]struct{}, len(symbols))
	for _, symbol := range symbols {
		symbolsSet[symbol] = struct{}{}
	}
	return alphabet[T]{
		symbols: symbolsSet,
	}
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

// GetSymbols implements interfaces.Alphabet
func (a alphabet[T]) GetSymbols() []interfaces.Symbol[T] {
	symbols := make([]interfaces.Symbol[T], 0, len(a.symbols))
	for symbol := range a.symbols {
		// symbol.(S) is a type assertion, which is a way to convert an interface to a concrete type.
		// It panics if the type assertion fails.
		// The type assertion is safe here because we know that the symbols in the alphabet are of type S.
		symbols = append(symbols, symbol)
	}
	return symbols
}

// Iterator implements interfaces.Alphabet
func (a alphabet[T]) Iterator() iterator.Iterator[interfaces.Symbol[T]] {
	it := iterator.NewSetIterator(a.symbols)

	return it
}

// Length implements interfaces.Alphabet
func (a alphabet[T]) Length() int {
	return len(a.symbols)
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
