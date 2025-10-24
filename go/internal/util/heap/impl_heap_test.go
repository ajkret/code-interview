package heap

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

// TestNewMinHeap tests the creation of a new min-heap
func TestNewMinHeap(t *testing.T) {
    heap := NewMinHeap[int]()
    assert.NotNil(t, heap, "NewMinHeap should return a non-nil heap")
    assert.False(t, heap.isMaxHeap, "NewMinHeap should create a min-heap")
    assert.Empty(t, heap.data, "NewMinHeap should initialize an empty data slice")
    assert.True(t, heap.IsEmpty(), "NewMinHeap should create an empty heap")
    assert.Equal(t, 0, heap.Size(), "NewMinHeap should have size 0")
}

// TestNewMaxHeap tests the creation of a new max-heap
func TestNewMaxHeap(t *testing.T) {
    heap := NewMaxHeap[int]()
    assert.NotNil(t, heap, "NewMaxHeap should return a non-nil heap")
    assert.True(t, heap.isMaxHeap, "NewMaxHeap should create a max-heap")
    assert.Empty(t, heap.data, "NewMaxHeap should initialize an empty data slice")
    assert.True(t, heap.IsEmpty(), "NewMaxHeap should create an empty heap")
    assert.Equal(t, 0, heap.Size(), "NewMaxHeap should have size 0")
}

// TestMinHeap_PushAndPeek tests pushing elements and peeking at the min element
func TestMinHeap_PushAndPeek(t *testing.T) {
    heap := NewMinHeap[int]()

    // Test peek on empty heap
    val, ok := heap.Peek()
    assert.False(t, ok, "Peek on empty heap should return false")
    assert.Equal(t, 0, val, "Peek on empty heap should return zero value")

    // Push single element
    heap.Push(10)
    val, ok = heap.Peek()
    assert.True(t, ok, "Peek should return true after pushing element")
    assert.Equal(t, 10, val, "Peek should return the only element")
    assert.Equal(t, 1, heap.Size(), "Size should be 1 after pushing one element")

    // Push multiple elements
    heap.Push(5)
    val, ok = heap.Peek()
    assert.True(t, ok)
    assert.Equal(t, 5, val, "Peek should return the minimum element (5)")

    heap.Push(15)
    val, ok = heap.Peek()
    assert.True(t, ok)
    assert.Equal(t, 5, val, "Peek should still return 5 as minimum")

    heap.Push(3)
    val, ok = heap.Peek()
    assert.True(t, ok)
    assert.Equal(t, 3, val, "Peek should return the new minimum (3)")
    assert.Equal(t, 4, heap.Size(), "Size should be 4")
}

// TestMaxHeap_PushAndPeek tests pushing elements and peeking at the max element
func TestMaxHeap_PushAndPeek(t *testing.T) {
    heap := NewMaxHeap[int]()

    heap.Push(10)
    val, ok := heap.Peek()
    assert.True(t, ok)
    assert.Equal(t, 10, val, "Peek should return the only element")

    heap.Push(5)
    val, ok = heap.Peek()
    assert.True(t, ok)
    assert.Equal(t, 10, val, "Peek should return the maximum element (10)")

    heap.Push(15)
    val, ok = heap.Peek()
    assert.True(t, ok)
    assert.Equal(t, 15, val, "Peek should return the new maximum (15)")

    heap.Push(20)
    val, ok = heap.Peek()
    assert.True(t, ok)
    assert.Equal(t, 20, val, "Peek should return the new maximum (20)")
    assert.Equal(t, 4, heap.Size(), "Size should be 4")
}

// TestMinHeap_Pop tests popping elements from min-heap
func TestMinHeap_Pop(t *testing.T) {
    heap := NewMinHeap[int]()

    // Test pop on empty heap
    val, ok := heap.Pop()
    assert.False(t, ok, "Pop on empty heap should return false")
    assert.Equal(t, 0, val, "Pop on empty heap should return zero value")

    // Push and pop single element
    heap.Push(10)
    val, ok = heap.Pop()
    assert.True(t, ok, "Pop should return true")
    assert.Equal(t, 10, val, "Pop should return the element")
    assert.True(t, heap.IsEmpty(), "Heap should be empty after popping")

    // Push multiple elements and pop in order
    elements := []int{20, 10, 30, 5, 15}
    for _, elem := range elements {
        heap.Push(elem)
    }

    expected := []int{5, 10, 15, 20, 30}
    for _, exp := range expected {
        val, ok = heap.Pop()
        assert.True(t, ok, "Pop should succeed")
        assert.Equal(t, exp, val, "Pop should return elements in ascending order")
    }

    assert.True(t, heap.IsEmpty(), "Heap should be empty after popping all elements")
}

// TestMaxHeap_Pop tests popping elements from max-heap
func TestMaxHeap_Pop(t *testing.T) {
    heap := NewMaxHeap[int]()

    // Push multiple elements and pop in order
    elements := []int{20, 10, 30, 5, 15}
    for _, elem := range elements {
        heap.Push(elem)
    }

    expected := []int{30, 20, 15, 10, 5}
    for _, exp := range expected {
        val, ok := heap.Pop()
        assert.True(t, ok, "Pop should succeed")
        assert.Equal(t, exp, val, "Pop should return elements in descending order")
    }

    assert.True(t, heap.IsEmpty(), "Heap should be empty after popping all elements")
}

// TestMinHeap_Heapify tests converting a slice into a min-heap
func TestMinHeap_Heapify(t *testing.T) {
    heap := NewMinHeap[int]()

    elements := []int{20, 15, 10, 17, 25, 30, 12}
    heap.Heapify(elements)

    assert.Equal(t, len(elements), heap.Size(), "Size should match number of heapified elements")

    // Pop all elements and verify they come out in sorted order
    var result []int
    for !heap.IsEmpty() {
        val, ok := heap.Pop()
        assert.True(t, ok)
        result = append(result, val)
    }

    expected := []int{10, 12, 15, 17, 20, 25, 30}
    assert.Equal(t, expected, result, "Heapified elements should pop in sorted order")
}

// TestMaxHeap_Heapify tests converting a slice into a max-heap
func TestMaxHeap_Heapify(t *testing.T) {
    heap := NewMaxHeap[int]()

    elements := []int{20, 15, 10, 17, 25, 30, 12}
    heap.Heapify(elements)

    assert.Equal(t, len(elements), heap.Size(), "Size should match number of heapified elements")

    // Pop all elements and verify they come out in reverse sorted order
    var result []int
    for !heap.IsEmpty() {
        val, ok := heap.Pop()
        assert.True(t, ok)
        result = append(result, val)
    }

    expected := []int{30, 25, 20, 17, 15, 12, 10}
    assert.Equal(t, expected, result, "Heapified elements should pop in reverse sorted order")
}

// TestHeap_Clear tests clearing the heap
func TestHeap_Clear(t *testing.T) {
    heap := NewMinHeap[int]()

    elements := []int{10, 20, 30, 5, 15}
    for _, elem := range elements {
        heap.Push(elem)
    }

    assert.Equal(t, 5, heap.Size(), "Size should be 5 before clearing")
    assert.False(t, heap.IsEmpty(), "Heap should not be empty before clearing")

    heap.Clear()

    assert.Equal(t, 0, heap.Size(), "Size should be 0 after clearing")
    assert.True(t, heap.IsEmpty(), "Heap should be empty after clearing")

    val, ok := heap.Peek()
    assert.False(t, ok, "Peek should return false on cleared heap")
    assert.Equal(t, 0, val, "Peek should return zero value on cleared heap")
}

// TestHeap_ToSlice tests converting heap to slice
func TestHeap_ToSlice(t *testing.T) {
    heap := NewMinHeap[int]()

    // Test empty heap
    slice := heap.ToSlice()
    assert.Empty(t, slice, "ToSlice on empty heap should return empty slice")

    // Add elements
    elements := []int{10, 20, 30, 5, 15}
    for _, elem := range elements {
        heap.Push(elem)
    }

    slice = heap.ToSlice()
    assert.Equal(t, 5, len(slice), "ToSlice should return all elements")

    // Verify that the slice contains all elements (order may vary)
    for _, elem := range elements {
        assert.Contains(t, slice, elem, "ToSlice should contain all pushed elements")
    }
}

// TestHeap_Iterator tests iterating over heap elements
func TestHeap_Iterator(t *testing.T) {
    heap := NewMinHeap[int]()

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

// TestMinHeap_WithStrings tests min-heap with string type
func TestMinHeap_WithStrings(t *testing.T) {
    heap := NewMinHeap[string]()

    words := []string{"dog", "cat", "bird", "elephant", "ant"}
    for _, word := range words {
        heap.Push(word)
    }

    expected := []string{"ant", "bird", "cat", "dog", "elephant"}
    for _, exp := range expected {
        val, ok := heap.Pop()
        assert.True(t, ok)
        assert.Equal(t, exp, val, "Strings should pop in lexicographic order")
    }
}

// TestMaxHeap_WithStrings tests max-heap with string type
func TestMaxHeap_WithStrings(t *testing.T) {
    heap := NewMaxHeap[string]()

    words := []string{"dog", "cat", "bird", "elephant", "ant"}
    for _, word := range words {
        heap.Push(word)
    }

    expected := []string{"elephant", "dog", "cat", "bird", "ant"}
    for _, exp := range expected {
        val, ok := heap.Pop()
        assert.True(t, ok)
        assert.Equal(t, exp, val, "Strings should pop in reverse lexicographic order")
    }
}

// TestHeap_LargeDataSet tests heap with a larger dataset
func TestHeap_LargeDataSet(t *testing.T) {
    heap := NewMinHeap[int]()

    // Push 100 elements
    for i := 100; i > 0; i-- {
        heap.Push(i)
    }

    assert.Equal(t, 100, heap.Size(), "Size should be 100")

    // Pop all elements and verify they come out sorted
    for i := 1; i <= 100; i++ {
        val, ok := heap.Pop()
        assert.True(t, ok)
        assert.Equal(t, i, val, "Elements should pop in sorted order")
    }

    assert.True(t, heap.IsEmpty(), "Heap should be empty after popping all elements")
}

// TestHeap_DuplicateElements tests heap with duplicate elements
func TestHeap_DuplicateElements(t *testing.T) {
    heap := NewMinHeap[int]()

    elements := []int{5, 10, 5, 20, 10, 5}
    for _, elem := range elements {
        heap.Push(elem)
    }

    expected := []int{5, 5, 5, 10, 10, 20}
    for _, exp := range expected {
        val, ok := heap.Pop()
        assert.True(t, ok)
        assert.Equal(t, exp, val, "Duplicates should be handled correctly")
    }
}
