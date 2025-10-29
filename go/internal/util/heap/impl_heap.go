package heap

import (
    "interview_go/internal/util/iterator"
)

// ImplHeap is a minimal heap implementation, using a slice as the underlying data structure.
// The heap is always ordered, and the array is indexed, using the position to calculate the parent and children.
//
// Array representation and their corresponding indices:
// [0] [1] [2] [3] [4] [5] [6] [7] [8] ...
//
// Tree structure with indices:
//
//               0 (Root)
//              / \
//             /   \
//            1     2
//           / \   / \
//          3   4 5   6
//         / \
//        7   8
//
//
// Calculation rules from the code:
// Parent of node i: (i - 1) / 2
// Left child of node i: 2*i + 1
// Right child of node i: 2*i + 2
//
// Examples:
//
// Node at index 0 (Root):
//   - No parent (or calculated as -1/2 = 0, but handled by checking index > 0)
//   - Left child: 2*0 + 1 = 1
//   - Right child: 2*0 + 2 = 2
//
// Node at index 1:
//   - Parent: (1 - 1) / 2 = 0
//   - Left child: 2*1 + 1 = 3
//   - Right child: 2*1 + 2 = 4
//
// Node at index 2:
//   - Parent: (2 - 1) / 2 = 0
//   - Left child: 2*2 + 1 = 5
//   - Right child: 2*2 + 2 = 6
//
// Node at index 3:
//   - Parent: (3 - 1) / 2 = 1
//   - Left child: 2*3 + 1 = 7
//   - Right child: 2*3 + 2 = 8
type ImplHeap[T any] struct {
    data       []T
    comparator func(a, b T) int // Returns: <0 if a<b, 0 if a==b, >0 if a>b
    isMaxHeap  bool
}

// NewMinHeap creates and returns a new empty heap.
func NewMinHeap[T any](comparator func(a, b T) int) *ImplHeap[T] {
    return &ImplHeap[T]{data: make([]T, 0), comparator: comparator, isMaxHeap: false}
}

// NewMaxHeap creates and returns a new empty heap.
func NewMaxHeap[T any](comparator func(a, b T) int) *ImplHeap[T] {
    return &ImplHeap[T]{data: make([]T, 0), comparator: comparator, isMaxHeap: true}
}

// Compile-time check to ensure ImplHeap implements the Heap interface
var _ Heap[string] = (*ImplHeap[string])(nil)

// isHigherPriority return true if the element on i-th position has higher priority
// than the element on j-th position
func (h *ImplHeap[T]) isHigherPriority(i, j int) bool {
    cmp := h.comparator(h.data[i], h.data[j])
    if h.isMaxHeap {
        // MaxHeap: higher priority to the element with greater value
        return cmp > 0
    }
    // ImplHeap: higher priority to the element with a smaller value
    return cmp < 0
}

// Auxiliary functions for indexing the data array
// parent look for the parent of the given node
func parent(i int) int {
    return (i - 1) / 2
}

// leftChild give the index of the left child of the given node
func leftChild(i int) int {
    return 2*i + 1
}

// rightChild give the index of the right child of the given node
func rightChild(i int) int {
    return 2*i + 2
}

// siftUp moves the given node up the tree until it reaches the root
func (h *ImplHeap[T]) siftUp(index int) {
    if index == 0 {
        // Already at the root
        return
    }

    for index > 0 && h.isHigherPriority(index, parent(index)) {
        h.data[index], h.data[parent(index)] = h.data[parent(index)], h.data[index]
        index = parent(index)
    }
}

// siftDown moves the given node down inthe tree until it reaches a leaf.
// Needed for Heapify and ExtractMin.
//
// O(log n)
func (h *ImplHeap[T]) siftDown(index int) {
    n := len(h.data)

    // Loop until the node is a leaf
    for {
        parentIndex := index
        leftChildIndex := leftChild(index)
        rightChildIndex := rightChild(index)

        // Starts with the parent and checks if it has a child that is higher priority
        minIndex := parentIndex

        // Check if the left leaf exists and has higher priority
        if leftChildIndex < n && h.isHigherPriority(leftChildIndex, minIndex) {
            minIndex = leftChildIndex
        }

        // Check if the right leaf exists and has higher priority
        if rightChildIndex < n && h.isHigherPriority(rightChildIndex, minIndex) {
            minIndex = rightChildIndex
        }

        // Parent is already the minimum, so we can stop
        if minIndex == parentIndex {
            return
        }

        // Swap the parent with the minimum
        h.data[parentIndex], h.data[minIndex] = h.data[minIndex], h.data[parentIndex]

        // Move index and start the loop again
        index = minIndex
    }
}

func (h *ImplHeap[T]) Peek() (T, bool) {
    if len(h.data) == 0 {
        return *new(T), false
    }
    return h.data[0], true
}

func (h *ImplHeap[T]) Push(value T) {
    // Add the element to the end of the array
    h.data = append(h.data, value)

    // Sift up the new element to the correct position
    h.siftUp(len(h.data) - 1)
}

func (h *ImplHeap[T]) Pop() (T, bool) {
    // zeroValue is declared for return in case of failure (empty heap).
    var zeroValue T

    // Check for Empty Heap
    if len(h.data) == 0 {
        return zeroValue, false
    }

    size := len(h.data)

    // Store the value to be returned (the minimum/root element)
    minValue := h.data[0]

    // Substitution: Move the LAST element of the slice to the Root (index 0).
    // This temporarily breaks the Heap property.
    h.data[0] = h.data[size-1]

    // 5. Shrink the Slice (remove the last element, which has been moved to index 0).
    h.data = h.data[:size-1]

    // Restore the Heap Property (Sift-Down)
    // The new root element (at index 0) "sinks" to its correct position.
    if len(h.data) > 0 {
        h.siftDown(0)
    }

    return minValue, true
}

func (h *ImplHeap[T]) Size() int {
    return len(h.data)
}

func (h *ImplHeap[T]) IsEmpty() bool {
    return len(h.data) == 0
}

func (h *ImplHeap[T]) Clear() {
    h.data = make([]T, 0)
}

func (h *ImplHeap[T]) Heapify(elements []T) {
    // Copy data
    h.data = elements

    // Begin sifting down
    for i := len(h.data)/2 - 1; i >= 0; i-- {
        h.siftDown(i)
    }
}

// ToSlice returns a copy of the underlying data slice - similar to BFS
func (h *ImplHeap[T]) ToSlice() []T {
    return h.data
}

// Iterator returns an iterator that traverses the elements in the heap.
func (h *ImplHeap[T]) Iterator() iterator.Iterator[T] {
    return newHeapIterator(h.data)
}

// SortedIterator Create an iterator that returns the elements in sorted order. Performs a copy of the data.
func (h *ImplHeap[T]) SortedIterator() iterator.Iterator[T] {
    return newHeapSortedIterator(h)
}

// Clone Copies the original Heap and returns a new Heap with the copied data.
// Complexity: O(n) for slice copy
func (h *ImplHeap[T]) Clone() *ImplHeap[T] {
    // Copy slice
    newData := make([]T, len(h.data))
    copy(newData, h.data)

    // Create and copy a new heap
    return &ImplHeap[T]{
        data: newData, // Usa o slice original (já copiado acima),
        // ou o novo slice se preferir. Para MinHeap é mais seguro usar o original
        isMaxHeap:  h.isMaxHeap,
        comparator: h.comparator,
    }
}
