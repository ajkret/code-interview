package collection

import "interview_go/internal/util/iterator"

// Collection is the generic interface that defines the contract for a collection
// with focus on read-only methods (no mutators)
// T is a type parameter representing the type of elements in this collection.
type Collection[T any] interface {
	// Get retrieves the element at the specified position in the list.
	Get(index int) (T, bool)

	// GetFromEnd retrieves the element at the specified position, counting backwards from the end, if possible.
	GetFromEnd(index int) (T, bool)

	// Size returns the number of elements in the list.
	Size() int

	// IsEmpty checks if the list contains no elements.
	IsEmpty() bool

	// End retrieves the element at the end of the list without removing it.
	End() (T, bool)

	// Start retrieves the element at the start of the list without removing it.
	Start() (T, bool)

	// Iterator returns an iterator over elements of type T.
	Iterator() iterator.Iterator[T]
}
