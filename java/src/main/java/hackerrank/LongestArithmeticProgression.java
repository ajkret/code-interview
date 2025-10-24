package hackerrank;

import java.io.*;
import java.util.*;
import java.util.stream.*;
import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;

public class LongestArithmeticProgression {

    /*
     * Complete the 'findLongestArithmeticProgression' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts the following parameters:
     *  1. INTEGER_ARRAY arr
     *  2. INTEGER k
     */

    public static int findLongestArithmeticProgression(List<Integer> arr, int k) {
        if (arr == null || arr.isEmpty()) return 0;

        // Remove duplicates
        Set<Integer> elements = new HashSet<>(arr);
        int longest = 0;

        for (int num : elements) {
            if (!elements.contains(num - k)) {
                int current = num;
                int length = 1;

                while (elements.contains(current + k)) {
                    current += k;
                    length++;
                }

                longest = Math.max(longest, length);
            }
        }

        return longest;
    }

    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));

        // Get the size of the array on input (each element in a new line)
        int arrCount = Integer.parseInt(bufferedReader.readLine().trim());

        // Read elements
        List<Integer> arr = IntStream.range(0, arrCount).mapToObj(i -> {
                    try {
                        return bufferedReader.readLine().replaceAll("\\s+$", "");
                    } catch (IOException ex) {
                        throw new RuntimeException(ex);
                    }
                })
                .map(String::trim)
                .map(Integer::parseInt)
                .collect(toList());

        // Read k
        int k = Integer.parseInt(bufferedReader.readLine().trim());

        int result = LongestArithmeticProgression.findLongestArithmeticProgression(arr, k);

        System.out.println(result);

        bufferedReader.close();

    }
}
