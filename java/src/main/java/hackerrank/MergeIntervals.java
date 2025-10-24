package hackerrank;

import java.io.*;
import java.util.*;
import java.util.stream.*;

import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;


class MergeIntervals {

    /*
     * Complete the 'mergeHighDefinitionIntervals' function below.
     *
     * The function is expected to return a 2D_INTEGER_ARRAY.
     * The function accepts 2D_INTEGER_ARRAY intervals as parameter.
     */

    public static List<List<Integer>> mergeHighDefinitionIntervals(List<List<Integer>> intervals) {

        var left = intervals.stream().map(i -> i.get(0)).sorted().collect(toList());
        var right = intervals.stream().map(i -> i.get(1)).sorted().collect(toList());
        var result = new ArrayList<List<Integer>>();

        if (intervals.isEmpty()) {
            return Arrays.asList();
        }
        int intervalLeft = left.get(0);
        int intervalRight = right.get(0);
        int index = 0;

        for (; index < intervals.size(); index++) {

            if (left.get(index) > intervalRight) {
                result.add(Arrays.asList(intervalLeft, intervalRight));
                intervalLeft = left.get(index);
            }
            intervalRight = right.get(index);
        }

        // Last interval
        if (index > 0) {
            result.add(Arrays.asList(intervalLeft, intervalRight));
        }

        return result;

    }

    // Example entry [[1, 3], [2, 6], [8, 10], [15, 18]]
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));

        int intervalsRows = Integer.parseInt(bufferedReader.readLine().trim());
        int intervalsColumns = Integer.parseInt(bufferedReader.readLine().trim());

        List<List<Integer>> intervals = new ArrayList<>();

        IntStream.range(0, intervalsRows).forEach(i -> {
            try {
                intervals.add(
                        Stream.of(bufferedReader.readLine().replaceAll("\\s+$", "").split(" "))
                                .map(Integer::parseInt)
                                .collect(toList())
                );
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        });

        List<List<Integer>> result = MergeIntervals.mergeHighDefinitionIntervals(intervals);

        result.stream()
                .map(
                        r -> r.stream()
                                .map(Object::toString)
                                .collect(joining(" "))
                )
                .collect(toList())
                .forEach(System.out::println);

        bufferedReader.close();
    }
}