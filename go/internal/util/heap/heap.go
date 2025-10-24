package heap

import "interview_go/internal/util/iterator"

// Heap is a generic interface representing a heap data structure.
// A heap is a specialized tree-based structure that satisfies the heap property:
// - Min-Heap: Parent node is always less than or equal to its children
// - Max-Heap: Parent node is always greater than or equal to its children
//
// Example Min-Heap Structure:
//
//             10
//            /  \
//           15   20
//          / \   /
//         17 25 30
//
// In this min-heap:
// - The root (10) is the smallest element
// - Every parent is smaller than its children
// - The structure is complete (filled left to right)
//
// Heaps are typically used for:
// - Priority queues
// - Heap sort algorithm
// - Finding kth smallest/largest elements
// - Scheduling algorithms
//
type Heap[T any] interface {
    // Peek returns the root element (min or max depending on heap type) without removing it.
    // Returns the element and true if the heap is not empty, otherwise returns zero value and false.
    // This is the primary "read" operation for heaps.
    //
    // For a min-heap, this returns the smallest element.
    // For a max-heap, this returns the largest element.
    //
    // Time complexity: O(1)
    Peek() (T, bool)

    // Push inserts a new value into the heap while maintaining the heap property.
    // The element is added at the end and then "bubbled up" to its correct position.
    //
    // Example: Pushing 12 into this min-heap:
    //
    // Before:
    //        10
    //       /  \
    //      15   20
    //
    // After:
    //        10
    //       /  \
    //      12   20
    //     /
    //    15
    //
    // Time complexity: O(log n)
    Push(value T)

    // Pop removes and returns the root element (min or max depending on heap type).
    // The last element is moved to the root and then "bubbled down" to restore heap property.
    // Returns the removed element and true if successful, otherwise returns zero value and false.
    //
    // Example: Popping from this min-heap:
    //
    // Before:
    //        10
    //       /  \
    //      15   20
    //     /
    //    17
    //
    // After (returns 10):
    //        15
    //       /  \
    //      17   20
    //
    // Time complexity: O(log n)
    Pop() (T, bool)

    // Size returns the number of elements currently in the heap.
    //
    // Time complexity: O(1)
    Size() int

    // IsEmpty returns true if the heap contains no elements.
    //
    // Time complexity: O(1)
    IsEmpty() bool

    // Clear removes all elements from the heap.
    // After clearing, the heap will be empty, and Size() will return 0.
    //
    // Time complexity: O(1)
    Clear()

    // Heapify converts a slice of elements into a valid heap structure.
    // This is more efficient than inserting elements one by one.
    //
    // Example: Heapify [20, 15, 10, 17] into a min-heap:
    //
    // Input array: [20, 15, 10, 17]
    //
    // Result (min-heap):
    //       10
    //      /  \
    //     17   15
    //    /
    //   20
    //
    // Time complexity: O(n) - more efficient than n × Push which would be O(n log n)
    Heapify(elements []T)

    // ToSlice returns a slice containing all elements in the heap.
    // The order may be the internal array representation (not necessarily sorted).
    // For a sorted output, repeatedly call Pop() instead.
    //
    // Time complexity: O(n)
    ToSlice() []T

    // Iterator returns an iterator for traversing heap elements.
    // Note: The iteration order is typically level-order (breadth-first),
    // which reflects the underlying array structure, NOT sorted order.
    //
    // Example level-order traversal for this min-heap:
    //       10
    //      /  \
    //     15   20
    //    /  \
    //   17  25
    //
    // Iterator order: 10, 15, 20, 17, 25
    //
    // For sorted order, use Pop() repeatedly instead.
    Iterator() iterator.Iterator[T]
}
