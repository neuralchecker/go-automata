package interfaces

type Guard[T any] interface {
	// IsSatisfied returns true if the guard is satisfied by the given symbol.
	IsSatisfied(symbol Symbol[T]) bool
	// String returns a string representation of the guard.
	String() string
}
