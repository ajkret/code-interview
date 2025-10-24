package com.dersommer.stack;

import com.dersommer.list.DoubleLinkedList;

import java.util.Iterator;

public class DoubleLinkedListStack<T> implements Stack<T>{
    private DoubleLinkedList<T> list = new DoubleLinkedList<>();

    @Override
    public void push(T value) {
        if (value == null)
            return;

        list.add(value);
    }

    @Override
    public T pop() {
        if(list.isEmpty())
            return null;
        return list.removeFromEnd();
    }

    @Override
    public T peek() {
        if(!list.isEmpty())
            return list.end();
        return null;
    }

    @Override
    public int size() {
        return list.size();
    }

    @Override
    public boolean isEmpty() {
        return list.isEmpty();
    }

    @Override
    public Iterator<T> iterator() {
        return list.iterator();
    }
}
