package tree

import (
    "testing"
)

// TestInOrderIteratorBasic tests basic iterator functionality with a simple tree.
func TestInOrderIteratorBasic(t *testing.T) {
    tree := NewBinaryTree[int]()
    tree.Add(10)
    tree.Add(5)
    tree.Add(15)
    tree.Add(3)
    tree.Add(7)
    tree.Add(12)
    tree.Add(20)

    // Tree structure:
    //            10
    //         /     \
    //        5       15
    //       / \     /  \
    //      3   7   12  20
    //
    // Expected in-order: 3, 5, 7, 10, 12, 15, 20

    t.Run("TraversesAllElementsInOrder", func(t *testing.T) {
        it := tree.Iterator()
        expected := []int{3, 5, 7, 10, 12, 15, 20}

        for i, expect := range expected {
            if !it.HasNext() {
                t.Fatalf("HasNext() returned false at iteration %d, expected more elements", i)
            }
            nextVal := it.Next()
            if nextVal != expect {
                t.Errorf("Next() at position %d expected %d, got %d", i, expect, nextVal)
            }
        }

        if it.HasNext() {
            t.Error("HasNext() should be false after last element")
        }
    })
}

// TestInOrderIteratorEmptyTree tests iterator on an empty tree.
func TestInOrderIteratorEmptyTree(t *testing.T) {
    tree := NewBinaryTree[int]()

    t.Run("EmptyTreeHasNoElements", func(t *testing.T) {
        it := tree.Iterator()
        if it.HasNext() {
            t.Error("HasNext() on empty tree should be false")
        }
    })

    t.Run("NextOnEmptyTreePanics", func(t *testing.T) {
        it := tree.Iterator()

        defer func() {
            if r := recover(); r == nil {
                t.Error("Next() on empty iterator should panic")
            }
        }()

        it.Next() // Should panic
    })
}

// TestInOrderIteratorSingleNode tests iterator with a tree containing only the root.
func TestInOrderIteratorSingleNode(t *testing.T) {
    tree := NewBinaryTree[int]()
    tree.Add(42)

    t.Run("SingleNodeTraversal", func(t *testing.T) {
        it := tree.Iterator()

        if !it.HasNext() {
            t.Error("HasNext() should be true for single-node tree")
        }

        val := it.Next()
        if val != 42 {
            t.Errorf("Expected 42, got %d", val)
        }

        if it.HasNext() {
            t.Error("HasNext() should be false after single element")
        }
    })
}

// TestInOrderIteratorLeftSkewedTree tests iterator with a left-skewed tree.
func TestInOrderIteratorLeftSkewedTree(t *testing.T) {
    tree := NewBinaryTree[int]()
    tree.Add(50)
    tree.Add(40)
    tree.Add(30)
    tree.Add(20)
    tree.Add(10)

    // Tree structure (left-skewed):
    //     50
    //    /
    //   40
    //  /
    // 30
    // /
    // 20
    // /
    // 10
    //
    // Expected in-order: 10, 20, 30, 40, 50

    t.Run("LeftSkewedTraversal", func(t *testing.T) {
        it := tree.Iterator()
        expected := []int{10, 20, 30, 40, 50}

        for i, expect := range expected {
            if !it.HasNext() {
                t.Fatalf("HasNext() returned false at iteration %d", i)
            }
            if val := it.Next(); val != expect {
                t.Errorf("Expected %d, got %d", expect, val)
            }
        }

        if it.HasNext() {
            t.Error("HasNext() should be false after traversal")
        }
    })
}

// TestInOrderIteratorRightSkewedTree tests iterator with a right-skewed tree.
func TestInOrderIteratorRightSkewedTree(t *testing.T) {
    tree := NewBinaryTree[int]()
    tree.Add(10)
    tree.Add(20)
    tree.Add(30)
    tree.Add(40)
    tree.Add(50)

    // Tree structure (right-skewed):
    // 10
    //  \
    //   20
    //    \
    //     30
    //      \
    //       40
    //        \
    //         50
    //
    // Expected in-order: 10, 20, 30, 40, 50

    t.Run("RightSkewedTraversal", func(t *testing.T) {
        it := tree.Iterator()
        expected := []int{10, 20, 30, 40, 50}

        for i, expect := range expected {
            if !it.HasNext() {
                t.Fatalf("HasNext() returned false at iteration %d", i)
            }
            if val := it.Next(); val != expect {
                t.Errorf("Expected %d, got %d", expect, val)
            }
        }

        if it.HasNext() {
            t.Error("HasNext() should be false after traversal")
        }
    })
}

// TestInOrderIteratorBalancedTree tests iterator with a balanced tree.
func TestInOrderIteratorBalancedTree(t *testing.T) {
    tree := NewBinaryTree[int]()
    tree.Add(50)
    tree.Add(30)
    tree.Add(70)
    tree.Add(20)
    tree.Add(40)
    tree.Add(60)
    tree.Add(80)

    // Tree structure:
    //          50
    //         /  \
    //        /    \
    //      30      70
    //     /  \    /  \
    //    20  40  60  80
    //
    // Expected in-order: 20, 30, 40, 50, 60, 70, 80

    t.Run("BalancedTreeTraversal", func(t *testing.T) {
        it := tree.Iterator()
        expected := []int{20, 30, 40, 50, 60, 70, 80}

        for i, expect := range expected {
            if !it.HasNext() {
                t.Fatalf("HasNext() returned false at iteration %d", i)
            }
            if val := it.Next(); val != expect {
                t.Errorf("Expected %d, got %d", expect, val)
            }
        }

        if it.HasNext() {
            t.Error("HasNext() should be false after traversal")
        }
    })
}

// TestInOrderIteratorWithStrings tests iterator with string values.
func TestInOrderIteratorWithStrings(t *testing.T) {
    tree := NewBinaryTree[string]()
    tree.Add("dog")
    tree.Add("cat")
    tree.Add("elephant")
    tree.Add("ant")
    tree.Add("bird")

    // Expected in-order: ant, bird, cat, dog, elephant (alphabetical)

    t.Run("StringTreeTraversal", func(t *testing.T) {
        it := tree.Iterator()
        expected := []string{"ant", "bird", "cat", "dog", "elephant"}

        for i, expect := range expected {
            if !it.HasNext() {
                t.Fatalf("HasNext() returned false at iteration %d", i)
            }
            if val := it.Next(); val != expect {
                t.Errorf("Expected %s, got %s", expect, val)
            }
        }

        if it.HasNext() {
            t.Error("HasNext() should be false after traversal")
        }
    })
}

// TestInOrderIteratorMultipleIterators tests that multiple iterators can coexist.
func TestInOrderIteratorMultipleIterators(t *testing.T) {
    tree := NewBinaryTree[int]()
    tree.Add(10)
    tree.Add(5)
    tree.Add(15)

    t.Run("MultipleIteratorsIndependent", func(t *testing.T) {
        it1 := tree.Iterator()
        it2 := tree.Iterator()

        // Advance first iterator
        if val := it1.Next(); val != 5 {
            t.Errorf("Iterator 1 expected 5, got %d", val)
        }

        // Second iterator should still start from beginning
        if val := it2.Next(); val != 5 {
            t.Errorf("Iterator 2 expected 5, got %d", val)
        }

        // Advance both
        if val := it1.Next(); val != 10 {
            t.Errorf("Iterator 1 expected 10, got %d", val)
        }
        if val := it2.Next(); val != 10 {
            t.Errorf("Iterator 2 expected 10, got %d", val)
        }
    })
}

// TestInOrderIteratorPanicOnExhausted tests that calling Next() on exhausted iterator panics.
func TestInOrderIteratorPanicOnExhausted(t *testing.T) {
    tree := NewBinaryTree[int]()
    tree.Add(10)
    tree.Add(5)

    t.Run("NextPanicsWhenExhausted", func(t *testing.T) {
        it := tree.Iterator()

        // Exhaust the iterator
        it.Next() // 5
        it.Next() // 10

        if it.HasNext() {
            t.Error("HasNext() should be false after exhausting iterator")
        }

        // Try to call Next() again - should panic
        defer func() {
            if r := recover(); r == nil {
                t.Error("Next() on exhausted iterator should panic")
            }
        }()

        it.Next() // Should panic
    })
}

// TestInOrderIteratorComplexTree tests iterator with a more complex tree structure.
func TestInOrderIteratorComplexTree(t *testing.T) {
    tree := NewBinaryTree[int]()
    values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85}
    for _, v := range values {
        tree.Add(v)
    }

    // Tree structure:
    //                50
    //         /              \
    //        30              70
    //       /  \            /  \
    //     20    40        60    80
    //    /  \   / \      /  \   / \
    //   10  25 35 45    55 65  75 85
    //
    // Expected in-order: 10, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85

    t.Run("ComplexTreeTraversal", func(t *testing.T) {
        it := tree.Iterator()
        expected := []int{10, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85}

        actualValues := []int{}
        for it.HasNext() {
            actualValues = append(actualValues, it.Next())
        }

        if len(actualValues) != len(expected) {
            t.Errorf("Expected %d elements, got %d", len(expected), len(actualValues))
        }

        for i, expect := range expected {
            if i >= len(actualValues) {
                t.Errorf("Missing element at position %d", i)
                continue
            }
            if actualValues[i] != expect {
                t.Errorf("At position %d: expected %d, got %d", i, expect, actualValues[i])
            }
        }
    })
}

// TestInOrderIteratorDuplicateValues tests that duplicate values don't break the iterator.
func TestInOrderIteratorDuplicateValues(t *testing.T) {
    tree := NewBinaryTree[int]()
    tree.Add(10)
    tree.Add(5)
    tree.Add(15)
    tree.Add(10) // Duplicate - should not be added

    t.Run("DuplicatesIgnored", func(t *testing.T) {
        it := tree.Iterator()
        expected := []int{5, 10, 15} // Only unique values

        actualValues := []int{}
        for it.HasNext() {
            actualValues = append(actualValues, it.Next())
        }

        if len(actualValues) != len(expected) {
            t.Errorf("Expected %d elements, got %d", len(expected), len(actualValues))
        }

        for i, expect := range expected {
            if actualValues[i] != expect {
                t.Errorf("At position %d: expected %d, got %d", i, expect, actualValues[i])
            }
        }
    })
}

// TestInOrderIteratorLargeTree tests iterator with a large tree.
func TestInOrderIteratorLargeTree(t *testing.T) {
    tree := NewBinaryTree[int]()
    count := 1000

    // Add elements in a mixed order
    for i := 0; i < count; i += 2 {
        tree.Add(i)
    }
    for i := 1; i < count; i += 2 {
        tree.Add(i)
    }

    t.Run("LargeTreeTraversal", func(t *testing.T) {
        it := tree.Iterator()

        // Verify sorted order
        prev := -1
        elementCount := 0
        for it.HasNext() {
            current := it.Next()
            if current <= prev {
                t.Errorf("Elements not in sorted order: %d followed by %d", prev, current)
            }
            prev = current
            elementCount++
        }

        if elementCount != count {
            t.Errorf("Expected %d elements, got %d", count, elementCount)
        }
    })
}
