package tree

// Tree is a generic interface representing the structure and operations of a tree.
//
// Example Tree Structure:
//
//             10
//            /  \
//           5    15
//          / \     \
//         3   7     20
//
// In the tree example:
// - The root value is 10.
// - 5 and 15 are the children of 10.
// - 3 and 7 are the left and right children of 5, respectively.
// - 20 is the right child of 15.
//
// Tree defines fundamental methods for managing hierarchical data structures, such as insertion,
// search, removal, and clearing operations. All implementing types must ensure that these operations
// adhere to the tree's specific rules (e.g., ordering for a Binary Search Tree or balancing for AVL trees).
type Tree[T any] interface {
    // Root retrieves the value stored at the root node of the tree, or nil if the tree is empty.
    // Time complexity: O(1)
    Root() (T, bool)

    // Add inserts a value into the tree.
    // The rules for insertion (e.g., maintaining order or balance) are determined
    // by the specific implementation. For instance, in a Binary Search Tree, values
    // are inserted such that the tree remains ordered.
    // Time complexity:
    // - Binary Search Tree: O(h) where h is the height
    //   - Balanced BST: O(log n)
    //   - Skewed BST: O(n)
    // - AVL/Red-Black Tree: O(log n) guaranteed
    Add(value T)

    // Search looks for a value in the tree.
    // Returns a pointer to the value if found, or nil if the value does not exist in the tree.
    //
    // Time complexity:
    // - Binary Search Tree: O(h) where h is the height
    //   - Balanced BST: O(log n)
    //   - Skewed BST: O(n)
    // - AVL/Red-Black Tree: O(log n) guaranteed
    Search(key T) (T, bool)

    // Remove deletes a value from the tree.
    // Removal may impact the tree's structure. Implementing types must ensure that
    // the tree maintains its properties (e.g., ordering for Binary Search Trees, balancing for AVL trees).
    // Returns a pointer to the removed value if it was successfully deleted, or nil if the value was not found.
    //
    // Time complexity:
    // - Binary Search Tree: O(h) where h is the height
    //   - Balanced BST: O(log n)
    //   - Skewed BST: O(n)
    // - AVL/Red-Black Tree: O(log n) guaranteed
    Remove(value T) (T, bool)

    // Clear removes all elements from the tree.
    // After clearing, the tree will be empty, and Root() will return nil.
    // Time complexity: O(1)
    Clear()

    // Iterator retrieves an iterator for traversing the tree elements in its specific order.
    //
    // Example Traversal:
    // Take this tree:
    //            10
    //         /     \
    //        5       15
    //       / \     /  \
    //      3   7   12  20
    //
    // Using an in-order traversal (left-root-right):
    //  - First, visit all nodes on the left subtree.
    //  - Process the root node.
    //  - Then, visit all nodes on the right subtree.
    //
    // Output: 3, 5, 7, 10, 12, 15, 20.
    // Time complexity: O(1) to create the iterator
    // Space complexity: O(h) where h is the height (for the traversal stack)
    // Traversing all elements: O(n) where n is the number of nodes
    Iterator() Iterator[T]
}

// Iterator is a generic interface for traversing elements of the tree.
type Iterator[T any] interface {
    // HasNext returns true if there are more elements to traverse.
    HasNext() bool

    // Next returns the next sequential element in the tree.
    // If called when there are no remaining elements, it panics.
    Next() T
}
