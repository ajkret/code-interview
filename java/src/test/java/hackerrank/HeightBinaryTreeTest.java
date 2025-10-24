package hackerrank;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;
import java.util.List;

public class HeightBinaryTreeTest {

    /**
     * Helper function to simulate the height calculation by calling
     * the same private recursive logic indirectly.
     * Since calculateHeight is private, we assume a public
     * wrapper like getBinarySearchTreeHeight exists (as in main()).
     */
    private int getHeight(int[] values, int[] left, int[] right) {
        // Access HeightBinaryTree.getBinarySearchTreeHeight() directly
        // or simulate if not implemented yet.
        return HeightBinaryTree.getBinarySearchTreeHeight(
                List.of(values[0], values[1], values[2], values[3], values[4], values[5], values[6]),
                List.of(left[0], left[1], left[2], left[3], left[4], left[5], left[6]),
                List.of(right[0], right[1], right[2], right[3], right[4], right[5], right[6])
        );
    }

    @Test
    void testBalancedTree() {
        // Example tree from the problem statement:
        //        4
        //       / \
        //      2   6
        //     / \ / \
        //    1  3 5  7
        int[] values = {4, 2, 6, 1, 3, 5, 7};
        int[] left   = {1, 3, 5, -1, -1, -1, -1};
        int[] right  = {2, 4, 6, -1, -1, -1, -1};

        int height = HeightBinaryTree.getBinarySearchTreeHeight(
                List.of(4, 2, 6, 1, 3, 5, 7),
                List.of(1, 3, 5, -1, -1, -1, -1),
                List.of(2, 4, 6, -1, -1, -1, -1)
        );

        assertEquals(2, height, "Balanced tree should have height 2");
    }

    @Test
    void testSingleNodeTree() {
        int[] values = {1};
        int[] left   = {-1};
        int[] right  = {-1};

        int height = HeightBinaryTree.getBinarySearchTreeHeight(
                List.of(1),
                List.of(-1),
                List.of(-1)
        );

        assertEquals(0, height, "Single node tree should have height 0");
    }

    @Test
    void testLeftSkewedTree() {
        // A left-only chain (height = n-1)
        //     5
        //    /
        //   4
        //  /
        // 3
        // ...
        int[] values = {5, 4, 3};
        int[] left   = {1, 2, -1};
        int[] right  = {-1, -1, -1};

        int height = HeightBinaryTree.getBinarySearchTreeHeight(
                List.of(5, 4, 3),
                List.of(1, 2, -1),
                List.of(-1, -1, -1)
        );

        assertEquals(2, height, "Left skewed tree with 3 nodes should have height 2");
    }

    @Test
    void testRightSkewedTree() {
        // A right-only chain (height = n-1)
        // 1
        //  \
        //   2
        //    \
        //     3
        int[] values = {1, 2, 3};
        int[] left   = {-1, -1, -1};
        int[] right  = {1, 2, -1};

        int height = HeightBinaryTree.getBinarySearchTreeHeight(
                List.of(1, 2, 3),
                List.of(-1, -1, -1),
                List.of(1, 2, -1)
        );

        assertEquals(2, height, "Right skewed tree with 3 nodes should have height 2");
    }

    @Test
    void testUnbalancedTree() {
        //       10
        //      /  \
        //     5    15
        //    /
        //   3
        //  /
        // 2
        int[] values = {10, 5, 15, 3, 2};
        int[] left   = {1, 3, -1, 4, -1};
        int[] right  = {2, -1, -1, -1, -1};

        int height = HeightBinaryTree.getBinarySearchTreeHeight(
                List.of(10, 5, 15, 3, 2),
                List.of(1, 3, -1, 4, -1),
                List.of(2, -1, -1, -1, -1)
        );

        assertEquals(3, height, "Unbalanced tree should have height 3");
    }
}
