package iterator

import "errors"

func NewSetIterator[S map[T]struct{}, T comparable](set S) Iterator[T] {
	setSlice := make([]T, len(set))
	for key := range set {
		setSlice = append(setSlice, key)
	}
	return &sliceIterator[T]{
		slice: setSlice,
		index: 0,
	}
}

func NewSetAssertIterator[S comparable, T any](set map[S]struct{}) (Iterator[T], error) {
	setSlice := make([]T, len(set))
	i := 0
	for key := range set {
		// we use the any() function to convert the key to an interface{}, in order to be able to use the type assertion.
		if keyT, ok := (any(key)).(T); ok {
			setSlice[i] = keyT
			i++
		} else {
			return nil, errors.New("cannot assert key to type T")
		}
	}
	return &sliceIterator[T]{
		slice: setSlice,
		index: 0,
	}, nil
}
