package stack

import (
    "testing"
)

// TestNewStackIsEmpty checks the initial state of a newly created stack.
func TestNewStackIsEmpty(t *testing.T) {
    // Arrange
    stack := NewDoubleLinkedListStack[int]()

    // Assert
    if !stack.IsEmpty() {
        t.Errorf("New stack should be empty, but IsEmpty returned false")
    }
    if stack.Size() != 0 {
        t.Errorf("New stack size should be 0, but got %d", stack.Size())
    }
}

// TestPushElements checks the Push method and size update.
func TestPushElements(t *testing.T) {
    stack := NewDoubleLinkedListStack[int]()

    stack.Push(1)
    if stack.IsEmpty() {
        t.Errorf("Stack should not be empty after Push")
    }
    if stack.Size() != 1 {
        t.Errorf("Size expected 1, got %d", stack.Size())
    }

    stack.Push(2)
    stack.Push(3)
    if stack.Size() != 3 {
        t.Errorf("Size expected 3, got %d", stack.Size())
    }
}

// TestPeek checks that Peek returns the top element without removing it.
func TestPeek(t *testing.T) {
    stack := NewDoubleLinkedListStack[int]()

    // Test peek on empty stack
    t.Run("PeekEmptyStack", func(t *testing.T) {
        val, ok := stack.Peek()
        if ok {
            t.Errorf("Peek on empty stack should return false, got true with value %d", val)
        }
    })

    // Test peek on non-empty stack
    t.Run("PeekNonEmptyStack", func(t *testing.T) {
        stack.Push(10)
        stack.Push(20)
        stack.Push(30)

        val, ok := stack.Peek()
        if !ok {
            t.Errorf("Peek should return true for non-empty stack")
        }
        if val != 30 {
            t.Errorf("Peek expected 30, got %d", val)
        }

        // Verify size hasn't changed
        if stack.Size() != 3 {
            t.Errorf("Size should remain 3 after Peek, got %d", stack.Size())
        }
    })
}

// TestPop checks the Pop method (LIFO behavior).
func TestPop(t *testing.T) {
    t.Run("PopFromEmptyStack", func(t *testing.T) {
        stack := NewDoubleLinkedListStack[int]()
        val, ok := stack.Pop()
        if ok {
            t.Errorf("Pop on empty stack should return false, got true with value %d", val)
        }
        if stack.Size() != 0 {
            t.Errorf("Size should remain 0, got %d", stack.Size())
        }
    })

    t.Run("PopLIFOOrder", func(t *testing.T) {
        stack := NewDoubleLinkedListStack[int]()
        stack.Push(1)
        stack.Push(2)
        stack.Push(3)

        // Pop should return elements in LIFO order: 3, 2, 1
        val, ok := stack.Pop()
        if !ok || val != 3 {
            t.Errorf("First Pop expected 3, got %d (ok=%v)", val, ok)
        }
        if stack.Size() != 2 {
            t.Errorf("Size expected 2, got %d", stack.Size())
        }

        val, ok = stack.Pop()
        if !ok || val != 2 {
            t.Errorf("Second Pop expected 2, got %d (ok=%v)", val, ok)
        }
        if stack.Size() != 1 {
            t.Errorf("Size expected 1, got %d", stack.Size())
        }

        val, ok = stack.Pop()
        if !ok || val != 1 {
            t.Errorf("Third Pop expected 1, got %d (ok=%v)", val, ok)
        }
        if !stack.IsEmpty() {
            t.Error("Stack should be empty after all pops")
        }
    })

    t.Run("PopSingleElement", func(t *testing.T) {
        stack := NewDoubleLinkedListStack[int]()
        stack.Push(42)

        val, ok := stack.Pop()
        if !ok || val != 42 {
            t.Errorf("Pop expected 42, got %d (ok=%v)", val, ok)
        }
        if !stack.IsEmpty() {
            t.Error("Stack should be empty after popping single element")
        }
    })
}

// TestPushPopSequence tests interleaved push and pop operations.
func TestPushPopSequence(t *testing.T) {
    stack := NewDoubleLinkedListStack[int]()

    stack.Push(1)
    stack.Push(2)

    val, _ := stack.Pop()
    if val != 2 {
        t.Errorf("Expected 2, got %d", val)
    }

    stack.Push(3)
    stack.Push(4)

    val, _ = stack.Peek()
    if val != 4 {
        t.Errorf("Peek expected 4, got %d", val)
    }

    val, _ = stack.Pop()
    if val != 4 {
        t.Errorf("Expected 4, got %d", val)
    }

    val, _ = stack.Pop()
    if val != 3 {
        t.Errorf("Expected 3, got %d", val)
    }

    val, _ = stack.Pop()
    if val != 1 {
        t.Errorf("Expected 1, got %d", val)
    }

    if !stack.IsEmpty() {
        t.Error("Stack should be empty")
    }
}

// TestStackWithStrings tests the stack with a different type.
func TestStackWithStrings(t *testing.T) {
    stack := NewDoubleLinkedListStack[string]()

    stack.Push("first")
    stack.Push("second")
    stack.Push("third")

    if val, _ := stack.Peek(); val != "third" {
        t.Errorf("Peek expected 'third', got '%s'", val)
    }

    if val, _ := stack.Pop(); val != "third" {
        t.Errorf("Pop expected 'third', got '%s'", val)
    }

    if val, _ := stack.Pop(); val != "second" {
        t.Errorf("Pop expected 'second', got '%s'", val)
    }

    if stack.Size() != 1 {
        t.Errorf("Size expected 1, got %d", stack.Size())
    }

    if val, _ := stack.Pop(); val != "first" {
        t.Errorf("Pop expected 'first', got '%s'", val)
    }

    if !stack.IsEmpty() {
        t.Error("Stack should be empty")
    }
}

// TestStackInterfaceImplementation verifies that DoubleLinkedListStack implements Stack interface.
func TestStackInterfaceImplementation(t *testing.T) {
    var stack Stack[int] = NewDoubleLinkedListStack[int]()

    stack.Push(100)
    stack.Push(200)

    if val, ok := stack.Peek(); !ok || val != 200 {
        t.Errorf("Expected 200, got %d (ok=%v)", val, ok)
    }

    // It will panic if we try to pop from an empty stack.
    _ = stack.Pop()

    if stack.Size() != 1 {
        t.Errorf("Size expected 1, got %d", stack.Size())
    }
}

// TestLargeStack tests stack behavior with many elements.
func TestLargeStack(t *testing.T) {
    stack := NewDoubleLinkedListStack[int]()
    count := 1000

    // Push many elements
    for i := 0; i < count; i++ {
        stack.Push(i)
    }

    if stack.Size() != count {
        t.Errorf("Size expected %d, got %d", count, stack.Size())
    }

    // Pop and verify LIFO order
    for i := count - 1; i >= 0; i-- {
        val, ok := stack.Pop()
        if !ok {
            t.Fatalf("Pop failed at iteration %d", i)
        }
        if val != i {
            t.Errorf("Expected %d, got %d", i, val)
        }
    }

    if !stack.IsEmpty() {
        t.Error("Stack should be empty after popping all elements")
    }
}
