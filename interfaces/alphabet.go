package interfaces

import "github.com/neuralchecker/go-automata/internal/iterator"

type Alphabet[T any] interface {
	Contains(symbol Symbol[T]) bool
	Length() int
	Equals(other Alphabet[T]) bool
	String() string
	GetSymbols() []Symbol[T]
	Iterator() iterator.Iterator[Symbol[T]]
}
