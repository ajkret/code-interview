package com.dersommer.list;

import java.util.Iterator;
import java.util.NoSuchElementException;

public class DoubleLinkedList<T> implements List<T> {
    private Node<T> head;
    private Node<T> tail;
    private int size;

    /**
     * Adds an element to the end of the list.
     * <pre>
     * Before add(value):
     *
     *   head         tail
     *    ↓             ↓
     *   [A] ←→ [B] ←→ [C]
     *
     * After add(D):
     *
     *   head                 tail
     *    ↓                    ↓
     *   [A] ←→ [B] ←→ [C] ←→ [D]
     *
     * Special case - empty list:
     * Before: head = null, tail = null
     * After:  head → [A] ← tail
     *
     * </pre>
     *
     * @param value the element to be added
     */
    @Override
    public void add(T value) {
        if (value == null)
            return;

        Node<T> newNode = new Node<>(value);
        if (head == null) {
            head = newNode;
            tail = newNode;
        } else {
            newNode.prior = tail;
            tail.next = newNode;
            tail = newNode;
        }
        // Keep track of the size, otherwise when calling size()
        // we will have to traverse the entire list
        size++;
    }

    @Override
    public T get(int index) {
        if (index < 0 || index >= size)
            return null;

        Node<T> node = head;
        for (int i = 0; i < index; i++) {
            node = node.next;
        }
        return node.data;
    }

    @Override
    public T getFromEnd(int index) {
        if (index < 0 || index >= size)
            return null;

        Node<T> node = tail;
        for (int i = 0; i < index; i++) {
            node = node.prior;
        }
        return node.data;
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public void clear() {
        head = null;
        tail = null;
        size = 0;
    }

    @Override
    public boolean isEmpty() {
        if (head == null)
            return true;
        return false;
    }

    /**
     * Removes the node after the current node by updating the links.
     *
     * <pre>
     * Before:
     *
     *          node       toRemove     nextNode
     *           ↓            ↓            ↓
     *   ... ←→ [A]    ←→    [B]    ←→    [C]    ←→ ...
     *
     * Step 1: node.next = node.next.next;
     *
     *          node             nextNode
     *           ↓                  ↓
     *   ... ←→ [A] ─────────────→ [C] ←→ ...
     *              ←─ [B] ←──────┘
     *              (orphaned)
     *
     * Step 2: node.next.prior = node;
     *
     *          node           nextNode
     *           ↓                ↓
     *   ... ←→ [A] ←──────────→ [C] ←→ ...
     *              [B] (removed, no references)
     *
     * After:
     *
     *        node       nextNode
     *         ↓            ↓
     *   ... ←→ [A]   ←→   [C] ←→ ...
     *
     * Node [B] is now unreachable and will be garbage collected.
     * </pre>
     */
    @Override
    public void remove(int index) {
        if (index < 0 || index >= size)
            return;
        // Special case - head
        if (index == 0) {
            head = head.next;
        } else {
            Node<T> node = head;
            for (int i = 0; i < index - 1; i++) {
                node = node.next;
            }

            node.next.prior = node;
            node.next = node.next.next;
        }
        size--;
    }

    @Override
    public T removeFromEnd() {
        if(tail==null)
            return null;
        T data = tail.data;
        tail = tail.prior;
        if(tail!=null)
            tail.next = null;
        else
            head = null;
        size--;
        return data;
    }

    @Override
    public T end() {
        if(isEmpty())
            return null;
        return tail.data;
    }

    @Override
    public T start() {
        if(isEmpty())
            return null;
        return head.data;
    }

    @Override
    public Iterator<T> iterator() {
        // TODO: implement iterator from the last element to the first element
        return new DoubleLinkedListIterator();
    }

    /**
     * Iterator implementation for traversing the doubly linked list from head to tail.
     */
    private class DoubleLinkedListIterator implements Iterator<T> {
        private Node<T> current;

        /**
         * Initializes the iterator at the head of the list.
         */
        public DoubleLinkedListIterator() {
            this.current = head;
        }

        /**
         * Checks if there are more elements to iterate over.
         *
         * @return {@code true} if the iteration has more elements
         */
        @Override
        public boolean hasNext() {
            return current != null;
        }

        /**
         * Returns the next element in the iteration.
         * <pre>
         * Iteration flow:
         *
         *   head                tail
         *    ↓                    ↓
         *   [A] ←→ [B] ←→ [C] ←→ [D]
         *    ↑
         *  current (start)
         *
         * After next(): returns A
         *           ↓
         *   [A] ←→ [B] ←→ [C] ←→ [D]
         *           ↑
         *        current
         *
         * After next(): returns B
         *                  ↓
         *   [A] ←→ [B] ←→ [C] ←→ [D]
         *                  ↑
         *               current
         * </pre>
         *
         * @return the next element in the iteration
         * @throws NoSuchElementException if the iteration has no more elements
         */
        @Override
        public T next() {
            if (!hasNext()) {
                throw new NoSuchElementException("No more elements in the list");
            }
            T data = current.data;
            current = current.next;
            return data;
        }
    }

    private static class Node<T> {
        T data;
        Node<T> next;
        Node<T> prior;

        Node(T data) {
            this.data = data;
            this.next = null;
            this.prior = null;
        }
    }
}


