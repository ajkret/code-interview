package hackerrank;

import java.util.*;
import java.lang.*;

public class InterviewHashMap<K extends Comparable<K>, V> {
    private static final int INITIAL_CAPACITY = 16;

    public Bucket<K, V>[] getBuckets() {
        return buckets;
    }

    private Bucket<K, V>[] buckets;

    public InterviewHashMap() {
        this(INITIAL_CAPACITY);
    }

    public InterviewHashMap(int capacity) {
        buckets = new Bucket[capacity];
        for (int i = 0; i < capacity; i++) {
            buckets[i] = new Bucket<>();
        }
    }

    private int getBucketIndex(K key) {
        if (key == null) return 0;
        else return Math.abs(key.hashCode()) % buckets.length;
    }

    public void put(K key, V value) {
        int index = getBucketIndex(key);
        Bucket<K, V> bucket = buckets[index];

        bucket.put(key, value);
    }

    public V get(K key) {
        int index = getBucketIndex(key);
        return buckets[index].get(key);
    }

    public void remove(K key) {
        int index = getBucketIndex(key);
        buckets[index].remove(key);
    }

    // Nested Bucket class
    static class Bucket<K extends Comparable<K>, V> {

        private static final int THRESHOLD = 8; // Threshold to transform to a binary tree

        private LinkedList<Entry<K, V>> entries = new LinkedList<>();
        private TreeNode<K, V> treeRoot = null;

        public boolean isTree() {
            return isTree;
        }

        private boolean isTree = false;

        public void put(K key, V value) {
            if(isTree) {
                treePut(treeRoot, key, value);
            } else {
                listPut(key, value);
                if(size()>THRESHOLD) {
                    transformToTree();
                }
            }
        }

        public V get(K key) {
            if(isTree) {
                return treeGet(treeRoot, key);
            } else {
                return listGet(key);
            }
        }

        private void listPut(K key, V value) {
            for (Entry<K, V> entry : entries) {
                if ((key == null && entry.key == null) || (key != null && key.equals(entry.key))) {
                    entry.value = value;
                    return;
                }
            }
            entries.add(new Entry<>(key, value));
        }

        private V listGet(K key) {
            for (Entry<K, V> entry : entries) {
                if ((key == null && entry.key == null) || (key != null && key.equals(entry.key))) {
                    return entry.value;
                }
            }
            return null;
        }

        public void remove(K key) {
            if(isTree) {
                treeRoot = treeRemove(treeRoot, key);
            } else {
                listRemove(key);
            }
        }

        private void listRemove(K key) {
            entries.removeIf(entry -> (key == null && entry.key == null) || (key != null && key.equals(entry.key)));
        }

        public int size() {
            return isTree ? countTreeNodes(treeRoot) : entries.size();
        }

        private void transformToTree() {
            treeRoot = null;
            for (Entry<K, V> entry : entries) {
                treeRoot = treePut(treeRoot, entry.key, entry.value);
            }
            entries = null;
            isTree = true;
        }

        // Binary Tree Operations
        private TreeNode<K, V> treePut(TreeNode<K, V> root, K key, V value) {
            if (root == null) {
                return new TreeNode<>(key, value);
            }
            int cmp = key.compareTo(root.key);
            if (cmp < 0) {
                root.left = treePut(root.left, key, value);
            } else if (cmp > 0) {
                root.right = treePut(root.right, key, value);
            } else {
                root.value = value;
            }
            return root;
        }

        private V treeGet(TreeNode<K, V> root, K key) {
            if (root == null) {
                return null;
            }
            int cmp = key.compareTo(root.key);
            if (cmp < 0) {
                return treeGet(root.left, key);
            } else if (cmp > 0) {
                return treeGet(root.right, key);
            } else {
                return root.value;
            }
        }

        private TreeNode<K, V> treeRemove(TreeNode<K, V> root, K key) {
            if (root == null) {
                return null;
            }
            int cmp = key.compareTo(root.key);
            if (cmp < 0) {
                root.left = treeRemove(root.left, key);
            } else if (cmp > 0) {
                root.right = treeRemove(root.right, key);
            } else {
                if (root.left == null) return root.right;
                if (root.right == null) return root.left;

                TreeNode<K, V> successor = findMin(root.right);
                root.key = successor.key;
                root.value = successor.value;
                root.right = treeRemove(root.right, successor.key);
            }

            return root;
        }

        private TreeNode<K, V> findMin(TreeNode<K, V> root) {
            while (root.left != null) {
                root = root.left;
            }
            return root;
        }

        private int countTreeNodes(TreeNode<K, V> root) {
            if (root == null) {
                return 0;
            }
            return 1 + countTreeNodes(root.left) + countTreeNodes(root.right);
        }
    }

    // Entry and TreeNode classes
    private static class Entry<K, V> {
        K key;
        V value;

        Entry(K key, V value) {
            this.key = key;
            this.value = value;
        }
    }

    private static class TreeNode<K, V> {
        K key;
        V value;
        TreeNode<K, V> left;
        TreeNode<K, V> right;

        TreeNode(K key, V value) {
            this.key = key;
            this.value = value;
        }
    }

    public static void main (String[] args) throws java.lang.Exception {
        InterviewHashMap map = new InterviewHashMap<Integer, String>();


        for(int i=0;i<20000;i++) {
            map.put(i,String.format("Value %s",i));
        }

        map.remove(12);

        System.out.println(map.get(12));

    }
}