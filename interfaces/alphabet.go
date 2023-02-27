package interfaces

import "github.com/neuralchecker/go-automata/internal/iterator"

type Alphabet[T any] interface {
	Contains(symbol Symbol[T]) bool
	Length() int
	Equal(other Alphabet[T]) bool
	String() string
	GetSymbols() []Symbol[T]
	GetSymbolAt(i int) Symbol[T]
	Iterator() iterator.Iterator[Symbol[T]]
	IsComplete() bool
}
