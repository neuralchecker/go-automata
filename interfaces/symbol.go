package interfaces

type Symbol[T any] interface {
	// GetValue returns the value of the symbol.
	GetValue() T
	// AddToValue adds the given value to the symbol's value.
	AddToValue(value T) Symbol[T]
	String() string
	Equals(other Symbol[T]) bool
	LesserThan(other Symbol[T]) bool
	GreaterThan(other Symbol[T]) bool
	LesserEqualThan(other Symbol[T]) bool
	GreaterEqualThan(other Symbol[T]) bool
}
