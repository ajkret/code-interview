package tree

import (
    "interview_go/internal/util/stack"

    "golang.org/x/exp/constraints"
)

// Node represents a single node in a binary tree.
// Each node contains a value and pointers to its left and right children.
//
// Example node structure:
//
//        [50]
//       /    \
//    [30]    [70]
//
type Node[T constraints.Ordered] struct {
    value T
    left  *Node[T]
    right *Node[T]
}

// newNode creates and returns a new node with the given value.
// The left and right children are initialized to nil.
func newNode[T constraints.Ordered](value T) *Node[T] {
    return &Node[T]{value: value}
}

// BinaryTree is a generic Binary Search Tree implementation.
// It maintains the BST property: for any node, all values in the left subtree
// are less than the node's value, and all values in the right subtree are greater.
//
// Example tree structure:
//
//           50
//          /  \
//         /    \
//       30      70
//      /  \    /  \
//     20  40  60  80
//
// Properties:
// - Left subtree values < parent value
// - Right subtree values > parent value
// - No duplicate values allowed
//
type BinaryTree[T constraints.Ordered] struct {
    root *Node[T]
    size int
}

// NewBinaryTree creates and returns a new empty binary tree.
func NewBinaryTree[T constraints.Ordered]() *BinaryTree[T] {
    return &BinaryTree[T]{}
}

// Compile-time check to ensure BinaryTree implements the Tree interface
var _ Tree[string] = (*BinaryTree[string])(nil)

// Root returns the value at the root of the tree.
// Returns the zero value and false if the tree is empty, otherwise returns the root value and true.
func (b *BinaryTree[T]) Root() (T, bool) {
    if b.root == nil {
        return *new(T), true
    }
    return b.root.value, true
}

// Add inserts a new value into the binary search tree while maintaining BST properties.
// If the value already exists, it is not added again (no duplicates allowed).
//
// Example: Adding 25 to this tree:
//
// Before:
//           50
//          /  \
//        30    70
//       /  \
//      20  40
//
// After:
//           50
//          /  \
//        30    70
//       /  \
//      20  40
//     /
//    25
//
// Time complexity: O(h) where h is the height of the tree
// In a balanced tree: O(log n), in worst case (skewed): O(n)
//
func (b *BinaryTree[T]) Add(value T) {
    // Commits the size before adding the value
    b.size += 1

    // Special case, empty tree
    if b.root == nil {
        b.root = newNode(value)
        return
    }

    node := b.root
    for {
        // Special case - value already exists
        if node.value == value {
            return
        }

        if value < node.value {
            if node.left == nil {
                node.left = newNode(value)
                return
            }
            node = node.left
        } else {
            if node.right == nil {
                node.right = newNode(value)
                return
            }
            node = node.right
        }
    }
}

// Search looks for a value in the binary search tree.
// Returns the value and true if found, otherwise returns zero value and false.
//
// Example: Searching for 40 in this tree:
//
//           50
//          /  \
//        30    70
//       /  \
//      20  40  <- Found here
//
// The search path follows: 50 -> 30 -> 40
//
// Time complexity: O(h) where h is the height of the tree
//
func (b *BinaryTree[T]) Search(value T) (T, bool) {
    node := b.root

    for {
        if node == nil {
            return *new(T), false
        }

        if node.value == value {
            return node.value, true
        }

        if value < node.value {
            node = b.root.left
        } else {
            node = b.root.right
        }

        if node == nil {
            return *new(T), false
        }
    }
}

// Remove deletes a value from the binary search tree while maintaining BST properties.
// Returns the removed value and true if found and removed, otherwise returns zero value and false.
//
// There are three cases to handle:
//
// Case 1: Node is a leaf (no children)
//
//        50              50
//       /  \            /  \
//     30    70   =>   30    70
//          /                /
//        60               60
//       /
//     55 <- Remove
//
// Case 2: Node has one child
//
//        50              50
//       /  \            /  \
//     30    70   =>   30    60
//          /
//        60 <- Remove
//
// Case 3: Node has two children (see detailed example below)
//
// Time complexity: O(h) where h is the height of the tree
//
func (b *BinaryTree[T]) Remove(value T) (T, bool) {
    var node = b.root
    var parent *Node[T] = nil
    var parentFromLeft = true

    // First, locate the node to remove
    for node != nil {
        if node.value == value {
            break
        }
        parent = node
        if value < node.value {
            node = node.left
            parentFromLeft = true
        } else {
            node = node.right
            parentFromLeft = false
        }
    }

    // if the node is not found, return false
    if node == nil {
        return *new(T), false
    }

    // Cases where a leaf is empty
    b.size -= 1
    var child *Node[T] = nil
    if node.left == nil && node.right == nil {

        if node.left == nil {
            child = node.right
        } else {
            child = node.left
        }

        // Special case - root
        if parent == nil {
            b.root = child
        } else {
            if parentFromLeft {
                parent.left = child
            } else {
                parent.right = child
            }
        }
    } else {
        // Challenging case: Node has both left and right children
        // We need to find a successor to replace the node being removed.
        //
        // Example: Removing node 70 from this tree:
        //
        //          50
        //         /  \
        //        /    \
        //      30      70  <- Node to remove (has both children)
        //     /  \    /  \
        //    20  40  60  80
        //
        // Strategy: Find the in-order successor (leftmost node in right subtree)
        // In this case, successor is 80 (no left child in right subtree)
        //
        // After removal, tree becomes:
        //
        //          50
        //         /  \
        //        /    \
        //      30      80  <- Successor replaces removed node
        //     /  \    /
        //    20  40  60
        //
        // get the right child and keep searching for the leftmost node
        // use it to replace the node

        successor := node.right
        parent = node // can be root
        for successor.left != nil {
            parent = successor
            successor = successor.left
        }

        // replace the node's value, not the node itself. It is up to the user
        // to use references or pointers
        node.value = successor.value

        // Now delete successor
        if parent.left != nil && parent.left.value == successor.value {
            // Replace a parent's left child with a successor's right child
            parent.left = successor.right
        } else {
            parent.right = successor.right

        }
    }

    // Return value as a result
    return value, true
}

// Clear removes all elements from the tree, leaving it empty.
// After calling Clear, the tree will have no nodes and size will be 0.
func (b *BinaryTree[T]) Clear() {
    b.root = nil
    b.size = 0
}

// Iterator returns an iterator for traversing the tree elements.
// The traversal order depends on the implementation (typically in-order for BST).
//
// Example in-order traversal:
//
//           50
//          /  \
//        30    70
//       /  \  /  \
//      20 40 60  80
//
// Iterator order: 20, 30, 40, 50, 60, 70, 80
//
func (b *BinaryTree[T]) Iterator() Iterator[T] {
    return newInOrderIterator(b.root)
}

// inOrderIterator implements an in-order tree traversal using a stack.
// It visits nodes in the order: left subtree, root, right subtree.
type inOrderIterator[T constraints.Ordered] struct {
    stack stack.Stack[*Node[T]]
}

// newInOrderIterator creates a new in-order iterator starting from the given root node.
func newInOrderIterator[T constraints.Ordered](root *Node[T]) Iterator[T] {
    it := &inOrderIterator[T]{
        stack: stack.NewDoubleLinkedListStack[*Node[T]](),
    }
    it.pushLeft(root)
    return it
}

// Compile-time check to ensure inOrderIterator implements the Iterator interface
var _ Iterator[int] = (*inOrderIterator[int])(nil)

// pushLeft pushes all nodes on the left side of the given node to the stack.
// This effectively traverses the leftmost path from the given node.
//
// Example: Starting from root 10:
//
//        10
//       /  \
//      5    15
//     /
//    3
//
// Stack after pushLeft(10): [10, 5, 3] (3 is on top)
//
func (it *inOrderIterator[T]) pushLeft(node *Node[T]) {
    for node != nil {
        it.stack.Push(node)
        node = node.left
    }
}

// HasNext returns true if there are more elements to traverse.
func (it *inOrderIterator[T]) HasNext() bool {
    return !it.stack.IsEmpty()
}

// Next returns the next element in the in-order traversal.
// Panics if called when there are no remaining elements.
//
// Algorithm explanation:
//
// Take this tree:
//            10
//         /     \
//        5       15
//       / \     /  \
//      3   7   12  20
//
// When the iterator initializes, it pushes the root node to the stack,
// then proceeds to push all nodes on the left side.
// The stack will look like this:
//
// [10, 5, 3] (3 is on top)
//
// Then it pops the last node *3*, saves it in memory (this point is very important)
// and starts to push all nodes in the right subtree.
//
// In our example, 3 has no right leaf. The stack will look like this:
//
// [10, 5]
//
// In the following call to next(), we will pop the last node *5*,
// then push the right subtree. The stack will look like this:
//
// [10, 7]
//
// Next calls:
//
// [10]        -> returns 7
// [15, 12]    -> returns 10
// [15]        -> returns 12
// [20]        -> returns 15
// []          -> returns 20
//
func (it *inOrderIterator[T]) Next() T {
    if it.stack.IsEmpty() {
        // Follow the classic iterator contract
        panic("iterator has no more elements")
    }

    node := it.stack.Pop()
    if node.right != nil {
        it.pushLeft(node.right)
    }
    return node.value
}
