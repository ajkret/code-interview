package com.dersommer.tree;

/**
 * A generic interface representing the structure and operations of a tree.
 * <p>
 * The {@code Tree} interface defines fundamental methods for managing
 * hierarchical data structures, such as insertion, search, removal, and clearing
 * operations. All implementing classes must ensure that these operations adhere
 * to the tree's specific rules (e.g., ordering for a Binary Search Tree or balancing for AVL trees).
 * </p>
 *
 * <h2>Example Tree Structure</h2>
 * Below is an example of a generic binary tree:
 * <pre>
 *          10
 *         /  \
 *        5    15
 *       / \     \
 *      3   7     20
 * </pre>
 * In the tree example:
 * <ul>
 *     <li>The root value is 10.</li>
 *     <li>5 and 15 are the children of 10.</li>
 *     <li>3 and 7 are the left and right children of 5, respectively.</li>
 *     <li>20 is the right child of 15.</li>
 * </ul>
 *
 * <p>This interface uses a generic type parameter {@code V}, where {@code V}
 * represents the type of values stored in the tree. The type {@code V} must
 * implement {@link Comparable} to ensure elements can be ordered.</p>
 *
 * @param <V> the type of elements maintained by this tree, which must be comparable
 * @see Iterable
 */
public interface Tree<V extends Comparable<V>> extends Iterable<V> {

    /**
     * Retrieves the value stored at the root node of the tree, if it exists.
     *
     * @return the value of the root node, or {@code null} if the tree is empty
     */
    V root();

    /**
     * Inserts a value into the tree.
     * <p>
     * The rules for insertion (e.g., maintaining order or balance) are
     * determined by the specific implementation. In a Binary Search Tree,
     * for instance, values are inserted such that the tree remains ordered.
     * </p>
     *
     * @param value the value to insert
     * @throws NullPointerException if {@code value} is {@code null}
     */
    void add(V value);

    /**
     * Searches for a value in the tree.
     * <p>
     * Implementations should specify whether the search is case-sensitive (if applicable)
     * and how duplicate values, if any, are handled. If the value exists in the tree,
     * it will typically return the first occurrence found during traversal.
     * </p>
     *
     * @param key the value to search for
     * @return the value if found, or {@code null} if the value does not exist in the tree
     */
    V search(V key);

    /**
     * Removes a value from the tree.
     * <p>
     * Removal may impact the tree's structure. Implementations must ensure that the
     * tree maintains its properties (e.g., ordering for Binary Search Trees, balancing for AVL trees).
     * </p>
     *
     * @param value the value to remove
     * @return the removed value if it was successfully deleted, or {@code null} if the value was not found
     */
    V remove(V value);

    /**
     * Clears all elements from the tree.
     * <p>
     * After clearing, the tree will be empty and {@link #root()} will return {@code null}.
     * </p>
     */
    void clear();
}
