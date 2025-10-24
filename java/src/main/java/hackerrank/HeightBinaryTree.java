package hackerrank;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.List;
import java.util.stream.IntStream;

import static java.util.stream.Collectors.toList;

/**
 * The height of a binary tree is the number of edges between the tree’s root and its furthest leaf.
 * For example, the following binary tree is of height 2:
 *
 * Complete the getHeight or height function in the editor. It must return the height of a binary tree as an integer.
 *
 * getHeight or height has the following parameter(s):
 *
 * root: a reference to the root of a binary tree.
 *
 * Note: The height of a binary tree with a single node is taken as zero.
 *
 * Input Format
 * The first line contains an integer n, the number of nodes in the tree.
 * The next line contains n space-separated integers, where the ith integer denotes node[i].data.
 *
 * Note: Node values are inserted into a binary search tree before a reference to the tree’s root node is passed to your function.
 * In a binary search tree, all nodes on the left branch of a node are less than the node value, and all nodes on the right branch are greater than the node value.
 */
public class HeightBinaryTree {
    public static int calculateHeight(int nodeIndex, int[] leftChild, int[] rightChild) {
        if (nodeIndex == -1) {
            return 0;
        }

        int leftHeight = calculateHeight(leftChild[nodeIndex], leftChild, rightChild);
        int rightHeight = calculateHeight(rightChild[nodeIndex], leftChild, rightChild);

        return 1 + Math.max(leftHeight, rightHeight);
    }

    public static int getBinarySearchTreeHeight(List<Integer> values, List<Integer> leftChild, List<Integer> rightChild) {
        int[] leftChildArray = leftChild.stream().mapToInt(Integer::intValue).toArray();
        int[] rightChildArray = rightChild.stream().mapToInt(Integer::intValue).toArray();
        return calculateHeight(0, leftChildArray, rightChildArray);
    }

    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));

        int valuesCount = Integer.parseInt(bufferedReader.readLine().trim());

        List<Integer> values = IntStream.range(0, valuesCount).mapToObj(i -> {
                    try {
                        return bufferedReader.readLine().replaceAll("\\s+$", "");
                    } catch (IOException ex) {
                        throw new RuntimeException(ex);
                    }
                })
                .map(String::trim)
                .map(Integer::parseInt)
                .collect(toList());

        int leftChildCount = Integer.parseInt(bufferedReader.readLine().trim());

        List<Integer> leftChild = IntStream.range(0, leftChildCount).mapToObj(i -> {
                    try {
                        return bufferedReader.readLine().replaceAll("\\s+$", "");
                    } catch (IOException ex) {
                        throw new RuntimeException(ex);
                    }
                })
                .map(String::trim)
                .map(Integer::parseInt)
                .collect(toList());

        int rightChildCount = Integer.parseInt(bufferedReader.readLine().trim());

        List<Integer> rightChild = IntStream.range(0, rightChildCount).mapToObj(i -> {
                    try {
                        return bufferedReader.readLine().replaceAll("\\s+$", "");
                    } catch (IOException ex) {
                        throw new RuntimeException(ex);
                    }
                })
                .map(String::trim)
                .map(Integer::parseInt)
                .collect(toList());

        int result = HeightBinaryTree.getBinarySearchTreeHeight(values, leftChild, rightChild);

        System.out.println(result);

        bufferedReader.close();
    }


}
