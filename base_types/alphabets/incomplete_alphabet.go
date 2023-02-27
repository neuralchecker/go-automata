package alphabets

import (
	"strings"

	"github.com/neuralchecker/go-automata/interfaces"
)

type incompleteAlphabet[T any] struct {
	alphabet[T]
}

func NewIncomplete[T any](symbols []interfaces.Symbol[T]) interfaces.Alphabet[T] {
	//struct{}{} is an empty struct, so it takes no space in memory.
	symbolsSet := make(map[interfaces.Symbol[T]]struct{}, len(symbols))
	for _, symbol := range symbols {
		symbolsSet[symbol] = struct{}{}
	}
	return incompleteAlphabet[T]{alphabet[T]{
		symbols: symbolsSet,
	}}
}

func (a incompleteAlphabet[T]) Equal(other interfaces.Alphabet[T]) bool {
	if a.Length() > other.Length() {
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

func (a incompleteAlphabet[T]) String() string {
	count := 0
	strBuilder := strings.Builder{}
	strBuilder.WriteString("IncompleteAlphabet{")
	for symbol := range a.symbols {
		strBuilder.WriteString(symbol.String())
		if count < len(a.symbols)-1 {
			strBuilder.WriteString(", ")
		}
	}
	strBuilder.WriteString("}")
	return strBuilder.String()
}

func (a incompleteAlphabet[T]) IsComplete() bool {
	return false
}
