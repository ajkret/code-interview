package com.dersommer.stack;

/**
 * My implementation of a stack.
 */
public interface Stack<T> extends Iterable<T> {

    void push(T value);

    T pop();

    T peek();

    int size();

    boolean isEmpty();

}
