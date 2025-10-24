package list

import (
	"interview_go/internal/util/collection"
	"interview_go/internal/util/iterator"
)

// node[T] is the private structure representing an element in the list.
// Note: It starts with a lowercase letter, making it unexported (private)
// outside the 'list' package.
type node[T any] struct {
	value T
	prior *node[T] // Pointer to the previous node
	next  *node[T] // Pointer to the next node
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		value: value,
		prior: nil,
		next:  nil,
	}
}

// DoubleLinkedList implements the List[T] interface. It implements generics T
type DoubleLinkedList[T any] struct {
	head *node[T] // Start of the list
	tail *node[T] // End of the list
	size int
}

// NewDoubleLinkedList is the idiomatic Go "constructor" to initialize the list.
func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	// Returns a pointer to a zero-valued DoubleLinkedList struct.
	// head, tail, and size are automatically initialized to their zero values (nil, nil, 0).
	return &DoubleLinkedList[T]{}
}

// Compilation assert: force contract
var _ collection.Collection[any] = (*DoubleLinkedList[any])(nil)
var _ List[any] = (*DoubleLinkedList[any])(nil)

// Add element to list
func (d *DoubleLinkedList[T]) Add(value T) {
	if any(value) == nil {
		return
	}

	node := newNode[T](value)
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.prior = d.tail
		d.tail.next = node
		d.tail = node
	}
	d.size++
}

// Get the element from the nth position
func (d *DoubleLinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= d.size {
		return *new(T), false
	}

	node := d.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value, true
}

// GetFromEnd get the nth-element from the end
func (d *DoubleLinkedList[T]) GetFromEnd(index int) (T, bool) {
	if index < 0 || index >= d.size {
		return *new(T), false
	}

	node := d.tail
	for i := 0; i < index; i++ {
		node = node.prior
	}
	return node.value, true
}

// Size return the size of the list
func (d *DoubleLinkedList[T]) Size() int {
	return d.size
}

// Clear removes all elements
func (d *DoubleLinkedList[T]) Clear() {
	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *DoubleLinkedList[T]) IsEmpty() bool {
	return d.size == 0
}

// Remove the node after the current node by updating the links.
//
// Before:
//
//	   //
//		       node       toRemove     nextNode
//		        ↓            ↓            ↓
//	       ←→ [A]    ←→    [B]    ←→    [C]    ←→ ...
//
// Step 1: node.next = node.next.next;
//
//	   //
//		       node             nextNode
//		        ↓                  ↓
//		   ←→ [A] ─────────────→ [C] ←→ ...
//		           ←─ [B] ←──────┘
//		           (orphaned)
//
// Step 2: node.next.prior = node;
//
//	   //
//		       node           nextNode
//		        ↓                ↓
//		   ←→ [A] ←──────────→ [C] ←→ ...
//		           [B] (removed, no references)
//
// After:
//
//	   //
//		     node       nextNode
//		      ↓            ↓
//		   ←→ [A]   ←→   [C] ←→ ...
//
// Node [B] is now unreachable and will be garbage collected.
func (d *DoubleLinkedList[T]) Remove(index int) bool {
	if index < 0 || index >= d.size {
		return false
	}
	if index == 0 {
		// special case, delete head
		d.head = d.head.next
	} else {
		node := d.head
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		node.next.prior = node
		node.next = node.next.next
	}
	d.size--
	return true
}

func (d *DoubleLinkedList[T]) RemoveFromEnd() (T, bool) {
	if d.IsEmpty() {
		return *new(T), false
	}
	data := d.tail.value
	d.tail = d.tail.prior
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.size--
	return data, true
}

func (d *DoubleLinkedList[T]) End() (T, bool) {
	if d.IsEmpty() {
		return *new(T), false
	}
	return d.tail.value, true
}

func (d *DoubleLinkedList[T]) Start() (T, bool) {
	if d.IsEmpty() {
		return *new(T), false
	}
	return d.head.value, true
}

// Iterator implements an iterable pattern to traverse the list
func (d *DoubleLinkedList[T]) Iterator() iterator.Iterator[T] {
	return newDoubleLinkedListIterator(d.head)
}

type doubleLinkedListIterator[T any] struct {
	current *node[T]
}

func newDoubleLinkedListIterator[T any](startNode *node[T]) iterator.Iterator[T] {
	return &doubleLinkedListIterator[T]{current: startNode}
}

var _ iterator.Iterator[any] = (*doubleLinkedListIterator[any])(nil)

func (d *doubleLinkedListIterator[T]) HasNext() bool {
	if d.current == nil {
		return false
	}
	return true
}

// Next Returns the next element in the iteration.
//
//	   //
//
//	     head                 tail
//	       ↓                    ↓
//	      [A] ←→ [B] ←→ [C] ←→ [D]
//	       ↑
//		 current (start)
//
//		After next(): returns A
//		        ↓
//		[A] ←→ [B] ←→ [C] ←→ [D]
//		        ↑
//		    current
//
//		 After next(): returns B
//		                ↓
//		 [A] ←→ [B] ←→ [C] ←→ [D]
//		                ↑
//		             current
func (d *doubleLinkedListIterator[T]) Next() T {
	if d.current == nil {
		panic("empty")
	}
	data := d.current.value
	d.current = d.current.next
	return data
}
