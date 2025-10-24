package com.dersommer.tree;

import com.dersommer.stack.DoubleLinkedListStack;
import com.dersommer.stack.Stack;

import java.util.Iterator;
import java.util.NoSuchElementException;

/**
 * A generic implementation of a binary tree that supports basic operations such as
 * adding, searching, removing, and iterating over its elements. This binary tree is
 * structured to maintain ordering based on the natural ordering of the elements
 * (a Comparable requirement).
 *
 * @param <V> the type of elements held in the tree, which must implement Comparable.
 */
public class BinaryTree<V extends Comparable<V>> implements Tree<V> {
    private Node<V> root;
    private int size;

    @Override
    public V root() {
        return root != null ? root.data : null;
    }

    @Override
    public void add(V value) {
        ++size;
        var newNode = new Node<>(value);
        if (root == null) {
            // Special case for the root
            root = newNode;
            return;
        }

        // Insertion - avoid recursivity
        var node = root;
        do {
            int direction = value.compareTo(node.data);
            if (direction == 0) {
                // Special case for duplicate values
                return;
            }
            if (direction < 0) {
                if (node.left == null) {
                    node.left = newNode;
                    return;
                }
                node = node.left;
            } else {
                if (node.right == null) {
                    node.right = newNode;
                    return;
                }
                node = node.right;
            }
        } while (node != null);
    }

    @Override
    public V search(V value) {
        var node = root;
        do {
            if (node == null)
                break;
            int direction = value.compareTo(node.data);
            if (direction == 0) {
                // Found
                return node.data;
            }
            if (direction < 0) {
                if (node.left == null) {
                    break;
                }
                node = node.left;
            } else {
                if (node.right == null) {
                    break;
                }
                node = node.right;
            }
        } while (node != null);
        return null;
    }

    @Override
    public V remove(V value) {
        if (root == null)
            return null;

        // Find the node to remove - loop while there is a child or find a match
        Node<V> parent = null;
        Node<V> node = root;

        boolean cameFromLeft = false; // This will avoid a comparison later
        for (int direction = 0; node != null && (direction = value.compareTo(node.data)) != 0; ) {
            parent = node;
            if (direction < 0) {
                cameFromLeft = true;
                node = node.left;
            } else {
                cameFromLeft = false;
                node = node.right;
            }
        }
        if (node == null) {
            // Did not find a match
            return null;
        }

        // Cases where one leaf is null
        Node<V> child = null;
        if (node.left == null || node.right == null) {

            child = node.left != null ? node.left : node.right;

            if (parent == null) {
                // Deleting the root node - replace it with child since one leaf is null
                root = child;
            } else {
                // Identifies which leaf to bind with parent
                if (cameFromLeft) {
                    // Came from left
                    parent.left = child;
                } else {
                    parent.right = child;
                }
            }
        } else {

            // Now the challenging case, find a successor to replace the node to remove
            // It will go to the right leaf and keep searching until it finds a leaf on the left
            Node<V> successor = node.right;
            parent = node;
            while (successor.left != null) {
                parent = successor;
                successor = successor.left;
            }

            // Replace the value, not the node itself; the node can be root
            node.data = successor.data;

            // Now delete successor
            if (parent.left != null && parent.left.data.compareTo(successor.data) == 0) {
                parent.left = successor.right;
            } else {
                parent.right = successor.right;
            }
        }

        // Return a valid value to identify removal
        return value;
    }

    @Override
    public void clear() {
        root = null;
        size = 0;
    }

    @Override
    public Iterator<V> iterator() {
        return new InOrderIterator();
    }

    private static class Node<V> {
        V data;
        Node<V> left;
        Node<V> right;

        Node(V data) {
            this.data = data;
            this.left = null;
            this.right = null;
        }
    }

    /**
     * In-order iterator implementation (DLR).
     * It is possible to have a reverse order by having an implementation that pushes
     * to the stack from the right
     */
    private final class InOrderIterator implements Iterator<V> {
        private final Stack<Node<V>> stack = new DoubleLinkedListStack<>();

        InOrderIterator() {
            pushLeft(root);
        }

        /**
         * Push all nodes left side of the given node to the stack
         * @param node root node to start from
         */
        private void pushLeft(Node<V> node) {
            while (node != null) {
                stack.push(node);
                node = node.left;
            }
        }

        @Override
        public boolean hasNext() {
            return !stack.isEmpty();
        }

        /**
         * Take this tree
         *            10
         *         /     \
         *        5       15
         *       / \     |  \
         *      3   7   12  20
         *
         *  When the class initializes, it pushes the root node to the stack,
         *  then proceeds to push all nodes left side.
         *  The stack will look like this:
         *
         *  10,5,3
         *
         *  Then it pops the last node *3*, saves it in memory (this point is very important)
         *  and pushes start to push all nodes in the right subtree.
         *
         *  In our example, 3 has no right leaf. the stack will look like this:
         *
         *  10,5
         *
         *  In the following call to next(), we will pop the last node *5*,
         *  then push the right subtree. The stack will look like this:
         *
         *  10,7
         *
         *  Next calls:
         *
         *  10
         *  15,12
         *  15
         *  20
         */
        @Override
        public V next() {
            if (stack.isEmpty()) {
                // Follow the classic iterator contract
                throw new NoSuchElementException();
            }

            Node<V> n = stack.pop();
            if (n.right != null) {
                pushLeft(n.right);
            }
            return n.data;
        }
    }
}
