package heap

import (
    "strings"
    "testing"

    "github.com/stretchr/testify/assert"
)

// TestHeap_Iterator tests iterating over heap elements
func TestHeap_Iterator(t *testing.T) {
    heap := NewMinHeap[int](func(a, b int) int {
        return a - b
    })

    elements := []int{10, 20, 30, 5, 15}
    for _, elem := range elements {
        heap.Push(elem)
    }

    iter := heap.Iterator()
    assert.NotNil(t, iter, "Iterator should not be nil")

    var result []int
    for iter.HasNext() {
        val := iter.Next()
        result = append(result, val)
    }

    assert.Equal(t, 5, len(result), "Iterator should iterate over all elements")

    // Verify all elements are present (order may vary)
    for _, elem := range elements {
        assert.Contains(t, result, elem, "Iterator should return all elements")
    }

    // Test that Next returns false after iteration is complete
    assert.Panics(t, func() {
        iter.Next()
    }, "Next should return zero value after iteration completes")
}

// TestHeap_SortedIterator tests the SortedIterator for min-heap
func TestHeap_SortedIterator(t *testing.T) {
    heap := NewMinHeap[int](func(a, b int) int {
        return a - b
    })

    // Test SortedIterator on empty heap
    iter := heap.SortedIterator()
    assert.NotNil(t, iter, "SortedIterator should not be nil even for empty heap")
    assert.False(t, iter.HasNext(), "SortedIterator on empty heap should have no elements")

    // Add elements to heap
    elements := []int{20, 10, 30, 5, 15, 25, 35, 40}
    for _, elem := range elements {
        heap.Push(elem)
    }

    // Get sorted iterator
    iter = heap.SortedIterator()
    assert.NotNil(t, iter, "SortedIterator should not be nil")

    // Collect all elements from the iterator
    var result []int
    for iter.HasNext() {
        val := iter.Next()
        result = append(result, val)
    }

    // Verify elements come out in sorted order (ascending for min-heap)
    expected := []int{5, 10, 15, 20, 25, 30, 35, 40}
    assert.Equal(t, expected, result, "SortedIterator should return elements in ascending order")

    // Verify the original heap is not modified
    assert.Equal(t, len(elements), heap.Size(), "SortedIterator should not modify original heap size")
    peekVal, peekOk := heap.Peek()
    assert.True(t, peekOk, "Original heap should still be valid")
    assert.Equal(t, 5, peekVal, "Original heap root should be unchanged")

    // Test that Next panics after iteration is complete
    assert.Panics(t, func() {
        iter.Next()
    }, "Next should panic after iteration completes")
}

// TestMaxHeap_SortedIterator tests the SortedIterator for max-heap
func TestMaxHeap_SortedIterator(t *testing.T) {
    heap := NewMaxHeap[int](func(a, b int) int {
        return a - b
    })

    elements := []int{20, 10, 30, 5, 15, 25}
    for _, elem := range elements {
        heap.Push(elem)
    }

    // Get sorted iterator
    iter := heap.SortedIterator()
    assert.NotNil(t, iter, "SortedIterator should not be nil")

    // Collect all elements from the iterator
    var result []int
    for iter.HasNext() {
        val := iter.Next()
        result = append(result, val)
    }

    // Verify elements come out in descending order (max-heap)
    expected := []int{30, 25, 20, 15, 10, 5}
    assert.Equal(t, expected, result, "SortedIterator should return elements in descending order for max-heap")

    // Verify the original heap is not modified
    assert.Equal(t, len(elements), heap.Size(), "SortedIterator should not modify original heap size")
}

// TestHeap_SortedIteratorWithStrings tests SortedIterator with string type
func TestHeap_SortedIteratorWithStrings(t *testing.T) {
    heap := NewMinHeap[string](func(a, b string) int {
        return strings.Compare(a, b)
    })

    words := []string{"dog", "cat", "bird", "elephant", "ant", "zebra"}
    for _, word := range words {
        heap.Push(word)
    }

    iter := heap.SortedIterator()
    var result []string
    for iter.HasNext() {
        result = append(result, iter.Next())
    }

    expected := []string{"ant", "bird", "cat", "dog", "elephant", "zebra"}
    assert.Equal(t, expected, result, "SortedIterator should return strings in alphabetical order")
}

// TestHeap_SortedIteratorSingleElement tests SortedIterator with single element
func TestHeap_SortedIteratorSingleElement(t *testing.T) {
    heap := NewMinHeap[int](func(a, b int) int {
        return a - b
    })

    heap.Push(42)

    iter := heap.SortedIterator()
    assert.True(t, iter.HasNext(), "SortedIterator should have one element")

    val := iter.Next()
    assert.Equal(t, 42, val, "SortedIterator should return the single element")

    assert.False(t, iter.HasNext(), "SortedIterator should have no more elements")
}

// TestHeap_SortedIteratorDoesNotModifyOriginal tests that SortedIterator creates a copy
func TestHeap_SortedIteratorDoesNotModifyOriginal(t *testing.T) {
    heap := NewMinHeap[int](func(a, b int) int {
        return a - b
    })

    elements := []int{10, 20, 30, 5, 15}
    for _, elem := range elements {
        heap.Push(elem)
    }

    // Get the original state
    originalSize := heap.Size()
    originalPeek, _ := heap.Peek()

    // Create sorted iterator and consume all elements
    iter := heap.SortedIterator()
    for iter.HasNext() {
        iter.Next()
    }

    // Verify original heap is unchanged
    assert.Equal(t, originalSize, heap.Size(), "Original heap size should be unchanged")
    currentPeek, _ := heap.Peek()
    assert.Equal(t, originalPeek, currentPeek, "Original heap root should be unchanged")

    // Pop all elements from original heap to verify all are still there
    var poppedCount int
    for !heap.IsEmpty() {
        heap.Pop()
        poppedCount++
    }
    assert.Equal(t, originalSize, poppedCount, "Original heap should still contain all elements")
}
