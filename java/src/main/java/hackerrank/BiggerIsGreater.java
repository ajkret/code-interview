package hackerrank;

import java.io.*;
import java.util.*;
import java.util.stream.*;

import static java.util.stream.Collectors.joining;

public class BiggerIsGreater {

    /**
     * Given a lowercase string {@code w}, this function finds the smallest lexicographically greater
     * permutation that can be formed by rearranging its characters.
     *
     * The algorithm follows the standard "next lexicographical permutation" approach:</br>
     * 1. Scan the string from right to left and find the first character (pivot) that is smaller
     *    than a character to its right.</br>
     * 2. Find the smallest character on the right side of the pivot that is greater than the pivot.</br>
     * 3. Swap them.</br>
     * 4. Sort (or reverse) the suffix after the pivot index to get the smallest possible greater word.</br>
     *
     * Example:
     *   w = "abcd" → next permutation = "abdc"
     *   w = "dcba" → no answer (already the highest permutation)
     *
     * @param w the input string
     * @return the next lexicographically greater string, or "no answer" if none exists
     *
     * Constraints:
     *  - 1 ≤ length of w ≤ 100
     *  - w contains only lowercase ASCII letters [a-z]
     *
     * Problem reference: HackerRank "Bigger is Greater"
     */
    public static String biggerIsGreater(String w) {
        // abcd
        // abd c
        // a c das
        // acdb
        char[] chars = w.toCharArray();
        for (int i = chars.length-1; i >= 1; i--) {
            if (chars[i - 1] < chars[i]) {
                for (int j = chars.length-1; j >= 1; j--) {
                    if (chars[i - 1] < chars[j]) {
                        char swap = chars[i - 1];
                        chars[i - 1] = chars[j];
                        chars[j] = swap;
                        Arrays.sort(chars, i, chars.length); // Could just insert -> move
                        return new String(chars);
                    }
                }
            }
        }
        return "no answer";
    }

    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));

        int T = Integer.parseInt(bufferedReader.readLine().trim());

        IntStream.range(0, T).forEach(TItr -> {
            try {
                String w = bufferedReader.readLine();

                String result = BiggerIsGreater.biggerIsGreater(w);

                System.out.println(result);
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        });

        bufferedReader.close();
    }
}