package hackerrank;

import org.junit.jupiter.api.Test;
import java.util.*;
import static org.junit.jupiter.api.Assertions.*;

public class InterviewHashMapTest {

    @Test
    void testPutAndGetSimpleValues() {
        InterviewHashMap<Integer, String> map = new InterviewHashMap<>();

        map.put(1, "A");
        map.put(2, "B");

        assertEquals("A", map.get(1));
        assertEquals("B", map.get(2));
        assertNull(map.get(3), "Non-existent key should return null");
    }

    @Test
    void testUpdateExistingKey() {
        InterviewHashMap<String, String> map = new InterviewHashMap<>();

        map.put("key", "first");
        map.put("key", "updated");

        assertEquals("updated", map.get("key"), "Should replace value for same key");
    }

    @Test
    void testRemoveKey() {
        InterviewHashMap<Integer, String> map = new InterviewHashMap<>();

        map.put(1, "one");
        map.put(2, "two");
        map.remove(1);

        assertNull(map.get(1), "Removed key should return null");
        assertEquals("two", map.get(2), "Other keys should remain intact");
    }

    @Test
    void testNullKeyHandling() {
        InterviewHashMap<String, Integer> map = new InterviewHashMap<>();

        map.put(null, 99);
        map.put("key", 100);

        assertEquals(99, map.get(null));
        assertEquals(100, map.get("key"));
    }

    @Test
    void testManyInsertionsAcrossBuckets() {
        InterviewHashMap<Integer, String> map = new InterviewHashMap<>();

        for (int i = 0; i < 100; i++) {
            map.put(i, "v" + i);
        }

        assertEquals("v0", map.get(0));
        assertEquals("v50", map.get(50));
        assertEquals("v99", map.get(99));
    }

    @Test
    void testBucketTransformationToTree() throws Exception {
        // Force collisions by mocking keys that hash to the same bucket
        InterviewHashMap<BadKey, String> map = new InterviewHashMap<>(1); // one bucket → all collide

        for (int i = 0; i < 10; i++) {
            map.put(new BadKey(i), "v" + i);
        }

        // The bucket should have transformed into a tree (threshold = 8)
        InterviewHashMap.Bucket<BadKey, String> bucket = map.getBuckets()[0];
        assertTrue(bucket.isTree(), "Bucket should transform into tree after threshold exceeded");

        assertEquals("v5", map.get(new BadKey(5)));
        assertEquals("v9", map.get(new BadKey(9)));
    }

    @Test
    void testTreeNodeRemoval() {
        InterviewHashMap<BadKey, String> map = new InterviewHashMap<>(1);

        for (int i = 0; i < 10; i++) {
            map.put(new BadKey(i), "val" + i);
        }

        map.remove(new BadKey(5));
        assertNull(map.get(new BadKey(5)), "Removed node from tree should return null");
        assertEquals("val4", map.get(new BadKey(4)));
    }

    @Test
    void testLargeScaleInsertions() {
        InterviewHashMap<Integer, String> map = new InterviewHashMap<>();

        int size = 5000;
        for (int i = 0; i < size; i++) {
            map.put(i, "v" + i);
        }

        for (int i = 0; i < size; i += 500) {
            assertEquals("v" + i, map.get(i));
        }
    }

    @Test
    void testOverwriteValueInTreeMode() {
        InterviewHashMap<BadKey, String> map = new InterviewHashMap<>(1);

        for (int i = 0; i < 10; i++) {
            map.put(new BadKey(i), "v" + i);
        }

        map.put(new BadKey(5), "updated");
        assertEquals("updated", map.get(new BadKey(5)), "Existing node in tree should update value");
    }

    // Custom key class to force hash collisions
    static class BadKey implements Comparable<BadKey> {
        private final int value;

        BadKey(int value) {
            this.value = value;
        }

        @Override
        public int hashCode() {
            return 42; // Force all to same bucket
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (!(o instanceof BadKey)) return false;
            return this.value == ((BadKey) o).value;
        }

        @Override
        public int compareTo(BadKey other) {
            return Integer.compare(this.value, other.value);
        }

        @Override
        public String toString() {
            return "BadKey(" + value + ")";
        }
    }
}
