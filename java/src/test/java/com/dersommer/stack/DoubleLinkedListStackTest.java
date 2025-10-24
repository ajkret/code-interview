package com.dersommer.stack;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.DisplayName;

import static org.junit.jupiter.api.Assertions.*;

class DoubleLinkedListStackTest {

    private DoubleLinkedListStack<Integer> stack;

    @BeforeEach
    void setUp() {
        stack = new DoubleLinkedListStack<>();
    }

    @Test
    @DisplayName("New stack should be empty")
    void testNewStackIsEmpty() {
        assertTrue(stack.isEmpty());
        assertEquals(0, stack.size());
    }

    @Test
    @DisplayName("Push single element")
    void testPushSingleElement() {
        stack.push(1);
        assertFalse(stack.isEmpty());
        assertEquals(1, stack.size());
    }

    @Test
    @DisplayName("Push multiple elements")
    void testPushMultipleElements() {
        stack.push(1);
        stack.push(2);
        stack.push(3);

        assertFalse(stack.isEmpty());
        assertEquals(3, stack.size());
    }

    @Test
    @DisplayName("Push null should not add element")
    void testPushNull() {
        stack.push(1);
        stack.push(null);
        stack.push(2);

        assertEquals(2, stack.size());
    }

    @Test
    @DisplayName("Peek returns last pushed element without removing it")
    void testPeek() {
        stack.push(1);
        stack.push(2);
        stack.push(3);

        assertEquals(3, stack.peek());
        assertEquals(3, stack.size()); // Size should not change
        assertEquals(3, stack.peek()); // Should still return same value
    }

    @Test
    @DisplayName("Peek on empty stack returns null")
    void testPeekEmptyStack() {
        assertNull(stack.peek());
    }

    @Test
    @DisplayName("Pop returns and removes last pushed element")
    void testPop() {
        stack.push(1);
        stack.push(2);
        stack.push(3);

        assertEquals(3, stack.pop());
        assertEquals(2, stack.size());
        assertEquals(2, stack.peek());
    }

    @Test
    @DisplayName("Pop on empty stack returns null")
    void testPopEmptyStack() {
        assertNull(stack.pop());
        assertTrue(stack.isEmpty());
    }

    @Test
    @DisplayName("LIFO order - Last In First Out")
    void testLIFOOrder() {
        stack.push(1);
        stack.push(2);
        stack.push(3);
        stack.push(4);

        assertEquals(4, stack.pop());
        assertEquals(3, stack.pop());
        assertEquals(2, stack.pop());
        assertEquals(1, stack.pop());
        assertTrue(stack.isEmpty());
    }

    @Test
    @DisplayName("Pop all elements makes stack empty")
    void testPopAllElements() {
        stack.push(1);
        stack.push(2);
        stack.push(3);

        stack.pop();
        stack.pop();
        stack.pop();

        assertTrue(stack.isEmpty());
        assertEquals(0, stack.size());
        assertNull(stack.peek());
        assertNull(stack.pop());
    }

    @Test
    @DisplayName("Push after pop")
    void testPushAfterPop() {
        stack.push(1);
        stack.push(2);
        assertEquals(2, stack.pop());

        stack.push(3);
        assertEquals(3, stack.peek());
        assertEquals(2, stack.size());
    }

    @Test
    @DisplayName("Interleaved push and pop operations")
    void testInterleavedOperations() {
        stack.push(1);
        assertEquals(1, stack.pop());

        stack.push(2);
        stack.push(3);
        assertEquals(3, stack.pop());

        stack.push(4);
        assertEquals(4, stack.peek());
        assertEquals(2, stack.size());

        assertEquals(4, stack.pop());
        assertEquals(2, stack.pop());
        assertTrue(stack.isEmpty());
    }

    @Test
    @DisplayName("Stack with String type")
    void testStackWithStrings() {
        DoubleLinkedListStack<String> stringStack = new DoubleLinkedListStack<>();

        stringStack.push("first");
        stringStack.push("second");
        stringStack.push("third");

        assertEquals("third", stringStack.pop());
        assertEquals("second", stringStack.pop());
        assertEquals("first", stringStack.peek());
        assertEquals(1, stringStack.size());
    }

    @Test
    @DisplayName("Large number of operations")
    void testLargeNumberOfOperations() {
        for (int i = 1; i <= 1000; i++) {
            stack.push(i);
        }

        assertEquals(1000, stack.size());
        assertEquals(1000, stack.peek());

        for (int i = 1000; i >= 1; i--) {
            assertEquals(i, stack.pop());
        }

        assertTrue(stack.isEmpty());
    }
}