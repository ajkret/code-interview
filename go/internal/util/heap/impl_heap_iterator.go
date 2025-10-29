package heap

import (
    "interview_go/internal/util/iterator"
)

// heapIterator[T] holds the state for sequential traversal of the underlying array.
// It iterates through the 'data' slice in linear (index) order.
type heapIterator[T any] struct {
    // Reference to the slice being iterated over
    data []T
    // Current index in the slice
    index int
}

// newHeapIterator is the internal constructor for the iterator.
func newHeapIterator[T any](data []T) iterator.Iterator[T] {
    return &heapIterator[T]{
        // The iterator takes a COPY of the slice header.
        // NOTE: This copy still points to the same underlying array data.
        data:  data,
        index: 0,
    }
}

// HasNext implements the Iterator[T].HasNext method.
// Checks if the current index is within the bounds of the data slice.
func (it *heapIterator[T]) HasNext() bool {
    return it.index < len(it.data)
}

// Next implements the Iterator[T].Next method.
// Returns the element at the current index and advances the index.
func (it *heapIterator[T]) Next() T {
    // Check for panic condition (calling Next() after HasNext() is false)
    if it.index >= len(it.data) {
        panic("iterator: cannot call Next() after iteration finished")
    }

    // 1. Get the current value
    value := it.data[it.index]

    // 2. Advance the index
    it.index++

    // 3. Return the value (copy)
    return value
}

// -- Sorted Iterator --

// heapSortedIterator[T] holds the state for sequential traversal of the underlying array.
// It iterates through the 'data' slice in linear (index) order.
type heapSortedIterator[T any] struct {
    // Reference to the slice being iterated over
    heap Heap[T]
}

// newHeapSortedIterator create a sorted iterator
func newHeapSortedIterator[T any](heap Heap[T]) iterator.Iterator[T] {
    return &heapSortedIterator[T]{
        heap: heap.Clone(),
    }
}

func (h heapSortedIterator[T]) Next() T {
    var value, ok = h.heap.Pop()
    if !ok {
        panic("iterator: cannot call Next() after iteration finished")
    }
    return value
}

func (h heapSortedIterator[T]) HasNext() bool {
    return !h.heap.IsEmpty()
}
