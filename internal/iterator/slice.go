package iterator

type sliceIterator[T any] struct {
	slice []T
	index int
}

func NewSliceIterator[T any](slice []T) *sliceIterator[T] {
	return &sliceIterator[T]{
		slice: slice,
		index: 0,
	}
}

func (iterator *sliceIterator[T]) HasNext() bool {
	return iterator.index < len(iterator.slice)
}

func (iterator *sliceIterator[T]) Next() T {
	iterator.index++
	return iterator.slice[iterator.index-1]
}
