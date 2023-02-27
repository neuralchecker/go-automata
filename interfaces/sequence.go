package interfaces

import (
	"github.com/neuralchecker/go-automata/internal/iterator"
)

type Sequence[T any] interface {
	// GetPrefixes returns the prefixes of the sequence.
	GetPreffixes() []Sequence[T]
	// GetSuffixes returns the suffixes of the sequence.
	GetSuffixes() []Sequence[T]
	// Append appends the specified symbol(s) to a copy of the sequence and returns it.
	Append(symbols ...Symbol[T]) Sequence[T]
	// Prepend prepends the specified symbol(s) to a copy of the sequence and returns it.
	Prepend(symbols ...Symbol[T]) Sequence[T]
	// Length returns the length of the sequence.
	Length() int
	// IsEmpty returns true if the sequence is empty, false otherwise.
	IsEmpty() bool
	// GetSymbolAt returns the symbol at the specified index.
	GetSymbolAt(index int) Symbol[T]
	//AsSlice returns the sequence as a slice.
	AsSlice() []Symbol[T]

	String() string
	Equal(other Sequence[T]) bool
	Iterator() iterator.Iterator[Symbol[T]]
	Hash() int
}
