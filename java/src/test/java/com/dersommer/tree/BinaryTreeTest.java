package com.dersommer.tree;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.DisplayName;

import static org.junit.jupiter.api.Assertions.*;

@DisplayName("BinaryTree Tests")
class BinaryTreeTest {

    private BinaryTree<Integer> tree;

    @BeforeEach
    void setUp() {
        tree = new BinaryTree<>();
    }

    @Test
    @DisplayName("Should return null for root of empty tree")
    void testEmptyTreeRoot() {
        assertNull(tree.root());
    }

    @Test
    @DisplayName("Should add root node correctly")
    void testAddRootNode() {
        tree.add(10);
        assertEquals(10, tree.root());
    }

    @Test
    @DisplayName("Should add multiple nodes")
    void testAddMultipleNodes() {
        tree.add(10);
        tree.add(5);
        tree.add(15);
        tree.add(3);
        tree.add(7);

        assertEquals(10, tree.root());
    }

    @Test
    @DisplayName("Should handle duplicate values")
    void testAddDuplicateValues() {
        tree.add(10);
        tree.add(5);
        tree.add(10); // duplicate

        assertEquals(10, tree.root());
    }

    @Test
    @DisplayName("Should search for existing value")
    void testSearchExistingValue() {
        tree.add(10);
        tree.add(5);
        tree.add(15);
        tree.add(3);

        assertEquals(10, tree.search(10));
        assertEquals(5, tree.search(5));
    }

    @Test
    @DisplayName("Should return null when searching for non-existing value")
    void testSearchNonExistingValue() {
        tree.add(10);
        tree.add(5);
        tree.add(15);

        assertNull(tree.search(20));
        assertNull(tree.search(1));
    }

    @Test
    @DisplayName("Should return null when searching in empty tree")
    void testSearchInEmptyTree() {
        assertNull(tree.search(10));
    }

    @Test
    @DisplayName("Should clear the tree")
    void testClear() {
        tree.add(10);
        tree.add(5);
        tree.add(15);

        tree.clear();

        assertNull(tree.root());
        assertNull(tree.search(10));
    }

    @Test
    @DisplayName("Should work with String values")
    void testWithStrings() {
        BinaryTree<String> stringTree = new BinaryTree<>();

        stringTree.add("dog");
        stringTree.add("cat");
        stringTree.add("elephant");

        assertEquals("dog", stringTree.root());
        assertEquals("cat", stringTree.search("cat"));
        assertNull(stringTree.search("zebra"));
    }

    @Test
    @DisplayName("Should handle single node tree")
    void testSingleNodeTree() {
        tree.add(42);

        assertEquals(42, tree.root());
        assertEquals(42, tree.search(42));
        assertNull(tree.search(1));
    }

    @Test
    @DisplayName("Should add nodes in ascending order")
    void testAddAscendingOrder() {
        tree.add(1);
        tree.add(2);
        tree.add(3);
        tree.add(4);
        tree.add(5);

        assertEquals(1, tree.root());
        assertEquals(3, tree.search(3)); // Note: Due to implementation, this might not work as expected
    }

    @Test
    @DisplayName("Should add nodes in descending order")
    void testAddDescendingOrder() {
        tree.add(5);
        tree.add(4);
        tree.add(3);
        tree.add(2);
        tree.add(1);

        assertEquals(5, tree.root());
    }


    @Test
    @DisplayName("Remove should return null (not implemented)")
    void testRemoveNotImplemented() {
        tree.add(10);
        tree.add(5);

        assertEquals(5,tree.remove(5));
        assertNull(tree.search(5));
        assertEquals(10, tree.search(10));
    }

    // Remove method tests
    @Test
    @DisplayName("Should return null when removing from empty tree")
    void testRemoveFromEmptyTree() {
        assertNull(tree.remove(10));
    }

    @Test
    @DisplayName("Should return null when removing non-existent value")
    void testRemoveNonExistentValue() {
        tree.add(10);
        tree.add(5);
        tree.add(15);

        assertNull(tree.remove(20));
        assertEquals(10, tree.root());
    }

    @Test
    @DisplayName("Should remove leftmost node")
    void testRemoveLeafNodeLeft() {
        tree.add(10);
        tree.add(5);
        tree.add(15);
        tree.add(3);

        assertEquals(3,tree.remove(3));
        assertNull(tree.search(3));
        assertEquals(10, tree.root());
        assertEquals(5, tree.search(5));
    }

    @Test
    @DisplayName("Should remove rightmost leaf")
    void testRemoveLeafNodeRight() {
        tree.add(10);
        tree.add(5);
        tree.add(15);
        tree.add(20);

        assertEquals(20,tree.remove(20));
        assertNull(tree.search(20));
        assertEquals(10, tree.root());
        assertEquals(15, tree.search(15));
    }

    @Test
    @DisplayName("Should remove node with only left child")
    void testRemoveNodeWithLeftChildOnly() {
        tree.add(10);
        tree.add(5);
        tree.add(15);
        tree.add(3);

        assertEquals(5, tree.remove(5));
        assertNull(tree.search(5));
        assertEquals(3, tree.search(3));
        assertEquals(10, tree.root());
    }

    @Test
    @DisplayName("Should remove node with only right child")
    void testRemoveNodeWithRightChildOnly() {
        tree.add(10);
        tree.add(5);
        tree.add(15);
        tree.add(7);

        assertEquals(5, tree.remove(5));
        assertNull(tree.search(5));
        assertEquals(7, tree.search(7));
        assertEquals(10, tree.root());
    }

    @Test
    @DisplayName("Should remove node with two children")
    void testRemoveNodeWithTwoChildren() {
        tree.add(10);
        tree.add(5);
        tree.add(15);
        tree.add(3);
        tree.add(7);
        tree.add(12);
        tree.add(20);

        assertEquals(5, tree.remove(5));
        assertNull(tree.search(5));
        assertEquals(7, tree.search(7));
        assertEquals(3, tree.search(3));
        assertEquals(10, tree.root());
    }

    @Test
    @DisplayName("Should remove root node with no children")
    void testRemoveRootNodeNoChildren() {
        tree.add(10);

        assertEquals(10, tree.remove(10));
        assertNull(tree.root());
        assertNull(tree.search(10));
    }

    @Test
    @DisplayName("Should remove root node with only left child")
    void testRemoveRootNodeWithLeftChildOnly() {
        tree.add(10);
        tree.add(5);
        tree.add(3);

        assertEquals(10, tree.remove(10));
        assertEquals(5, tree.root());
        assertNull(tree.search(10));
        assertEquals(3, tree.search(3));
    }

    @Test
    @DisplayName("Should remove root node with only right child")
    void testRemoveRootNodeWithRightChildOnly() {
        tree.add(10);
        tree.add(15);
        tree.add(20);

        assertEquals(10, tree.remove(10));
        assertEquals(15, tree.root());
        assertNull(tree.search(10));
        assertEquals(20, tree.search(20));
    }

    @Test
    @DisplayName("Should remove root node with two children")
    void testRemoveRootNodeWithTwoChildren() {
        tree.add(10);
        tree.add(5);
        tree.add(15);
        tree.add(3);
        tree.add(7);
        tree.add(12);
        tree.add(20);

        assertEquals(10, tree.remove(10));
        assertEquals(12, tree.root());
        assertNull(tree.search(10));
        assertEquals(5, tree.search(5));
        assertEquals(15, tree.search(15));
    }

    @Test
    @DisplayName("Should remove multiple nodes sequentially")
    void testRemoveMultipleNodes() {
        tree.add(10);
        tree.add(5);
        tree.add(15);
        tree.add(3);
        tree.add(7);
        tree.add(12);
        tree.add(20);

        // Safe Remove nodes - with a null leaf
        tree.remove(3); // leftmost
        tree.remove(20); // rightmost
        tree.remove(5); // new leftmost

        assertNull(tree.search(3));
        assertNull(tree.search(20));
        assertNull(tree.search(5));
        assertEquals(10, tree.root());
        assertEquals(7, tree.search(7));
        assertEquals(15, tree.search(15));
    }

    @Test
    @DisplayName("Should remove all nodes one by one")
    void testRemoveAllNodes() {
        tree.add(10);
        tree.add(5);
        tree.add(15);

        tree.remove(5);
        tree.remove(15);
        tree.remove(10);

        assertNull(tree.root());
        assertNull(tree.search(10));
    }

    @Test
    @DisplayName("Should handle complex removal scenario")
    void testComplexRemovalScenario() {
        // Build a more complex tree
        tree.add(50);
        tree.add(30);
        tree.add(70);
        tree.add(20);
        tree.add(40);
        tree.add(60);
        tree.add(80);
        tree.add(10);
        tree.add(25);
        tree.add(65);

        // Remove a node with two children
        assertEquals(30, tree.remove(30));
        assertNull(tree.search(30));
        assertEquals(40, tree.search(40));

        // Verify tree structure is still valid
        assertEquals(50, tree.root());
        assertEquals(20, tree.search(20));
        assertEquals(70, tree.search(70));
    }

    @Test
    @DisplayName("Should iterate over elements in in-order traversal")
    void testIterator() {
        // Case 1: Empty tree
        assertFalse(tree.iterator().hasNext(), "Iterator should have no elements for an empty tree.");

        // Case 2: Single-node tree
        tree.add(42);
        var iterator = tree.iterator();
        assertTrue(iterator.hasNext(), "Iterator should have one element for a single-node tree.");
        assertEquals(42, iterator.next(), "Iterator should return the single element in the tree.");
        assertFalse(iterator.hasNext(), "Iterator should have no more elements.");

        // Case 3: Multi-node tree
        tree.add(10);
        tree.add(20);
        tree.add(5);
        tree.add(15);
        tree.add(25);

        java.util.List<Integer> result = new java.util.ArrayList<>();
        for (Integer value : tree) {
            result.add(value);
        }

        assertEquals(java.util.Arrays.asList(5, 10, 15, 20, 25, 42),
                result, "Iterator should return elements in sorted (in-order traversal) order.");
    }

}
