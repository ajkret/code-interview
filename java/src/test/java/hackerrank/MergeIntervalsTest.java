package hackerrank;

import org.junit.jupiter.api.Test;
import java.util.*;
import static org.junit.jupiter.api.Assertions.*;

public class MergeIntervalsTest {

    @Test
    void testOverlappingIntervals() {
        // Example: [[1,3],[2,6],[8,10],[15,18]]
        // Expected: [[1,6],[8,10],[15,18]]
        List<List<Integer>> intervals = Arrays.asList(
                Arrays.asList(1, 3),
                Arrays.asList(2, 6),
                Arrays.asList(8, 10),
                Arrays.asList(15, 18)
        );

        List<List<Integer>> expected = Arrays.asList(
                Arrays.asList(1, 6),
                Arrays.asList(8, 10),
                Arrays.asList(15, 18)
        );

        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);
        assertEquals(expected, result);
    }

    @Test
    void testTouchingIntervalsShouldMerge() {
        // [[1,2],[2,3],[3,4]] -> [[1,4]]
        List<List<Integer>> intervals = Arrays.asList(
                Arrays.asList(1, 2),
                Arrays.asList(2, 3),
                Arrays.asList(3, 4)
        );

        List<List<Integer>> expected = Collections.singletonList(Arrays.asList(1, 4));
        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);
        assertEquals(expected, result);
    }

    @Test
    void testNonOverlappingIntervals() {
        // [[1,2],[4,5],[7,8]] -> same result (no merges)
        List<List<Integer>> intervals = Arrays.asList(
                Arrays.asList(1, 2),
                Arrays.asList(4, 5),
                Arrays.asList(7, 8)
        );

        List<List<Integer>> expected = Arrays.asList(
                Arrays.asList(1, 2),
                Arrays.asList(4, 5),
                Arrays.asList(7, 8)
        );

        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);
        assertEquals(expected, result);
    }

    @Test
    void testUnorderedIntervals() {
        // [[8,10],[1,3],[2,6]] -> same as [[1,3],[2,6],[8,10]] => [[1,6],[8,10]]
        List<List<Integer>> intervals = Arrays.asList(
                Arrays.asList(8, 10),
                Arrays.asList(1, 3),
                Arrays.asList(2, 6)
        );

        List<List<Integer>> expected = Arrays.asList(
                Arrays.asList(1, 6),
                Arrays.asList(8, 10)
        );

        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);
        assertEquals(expected, result);
    }

    @Test
    void testSingleInterval() {
        // [[5,7]] -> [[5,7]]
        List<List<Integer>> intervals = Collections.singletonList(Arrays.asList(5, 7));
        List<List<Integer>> expected = Collections.singletonList(Arrays.asList(5, 7));

        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);
        assertEquals(expected, result);
    }

    @Test
    void testEmptyList() {
        List<List<Integer>> intervals = Collections.emptyList();

        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);

        assertTrue(result.isEmpty(), "Empty list should produce empty result");
    }

    @Test
    void testNestedIntervals() {
        // [[1,10],[2,5],[3,4]] -> [[1,10]]
        List<List<Integer>> intervals = Arrays.asList(
                Arrays.asList(1, 10),
                Arrays.asList(2, 5),
                Arrays.asList(3, 4)
        );

        List<List<Integer>> expected = Collections.singletonList(Arrays.asList(1, 10));
        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);

        assertEquals(expected, result);
    }

    @Test
    void testIntervalsWithSameStartDifferentEnds() {
        // [[1,4],[1,5],[1,3]] -> [[1,5]]
        List<List<Integer>> intervals = Arrays.asList(
                Arrays.asList(1, 4),
                Arrays.asList(1, 5),
                Arrays.asList(1, 3)
        );

        List<List<Integer>> expected = Collections.singletonList(Arrays.asList(1, 5));
        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);

        assertEquals(expected, result);
    }

    @Test
    void testIntervalsWithSameEndDifferentStarts() {
        // [[1,5],[2,5],[3,5]] -> [[1,5]]
        List<List<Integer>> intervals = Arrays.asList(
                Arrays.asList(1, 5),
                Arrays.asList(2, 5),
                Arrays.asList(3, 5)
        );

        List<List<Integer>> expected = Collections.singletonList(Arrays.asList(1, 5));
        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);

        assertEquals(expected, result);
    }
}
