package com.dersommer.list;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import java.util.Iterator;
import java.util.NoSuchElementException;

import static org.junit.jupiter.api.Assertions.*;

@DisplayName("DoubleLinkedList Tests")
class DoubleLinkedListTest {

    private DoubleLinkedList<Integer> list;

    @BeforeEach
    void setUp() {
        list = new DoubleLinkedList<>();
    }

    @Test
    @DisplayName("New list should be empty")
    void testNewListIsEmpty() {
        assertTrue(list.isEmpty());
        assertEquals(0, list.size());
    }

    @Test
    @DisplayName("Add elements and verify size")
    void testAddElements() {
        list.add(1);
        assertFalse(list.isEmpty());
        assertEquals(1, list.size());

        list.add(2);
        list.add(3);
        assertEquals(3, list.size());
    }

    @Test
    @DisplayName("Get elements by index")
    void testGetByIndex() {
        list.add(10);
        list.add(20);
        list.add(30);

        assertEquals(10, list.get(0));
        assertEquals(20, list.get(1));
        assertEquals(30, list.get(2));
    }

    @Test
    @DisplayName("Get element from start")
    void testStart() {
        list.add(1);
        list.add(2);
        list.add(3);

        assertEquals(1, list.start());
    }

    @Test
    @DisplayName("Get element from end")
    void testEnd() {
        list.add(1);
        list.add(2);
        list.add(3);

        assertEquals(3, list.end());
    }

    @Test
    @DisplayName("Get from end by index")
    void testGetFromEnd() {
        list.add(10);
        list.add(20);
        list.add(30);

        assertEquals(30, list.getFromEnd(0));
        assertEquals(20, list.getFromEnd(1));
        assertEquals(10, list.getFromEnd(2));
    }

    @Test
    @DisplayName("Remove element by index")
    void testRemove() {
        list.add(1);
        list.add(2);
        list.add(3);

        list.remove(1); // Remove middle element
        assertEquals(2, list.size());
        assertEquals(1, list.get(0));
        assertEquals(3, list.get(1));
    }

    @Test
    @DisplayName("Remove first element")
    void testRemoveFirst() {
        list.add(1);
        list.add(2);
        list.add(3);

        list.remove(0);
        assertEquals(2, list.size());
        assertEquals(2, list.get(0));
        assertEquals(3, list.get(1));
    }

    @Test
    @DisplayName("Remove last element")
    void testRemoveLast() {
        list.add(1);
        list.add(2);
        list.add(3);

        list.remove(2);
        assertEquals(2, list.size());
        assertEquals(1, list.get(0));
        assertEquals(2, list.get(1));
    }

    @Test
    @DisplayName("Clear list")
    void testClear() {
        list.add(1);
        list.add(2);
        list.add(3);

        list.clear();
        assertTrue(list.isEmpty());
        assertEquals(0, list.size());
    }

    @Test
    @DisplayName("Iterator traverses all elements")
    void testIterator() {
        list.add(1);
        list.add(2);
        list.add(3);

        Iterator<Integer> iterator = list.iterator();
        assertTrue(iterator.hasNext());
        assertEquals(1, iterator.next());
        assertTrue(iterator.hasNext());
        assertEquals(2, iterator.next());
        assertTrue(iterator.hasNext());
        assertEquals(3, iterator.next());
        assertFalse(iterator.hasNext());
    }

    @Test
    @DisplayName("Iterator on empty list")
    void testIteratorOnEmptyList() {
        Iterator<Integer> iterator = list.iterator();
        assertFalse(iterator.hasNext());
    }

    @Test
    @DisplayName("Iterator throws exception when no more elements")
    void testIteratorNoSuchElement() {
        list.add(1);
        Iterator<Integer> iterator = list.iterator();
        iterator.next();

        assertThrows(NoSuchElementException.class, iterator::next);
    }

    @Test
    @DisplayName("Enhanced for loop works")
    void testEnhancedForLoop() {
        list.add(1);
        list.add(2);
        list.add(3);

        int sum = 0;
        for (Integer value : list) {
            sum += value;
        }
        assertEquals(6, sum);
    }

    @Test
    @DisplayName("Test with String type")
    void testWithStrings() {
        DoubleLinkedList<String> stringList = new DoubleLinkedList<>();
        stringList.add("Hello");
        stringList.add("World");

        assertEquals(2, stringList.size());
        assertEquals("Hello", stringList.get(0));
        assertEquals("World", stringList.get(1));
    }

    @Test
    @DisplayName("Remove throws exception for invalid index")
    void testRemoveInvalidIndex() {
        list.add(1);
        list.add(2);

        assertEquals(2, list.size());

        list.remove(-1);
        list.remove(2);

        assertEquals(2, list.size());
    }

    @Test
    @DisplayName("Add and remove single element")
    void testSingleElement() {
        list.add(42);
        assertEquals(1, list.size());
        assertEquals(42, list.get(0));

        list.remove(0);
        assertTrue(list.isEmpty());
        assertEquals(0, list.size());
    }

    @Test
    @DisplayName("Remove last element")
    void testRemoveFromEnd() {
        list.add(1);
        list.add(2);
        list.add(3);

        assertEquals(3, list.removeFromEnd());
        assertEquals(2, list.size());
        assertEquals(2,list.end());
    }

    @Test
    @DisplayName("RemoveFromEnd on empty list returns null")
    void testRemoveFromEndOnEmptyList() {
        assertNull(list.removeFromEnd());
        assertEquals(0, list.size());
        assertTrue(list.isEmpty());
    }

    @Test
    @DisplayName("RemoveFromEnd on single element list")
    void testRemoveFromEndSingleElement() {
        list.add(42);

        assertEquals(42, list.removeFromEnd());
        assertEquals(0, list.size());
        assertTrue(list.isEmpty());
        assertNull(list.start());
        assertNull(list.end());
    }

    @Test
    @DisplayName("RemoveFromEnd multiple times")
    void testRemoveFromEndMultipleTimes() {
        list.add(1);
        list.add(2);
        list.add(3);
        list.add(4);

        assertEquals(4, list.removeFromEnd());
        assertEquals(3, list.size());
        assertEquals(3, list.end());

        assertEquals(3, list.removeFromEnd());
        assertEquals(2, list.size());
        assertEquals(2, list.end());

        assertEquals(2, list.removeFromEnd());
        assertEquals(1, list.size());
        assertEquals(1, list.end());

        assertEquals(1, list.removeFromEnd());
        assertEquals(0, list.size());
        assertTrue(list.isEmpty());
    }

    @Test
    @DisplayName("RemoveFromEnd until empty then try again")
    void testRemoveFromEndUntilEmptyAndTryAgain() {
        list.add(1);
        list.add(2);

        assertEquals(2, list.removeFromEnd());
        assertEquals(1, list.removeFromEnd());
        assertTrue(list.isEmpty());

        // Try to remove from empty list
        assertNull(list.removeFromEnd());
        assertTrue(list.isEmpty());
    }

    @Test
    @DisplayName("RemoveFromEnd maintains head reference correctly")
    void testRemoveFromEndMaintainsHeadReference() {
        list.add(10);
        list.add(20);
        list.add(30);

        list.removeFromEnd();
        assertEquals(10, list.start());
        assertEquals(10, list.get(0));

        list.removeFromEnd();
        assertEquals(10, list.start());
        assertEquals(10, list.get(0));
    }

    @Test
    @DisplayName("RemoveFromEnd then add new elements")
    void testRemoveFromEndThenAdd() {
        list.add(1);
        list.add(2);

        assertEquals(2, list.removeFromEnd());

        list.add(3);
        list.add(4);

        assertEquals(3, list.size());
        assertEquals(1, list.start());
        assertEquals(4, list.end());
        assertEquals(1, list.get(0));
        assertEquals(3, list.get(1));
        assertEquals(4, list.get(2));
    }

    @Test
    @DisplayName("RemoveFromEnd with two elements")
    void testRemoveFromEndWithTwoElements() {
        list.add(100);
        list.add(200);

        assertEquals(200, list.removeFromEnd());
        assertEquals(1, list.size());
        assertEquals(100, list.start());
        assertEquals(100, list.end());
        assertFalse(list.isEmpty());
    }

    @Test
    @DisplayName("Interleave removeFromEnd with add operations")
    void testInterleaveRemoveFromEndAndAdd() {
        list.add(1);
        assertEquals(1, list.removeFromEnd());

        list.add(2);
        list.add(3);
        assertEquals(3, list.removeFromEnd());

        list.add(4);
        assertEquals(2, list.size());
        assertEquals(2, list.get(0));
        assertEquals(4, list.get(1));
    }

    @Test
    @DisplayName("RemoveFromEnd preserves iteration order")
    void testRemoveFromEndPreservesIteration() {
        list.add(1);
        list.add(2);
        list.add(3);
        list.add(4);

        list.removeFromEnd(); // Remove 4
        list.removeFromEnd(); // Remove 3

        Iterator<Integer> iterator = list.iterator();
        assertEquals(1, iterator.next());
        assertEquals(2, iterator.next());
        assertFalse(iterator.hasNext());
    }

    @Test
    @DisplayName("RemoveFromEnd with getFromEnd operations")
    void testRemoveFromEndWithGetFromEnd() {
        list.add(10);
        list.add(20);
        list.add(30);
        list.add(40);

        assertEquals(40, list.removeFromEnd());

        assertEquals(30, list.getFromEnd(0));
        assertEquals(20, list.getFromEnd(1));
        assertEquals(10, list.getFromEnd(2));
    }
}