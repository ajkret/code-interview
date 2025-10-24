package hackerrank;

import org.junit.jupiter.api.Test;
import java.util.*;
import static org.junit.jupiter.api.Assertions.*;

public class LongestArithmeticProgressionTest {

    @Test
    void testSimpleIncreasingSequence() {
        // 1, 3, 5, 7, 9 → difference k = 2
        List<Integer> arr = Arrays.asList(1, 3, 5, 7, 9);
        int k = 2;

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        assertEquals(5, result, "Full sequence forms an arithmetic progression with k=2");
    }

    @Test
    void testSimpleDecreasingSequence() {
        // 9, 7, 5, 3, 1 → difference k = -2
        List<Integer> arr = Arrays.asList(9, 7, 5, 3, 1);
        int k = -2;

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        assertEquals(5, result, "Decreasing sequence should also work with k=-2");
    }

    @Test
    void testUnorderedElements() {
        // 5, 1, 9, 7, 3 → difference k = 2
        // Sorted set → 1,3,5,7,9 forms a chain of 5
        List<Integer> arr = Arrays.asList(5, 1, 9, 7, 3);
        int k = 2;

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        assertEquals(5, result, "Order of elements should not matter");
    }

    @Test
    void testMissingElementBreaksChain() {
        // 1, 3, 7, 9 with k = 2 → chains: (1,3) and (7,9)
        List<Integer> arr = Arrays.asList(1, 3, 7, 9);
        int k = 2;

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        assertEquals(2, result, "Longest valid subchain length should be 2");
    }

    @Test
    void testDuplicateValuesIgnored() {
        // Duplicates should not affect result
        // 1, 1, 3, 3, 5 → k=2 → (1,3,5)
        List<Integer> arr = Arrays.asList(1, 1, 3, 3, 5);
        int k = 2;

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        assertEquals(3, result, "Duplicates should not inflate progression length");
    }

    @Test
    void testSingleElement() {
        List<Integer> arr = Collections.singletonList(42);
        int k = 3;

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        assertEquals(1, result, "Single element should form a progression of length 1");
    }

    @Test
    void testEmptyArray() {
        List<Integer> arr = Collections.emptyList();
        int k = 5;

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        assertEquals(0, result, "Empty array should return 0");
    }

    @Test
    void testNullArray() {
        int result = LongestArithmeticProgression.findLongestArithmeticProgression(null, 5);

        assertEquals(0, result, "Null input should safely return 0");
    }

    @Test
    void testNegativeStepGaps() {
        // 10, 8, 6, 4, 2 with k = -2
        List<Integer> arr = Arrays.asList(10, 8, 6, 4, 2);
        int k = -2;

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        assertEquals(5, result, "Progression with negative step should be detected correctly");
    }

    @Test
    void testMultipleProgressionsPickLongest() {
        // 1, 3, 5, 10, 15, 20 with k = 5 → two chains:
        // (1,3,5) length 3 and (10,15,20) length 3 → longest = 3
        List<Integer> arr = Arrays.asList(1, 3, 5, 10, 15, 20);
        int k = 5;

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        assertEquals(3, result, "Should return the length of the longest arithmetic chain");
    }
}
