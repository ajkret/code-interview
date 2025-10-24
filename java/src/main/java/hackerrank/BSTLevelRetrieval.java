package hackerrank;

import java.util.*;

/**
 * BST with various level-based retrieval methods
 */
public class BSTLevelRetrieval {

    static class TreeNode {
        int val;
        TreeNode left, right;

        TreeNode(int val) {
            this.val = val;
        }
    }

    static class BST {
        private TreeNode root;

        // Insert method
        public void insert(int val) {
            root = insertRec(root, val);
        }

        private TreeNode insertRec(TreeNode root, int val) {
            if (root == null) {
                return new TreeNode(val);
            }

            if (val < root.val) {
                root.left = insertRec(root.left, val);
            } else if (val > root.val) {
                root.right = insertRec(root.right, val);
            }

            return root;
        }

        // METHOD 1: Level-Order Traversal (BFS) - All elements level by level
        public List<Integer> levelOrderTraversal() {
            List<Integer> result = new ArrayList<>();
            if (root == null) return result;

            Queue<TreeNode> queue = new LinkedList<>();
            queue.offer(root);

            while (!queue.isEmpty()) {
                TreeNode current = queue.poll();
                result.add(current.val);

                if (current.left != null) {
                    queue.offer(current.left);
                }
                if (current.right != null) {
                    queue.offer(current.right);
                }
            }

            return result;
        }

        // METHOD 2: Get elements at specific level (0-indexed)
        public List<Integer> getElementsAtLevel(int targetLevel) {
            List<Integer> result = new ArrayList<>();
            getElementsAtLevelRec(root, 0, targetLevel, result);
            return result;
        }

        private void getElementsAtLevelRec(TreeNode node, int currentLevel, int targetLevel, List<Integer> result) {
            if (node == null) return;

            if (currentLevel == targetLevel) {
                result.add(node.val);
                return;
            }

            getElementsAtLevelRec(node.left, currentLevel + 1, targetLevel, result);
            getElementsAtLevelRec(node.right, currentLevel + 1, targetLevel, result);
        }

        // METHOD 3: Get all elements grouped by level
        public List<List<Integer>> getElementsByLevel() {
            List<List<Integer>> result = new ArrayList<>();
            if (root == null) return result;

            Queue<TreeNode> queue = new LinkedList<>();
            queue.offer(root);

            while (!queue.isEmpty()) {
                int levelSize = queue.size();
                List<Integer> currentLevel = new ArrayList<>();

                for (int i = 0; i < levelSize; i++) {
                    TreeNode current = queue.poll();
                    currentLevel.add(current.val);

                    if (current.left != null) {
                        queue.offer(current.left);
                    }
                    if (current.right != null) {
                        queue.offer(current.right);
                    }
                }

                result.add(currentLevel);
            }

            return result;
        }

        // METHOD 4: Get elements from level X to level Y
        public List<Integer> getElementsInLevelRange(int startLevel, int endLevel) {
            List<Integer> result = new ArrayList<>();
            getElementsInRangeRec(root, 0, startLevel, endLevel, result);
            return result;
        }

        private void getElementsInRangeRec(TreeNode node, int currentLevel, int startLevel, int endLevel, List<Integer> result) {
            if (node == null || currentLevel > endLevel) return;

            if (currentLevel >= startLevel) {
                result.add(node.val);
            }

            getElementsInRangeRec(node.left, currentLevel + 1, startLevel, endLevel, result);
            getElementsInRangeRec(node.right, currentLevel + 1, startLevel, endLevel, result);
        }

        // METHOD 5: Get leaf nodes (bottom level elements)
        public List<Integer> getLeafNodes() {
            List<Integer> result = new ArrayList<>();
            getLeafNodesRec(root, result);
            return result;
        }

        private void getLeafNodesRec(TreeNode node, List<Integer> result) {
            if (node == null) return;

            if (node.left == null && node.right == null) {
                result.add(node.val);
                return;
            }

            getLeafNodesRec(node.left, result);
            getLeafNodesRec(node.right, result);
        }

        // METHOD 6: Get elements at maximum depth level
        public List<Integer> getDeepestLevel() {
            if (root == null) return new ArrayList<>();

            int maxDepth = getMaxDepth();
            return getElementsAtLevel(maxDepth - 1); // 0-indexed
        }

        // METHOD 7: Get height/max depth of tree
        public int getMaxDepth() {
            return getMaxDepthRec(root);
        }

        private int getMaxDepthRec(TreeNode node) {
            if (node == null) return 0;

            int leftDepth = getMaxDepthRec(node.left);
            int rightDepth = getMaxDepthRec(node.right);

            return Math.max(leftDepth, rightDepth) + 1;
        }

        // METHOD 8: Get elements by level with their positions
        public Map<Integer, List<Integer>> getLevelPositionMap() {
            Map<Integer, List<Integer>> levelMap = new HashMap<>();
            getLevelPositionMapRec(root, 0, levelMap);
            return levelMap;
        }

        private void getLevelPositionMapRec(TreeNode node, int level, Map<Integer, List<Integer>> levelMap) {
            if (node == null) return;

            levelMap.computeIfAbsent(level, k -> new ArrayList<>()).add(node.val);

            getLevelPositionMapRec(node.left, level + 1, levelMap);
            getLevelPositionMapRec(node.right, level + 1, levelMap);
        }

        // METHOD 9: Right view of tree (rightmost element at each level)
        public List<Integer> getRightView() {
            List<Integer> result = new ArrayList<>();
            getRightViewRec(root, 0, result);
            return result;
        }

        private void getRightViewRec(TreeNode node, int level, List<Integer> result) {
            if (node == null) return;

            if (level == result.size()) {
                result.add(node.val);
            }

            // Visit right first to get rightmost element
            getRightViewRec(node.right, level + 1, result);
            getRightViewRec(node.left, level + 1, result);
        }

        // METHOD 10: Left view of tree (leftmost element at each level)
        public List<Integer> getLeftView() {
            List<Integer> result = new ArrayList<>();
            getLeftViewRec(root, 0, result);
            return result;
        }

        private void getLeftViewRec(TreeNode node, int level, List<Integer> result) {
            if (node == null) return;

            if (level == result.size()) {
                result.add(node.val);
            }

            // Visit left first to get leftmost element
            getLeftViewRec(node.left, level + 1, result);
            getLeftViewRec(node.right, level + 1, result);
        }

        // Utility method to print tree structure
        public void printTree() {
            System.out.println("Tree Structure:");
            printTreeRec(root, "", false);
        }

        private void printTreeRec(TreeNode node, String prefix, boolean isLeft) {
            if (node != null) {
                System.out.println(prefix + (isLeft ? "├── " : "└── ") + node.val);
                if (node.left != null || node.right != null) {
                    if (node.left != null) {
                        printTreeRec(node.left, prefix + (isLeft ? "│   " : "    "), true);
                    }
                    if (node.right != null) {
                        printTreeRec(node.right, prefix + (isLeft ? "│   " : "    "), false);
                    }
                }
            }
        }
    }

    public static void main(String[] args) {
        System.out.println("=== BST LEVEL-BASED RETRIEVAL DEMO ===");

        BST bst = new BST();

        // Build a sample BST
        int[] values = {50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 90};

        System.out.println("Inserting values: " + Arrays.toString(values));
        for (int val : values) {
            bst.insert(val);
        }

        bst.printTree();

        System.out.println("\n=== DIFFERENT LEVEL RETRIEVAL METHODS ===");

        // Method 1: Level-order traversal
        System.out.println("1. Level-Order Traversal: " + bst.levelOrderTraversal());

        // Method 2: Elements at specific levels
        System.out.println("2. Elements at level 0 (root): " + bst.getElementsAtLevel(0));
        System.out.println("   Elements at level 1: " + bst.getElementsAtLevel(1));
        System.out.println("   Elements at level 2: " + bst.getElementsAtLevel(2));
        System.out.println("   Elements at level 3: " + bst.getElementsAtLevel(3));

        // Method 3: All elements grouped by level
        System.out.println("3. Elements grouped by level:");
        List<List<Integer>> levelGroups = bst.getElementsByLevel();
        for (int i = 0; i < levelGroups.size(); i++) {
            System.out.println("   Level " + i + ": " + levelGroups.get(i));
        }

        // Method 4: Elements in level range
        System.out.println("4. Elements from level 1 to 3: " + bst.getElementsInLevelRange(1, 3));

        // Method 5: Leaf nodes
        System.out.println("5. Leaf nodes: " + bst.getLeafNodes());

        // Method 6: Deepest level elements
        System.out.println("6. Deepest level elements: " + bst.getDeepestLevel());

        // Method 7: Tree height
        System.out.println("7. Tree height: " + bst.getMaxDepth());

        // Method 8: Level position map
        System.out.println("8. Level-Position Map: " + bst.getLevelPositionMap());

        // Method 9: Right view
        System.out.println("9. Right view (rightmost at each level): " + bst.getRightView());

        // Method 10: Left view
        System.out.println("10. Left view (leftmost at each level): " + bst.getLeftView());

        System.out.println("\n=== TIME COMPLEXITIES ===");
        System.out.println("• Level-order traversal: O(n)");
        System.out.println("• Elements at specific level: O(n) worst case");
        System.out.println("• Elements grouped by level: O(n)");
        System.out.println("• Leaf nodes: O(n)");
        System.out.println("• Tree views: O(n)");
        System.out.println("• Tree height: O(n) worst case, O(log n) balanced");

        System.out.println("\n=== SPACE COMPLEXITIES ===");
        System.out.println("• BFS methods: O(w) where w is maximum width");
        System.out.println("• DFS methods: O(h) where h is tree height");
        System.out.println("• Result storage: O(k) where k is number of elements returned");
    }
}