package stack

// Stack basic implementation of a stack
type Stack[T any] interface {
    // Push adds an element to the top of the Stack
    Push(T)

    // Pop Retrieves the element from the top
    Pop() T

    // Peek check element on top of the stack
    Peek() (T, bool)

    // Size Stack "height"
    Size() int

    // IsEmpty Checks if the stack is empty
    IsEmpty() bool
}
