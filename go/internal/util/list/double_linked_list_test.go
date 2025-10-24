package list

import (
	"testing"
	// Optional: You might use an assertion library like 'testify/assert' later,
	// but for now, we stick to the standard library as it's the Go idiom.
)

// Helper: Define a type for the test subject to handle different Generics easily
var list List[int] = NewDoubleLinkedList[int]()

// TestNewListIsEmpty checks the initial state of a newly created list.
func TestNewListIsEmpty(t *testing.T) {
	// Arrange
	list := NewDoubleLinkedList[int]()

	// Assert
	if !list.IsEmpty() {
		t.Errorf("New list should be empty, but IsEmpty returned false")
	}
	if list.Size() != 0 {
		t.Errorf("New list size should be 0, but got %d", list.Size())
	}
}

// TestAddElements checks the Add method and size update.
func TestAddElements(t *testing.T) {
	list := NewDoubleLinkedList[int]()

	list.Add(1)
	if list.IsEmpty() {
		t.Errorf("List should not be empty after Add")
	}
	if list.Size() != 1 {
		t.Errorf("Size expected 1, got %d", list.Size())
	}

	list.Add(2)
	list.Add(3)
	if list.Size() != 3 {
		t.Errorf("Size expected 3, got %d", list.Size())
	}
}

// --- TABLE-DRIVEN TESTS for Get and Accessors ---

// TestListAccessors groups tests for Get, Start, End, and GetFromEnd.
func TestListAccessors(t *testing.T) {
	list := NewDoubleLinkedList[int]()
	list.Add(10)
	list.Add(20)
	list.Add(30)

	// Sub-test for GetByIndex
	t.Run("GetByIndex", func(t *testing.T) {
		if val, _ := list.Get(0); val != 10 {
			t.Errorf("Get(0) expected 10, got %d", val)
		}
		if val, _ := list.Get(1); val != 20 {
			t.Errorf("Get(1) expected 20, got %d", val)
		}
		if val, _ := list.Get(2); val != 30 {
			t.Errorf("Get(2) expected 30, got %d", val)
		}
		// NOTE: In Go, accessing invalid index usually results in a panic,
		// which should be tested with 'recover' or handled by the function.
	})

	// Sub-test for Start()
	t.Run("Start", func(t *testing.T) {
		if val, _ := list.Start(); val != 10 {
			t.Errorf("Start() expected 10, got %d", val)
		}
	})

	// Sub-test for End()
	t.Run("End", func(t *testing.T) {
		if val, _ := list.End(); val != 30 {
			t.Errorf("End() expected 30, got %d", val)
		}
	})

	// Sub-test for GetFromEnd
	t.Run("GetFromEnd", func(t *testing.T) {
		if val, _ := list.GetFromEnd(0); val != 30 {
			t.Errorf("GetFromEnd(0) expected 30, got %d", val)
		}
		if val, _ := list.GetFromEnd(1); val != 20 {
			t.Errorf("GetFromEnd(1) expected 20, got %d", val)
		}
		if val, _ := list.GetFromEnd(2); val != 10 {
			t.Errorf("GetFromEnd(2) expected 10, got %d", val)
		}
	})
}

// --- TABLE-DRIVEN TESTS for Remove ---

func TestRemove(t *testing.T) {
	tests := []struct {
		name          string
		initialData   []int
		removeIndex   int
		expectedSize  int
		expectedFinal []int // Expected list contents after removal
	}{
		{
			name:          "Remove Middle Element",
			initialData:   []int{1, 2, 3},
			removeIndex:   1,
			expectedSize:  2,
			expectedFinal: []int{1, 3},
		},
		{
			name:          "Remove First Element",
			initialData:   []int{1, 2, 3},
			removeIndex:   0,
			expectedSize:  2,
			expectedFinal: []int{2, 3},
		},
		{
			name:          "Remove Last Element",
			initialData:   []int{1, 2, 3},
			removeIndex:   2,
			expectedSize:  2,
			expectedFinal: []int{1, 2},
		},
		{
			name:          "Remove Single Element",
			initialData:   []int{42},
			removeIndex:   0,
			expectedSize:  0,
			expectedFinal: []int{},
		},
		// NOTE on invalid index: Go functions often handle invalid operations
		// by returning an error or using a panic. We will rely on the implementation
		// to either silently ignore or panic, and test the panic behavior separately.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewDoubleLinkedList[int]()
			for _, val := range tt.initialData {
				list.Add(val)
			}

			// Act
			list.Remove(tt.removeIndex)

			// Assert size
			if list.Size() != tt.expectedSize {
				t.Errorf("Size expected %d, got %d", tt.expectedSize, list.Size())
			}

			// Assert content by iterating over the expected final data
			if list.Size() > 0 {
				for i, expected := range tt.expectedFinal {
					if actual, _ := list.Get(i); actual != expected {
						t.Errorf("Index %d expected %d, got %d", i, expected, actual)
					}
				}
			} else if !list.IsEmpty() {
				t.Error("List should be empty, but IsEmpty returned false")
			}
		})
	}
}

// --- Tests for RemoveFromEnd (Stack Pop) ---

func TestRemoveFromEnd(t *testing.T) {
	// Sub-test for normal operation
	t.Run("NormalOperations", func(t *testing.T) {
		list := NewDoubleLinkedList[int]()
		list.Add(1)
		list.Add(2)
		list.Add(3)

		if val, _ := list.RemoveFromEnd(); val != 3 {
			t.Errorf("Expected 3, got %d", val)
		}
		if list.Size() != 2 {
			t.Errorf("Size expected 2, got %d", list.Size())
		}
		if val, _ := list.End(); val != 2 {
			t.Errorf("End expected 2, got %d", val)
		}

		if val, _ := list.RemoveFromEnd(); val != 2 {
			t.Errorf("Expected 2, got %d", val)
		}
		if val, _ := list.RemoveFromEnd(); val != 1 {
			t.Errorf("Expected 1, got %d", val)
		}
		if !list.IsEmpty() {
			t.Error("List should be empty")
		}
	})

	// Sub-test for single element case
	t.Run("SingleElement", func(t *testing.T) {
		list := NewDoubleLinkedList[int]()
		list.Add(42)
		if val, _ := list.RemoveFromEnd(); val != 42 {
			t.Errorf("Expected 42, got %d", val)
		}
		if !list.IsEmpty() {
			t.Error("List should be empty after single removal")
		}
	})

	// Test for Empty List (Handle zero-value return or panic)
	// NOTE: Since the Java version returns null for empty list,
	// we assume Go should return the zero value for the generic type T (which is 0 for int).
	t.Run("EmptyListZeroValue", func(t *testing.T) {
		list := NewDoubleLinkedList[int]()
		var expectedZero int // 0 for int

		if val, _ := list.RemoveFromEnd(); val != expectedZero {
			t.Errorf("Expected zero value %d, got %d", expectedZero, val)
		}
		if list.Size() != 0 {
			t.Errorf("Size expected 0, got %d", list.Size())
		}
	})
}

// --- Iterator Tests (Requires implementing the Iterator method) ---

func TestIterator(t *testing.T) {
	list := NewDoubleLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	t.Run("TraversesAllElements", func(t *testing.T) {
		it := list.Iterator()
		expected := []int{1, 2, 3}

		for _, expect := range expected {
			if !it.HasNext() {
				t.Fatalf("HasNext() returned false unexpectedly")
			}
			if nextVal := it.Next(); nextVal != expect {
				t.Errorf("Next() expected %d, got %d", expect, nextVal)
			}
		}

		if it.HasNext() {
			t.Error("HasNext() should be false after last element")
		}
	})

	t.Run("OnEmptyList", func(t *testing.T) {
		emptyList := NewDoubleLinkedList[int]()
		it := emptyList.Iterator()
		if it.HasNext() {
			t.Error("HasNext() on empty list should be false")
		}
		// NOTE: Testing for NoSuchElementException (panic in Go) should be done
		// with a defer/recover block, but we simplify the test for now.
	})
}

// TestClear functionality
func TestClear(t *testing.T) {
	list := NewDoubleLinkedList[int]()
	list.Add(1)
	list.Add(2)

	list.Clear()

	if !list.IsEmpty() {
		t.Error("List should be empty after Clear()")
	}
	if list.Size() != 0 {
		t.Errorf("Size expected 0 after Clear(), got %d", list.Size())
	}
}
