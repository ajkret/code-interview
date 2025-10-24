package stack

import "interview_go/internal/util/list"

// DoubleLinkedListStack implementation of a Stack using the DoubleLinkedList
type DoubleLinkedListStack[T any] struct {
    // Our stack is actually a double-linked list
    stack list.DoubleLinkedList[T]
}

// NewDoubleLinkedListStack Constructor
func NewDoubleLinkedListStack[T any]() *DoubleLinkedListStack[T] {
    return &DoubleLinkedListStack[T]{
        stack: *list.NewDoubleLinkedList[T](),
    }
}

var _ Stack[any] = (*DoubleLinkedListStack[any])(nil)

func (d *DoubleLinkedListStack[T]) Push(t T) {
    d.stack.Add(t)
}

func (d *DoubleLinkedListStack[T]) Pop() T {
    val, hasValue := d.stack.RemoveFromEnd()
    if hasValue {
        return val
    }
    panic("empty")
}

func (d *DoubleLinkedListStack[T]) Peek() (T, bool) {
    return d.stack.End()
}

func (d *DoubleLinkedListStack[T]) Size() int {
    return d.stack.Size()
}

func (d *DoubleLinkedListStack[T]) IsEmpty() bool {
    return d.stack.IsEmpty()
}
