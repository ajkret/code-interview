package app

import (
    "bufio"
    "fmt"
    "interview_go/internal/util/heap"
    "os"
    "strconv"
    "strings"
)

type Point struct {
    x        int
    y        int
    distance int
}

// K Closest Points
// Given a list of points on a 2D plane. Find k closest points to the origin (0, 0).
//
// Input: [(1, 1), (2, 2), (3, 3)], 1
//
// Output: [(1, 1)]
// Distance between two points (x1, y1) and (x2, y2) is sqrt((x1 - x2)^2 + (y1 - y2)^2).
// For distance to the origin, (x2, y2) is (0, 0) so the distance becomes sqrt(x1^2 + y1^2)
// parameters:
// - points: a list of points
// - k: the number of closest points
// return: a list of k closest points
//
func kClosestPoints(points [][]int, k int) [][]int {
    var resultHeap = heap.NewMinHeap[*Point](func(a, b *Point) int {
        return a.distance - b.distance
    })

    // Why use x² + y² instead of sqrt(x² + y²)?
    // When calculating the distance from a point (x, y) to the origin (0, 0), the formula is:
    //
    // distance = sqrt(x² + y²)
    // But in the code, we use:
    //
    // distance = x² + y²
    // The Reason: Comparing Distances
    // The square root function (sqrt) is a monotonically increasing function. This means that if a > b, then sqrt(a) > sqrt(b).
    // When you only need to compare which point is closer or farther, you don't need the actual distance value—just the order.
    // Comparing x² + y² is enough because the point with the smaller x² + y² will also have the smaller sqrt(x² + y²).

    for _, point := range points {
        resultHeap.Push(&Point{point[0], point[1], point[0]*point[0] + point[1]*point[1]})
    }

    result := make([][]int, 0, k)
    for i := 0; i < k; i++ {
        point, ok := resultHeap.Pop()
        if ok {
            result = append(result, []int{point.x, point.y})
        }
    }

    return result
}

func arrayAtoi(arr []string) []int {
    res := []int{}
    for _, x := range arr {
        v, _ := strconv.Atoi(x)
        res = append(res, v)
    }
    return res
}

func splitWords(s string) []string {
    if s == "" {
        return []string{}
    }
    return strings.Split(s, " ")
}

func arrayItoa(arr []int) []string {
    res := []string{}
    for _, v := range arr {
        res = append(res, strconv.Itoa(v))
    }
    return res
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    pointsLength, _ := strconv.Atoi(scanner.Text())
    points := [][]int{}
    for i := 0; i < pointsLength; i++ {
        scanner.Scan()
        points = append(points, arrayAtoi(splitWords(scanner.Text())))
    }
    scanner.Scan()
    k, _ := strconv.Atoi(scanner.Text())
    res := kClosestPoints(points, k)
    for _, row := range res {
        fmt.Println(strings.Join(arrayItoa(row), " "))
    }
}
