package com.dersommer.list;

import java.util.Iterator;

/**
 * A generic list interface that defines the contract for a linear collection of elements.
 * <p>
 * This interface provides basic list operations such as adding, retrieving, and removing elements,
 * as well as navigation and iteration capabilities.
 * </p>
 *
 * @param <T> the type of elements in this list
 */
public interface List<T> extends Iterable<T> {

    /**
     * Adds an element to the list.
     *
     * @param value the element to be added
     */
    void add(T value);

    /**
     * Retrieves the element at the specified position in the list.
     *
     * @param index the position of the element to retrieve
     * @return the element at the specified index
     */
    T get(int index);

    /**
     * Retrieves the element at the specified position in the list,
     * starting from the end of the list.
     *
     * @param index the position of the element to retrieve
     * @return the element at the specified index
     */
    T getFromEnd(int index);

    /**
     * Returns the number of elements in the list.
     *
     * @return the size of the list
     */
    int size();

    /**
     * Removes all elements from the list.
     */
    void clear();

    /**
     * Checks if the list contains no elements.
     *
     * @return {@code true} if the list is empty, {@code false} otherwise
     */
    boolean isEmpty();

    /**
     * Removes the element at the specified position in the list.
     *
     * @param index the position of the element to remove
     */
    void remove(int index);

    /**
     * Removes the element at the specified position in the list. Special case, that will enable stack's pop()
     *
     * @param index the position of the element to remove
     * @return the removed element
     */
    T removeFromEnd();

    /**
     * Retrieve the element at the end of the list.
     */
    T end();

    /**
     * Retrieve the element at the start of the list.
     */
    T start();

    /**
     * Returns an iterator over elements of type {@code T}.
     *
     * @return an Iterator
     */
    Iterator<T> iterator();


}