package iterator

// Iterator This is the common pattern used in Go when 'range' is not possible.
type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
