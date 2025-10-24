package list

// List is the generic interface that defines the contract for a linear collection (list).
// T is a type parameter representing the type of elements in this list.
// Only the mutator methods
type List[T any] interface {
	// Add appends an element to the end of the list.
	Add(value T)

	// Clear removes all elements from the list.
	Clear()

	// Remove deletes the element at the specified position in the list.
	Remove(index int) bool

	// RemoveFromEnd removes and returns the element at the end of the list (useful for Stack Pop).
	RemoveFromEnd() (T, bool)
}
