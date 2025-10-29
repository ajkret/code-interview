package app

import (
    "sort"
    "testing"

    "github.com/stretchr/testify/assert"
)

// TestKClosestPoints_BasicCase tests the example from the problem description
func TestKClosestPoints_BasicCase(t *testing.T) {
    points := [][]int{{1, 1}, {2, 2}, {3, 3}}
    k := 1

    result := kClosestPoints(points, k)

    assert.Equal(t, 1, len(result), "Should return 1 point")
    assert.Equal(t, []int{1, 1}, result[0], "Closest point should be (1,1)")
}

// TestKClosestPoints_MultiplePoints tests returning multiple closest points
func TestKClosestPoints_MultiplePoints(t *testing.T) {
    points := [][]int{{1, 3}, {-2, 2}, {5, 8}, {0, 1}}
    k := 2

    result := kClosestPoints(points, k)

    assert.Equal(t, 2, len(result), "Should return 2 points")

    // Sort results for consistent comparison
    sortPoints(result)
    expected := [][]int{{0, 1}, {-2, 2}}
    sortPoints(expected)

    assert.Equal(t, expected, result, "Should return 2 closest points")
}

// TestKClosestPoints_AllPoints tests when k equals number of points
func TestKClosestPoints_AllPoints(t *testing.T) {
    points := [][]int{{1, 1}, {2, 2}, {3, 3}}
    k := 3

    result := kClosestPoints(points, k)

    assert.Equal(t, 3, len(result), "Should return all 3 points")

    // Verify all points are present
    sortPoints(result)
    sortPoints(points)
    assert.Equal(t, points, result, "Should return all points")
}

// TestKClosestPoints_SinglePoint tests with only one point
func TestKClosestPoints_SinglePoint(t *testing.T) {
    points := [][]int{{5, 5}}
    k := 1

    result := kClosestPoints(points, k)

    assert.Equal(t, 1, len(result), "Should return 1 point")
    assert.Equal(t, []int{5, 5}, result[0], "Should return the only point")
}

// TestKClosestPoints_PointAtOrigin tests when a point is at origin
func TestKClosestPoints_PointAtOrigin(t *testing.T) {
    points := [][]int{{0, 0}, {1, 1}, {2, 2}}
    k := 1

    result := kClosestPoints(points, k)

    assert.Equal(t, 1, len(result), "Should return 1 point")
    assert.Equal(t, []int{0, 0}, result[0], "Origin should be closest")
}

// TestKClosestPoints_NegativeCoordinates tests with negative coordinates
func TestKClosestPoints_NegativeCoordinates(t *testing.T) {
    points := [][]int{{-1, -1}, {-2, -2}, {1, 1}, {2, 2}}
    k := 2

    result := kClosestPoints(points, k)

    assert.Equal(t, 2, len(result), "Should return 2 points")

    // (-1,-1) and (1,1) both have distance 2, closest points
    sortPoints(result)
    expected := [][]int{{-1, -1}, {1, 1}}
    sortPoints(expected)

    assert.Equal(t, expected, result, "Should return 2 closest points")
}

// TestKClosestPoints_SameDistance tests when multiple points have same distance
func TestKClosestPoints_SameDistance(t *testing.T) {
    points := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    k := 2

    result := kClosestPoints(points, k)

    assert.Equal(t, 2, len(result), "Should return 2 points")

    // All points have distance 1, any 2 are valid
    for _, point := range result {
        distance := point[0]*point[0] + point[1]*point[1]
        assert.Equal(t, 1, distance, "Each point should have distance 1")
    }
}

// TestKClosestPoints_LargerDataset tests with more points
func TestKClosestPoints_LargerDataset(t *testing.T) {
    points := [][]int{
        {3, 3},   // distance: 18
        {5, -1},  // distance: 26
        {-2, 4},  // distance: 20
        {1, 1},   // distance: 2
        {0, 2},   // distance: 4
        {-1, -1}, // distance: 2
    }
    k := 3

    result := kClosestPoints(points, k)

    assert.Equal(t, 3, len(result), "Should return 3 points")

    // Verify all returned points have distance <= 4
    for _, point := range result {
        distance := point[0]*point[0] + point[1]*point[1]
        assert.LessOrEqual(t, distance, 4, "Each point should have distance <= 4")
    }
}

// TestKClosestPoints_DistanceCalculation tests correct distance calculation
func TestKClosestPoints_DistanceCalculation(t *testing.T) {
    points := [][]int{
        {3, 4},  // distance: 25 (3²+4²)
        {5, 12}, // distance: 169 (5²+12²)
        {1, 0},  // distance: 1
    }
    k := 1

    result := kClosestPoints(points, k)

    assert.Equal(t, 1, len(result), "Should return 1 point")
    assert.Equal(t, []int{1, 0}, result[0], "Point with distance 1 should be closest")
}

// TestKClosestPoints_KEqualsZero tests edge case when k=0
func TestKClosestPoints_KEqualsZero(t *testing.T) {
    points := [][]int{{1, 1}, {2, 2}, {3, 3}}
    k := 0

    result := kClosestPoints(points, k)

    assert.Equal(t, 0, len(result), "Should return empty result when k=0")
}

// TestKClosestPoints_EmptyInput tests with no points
func TestKClosestPoints_EmptyInput(t *testing.T) {
    points := [][]int{}
    k := 5

    result := kClosestPoints(points, k)

    assert.Equal(t, 0, len(result), "Should return empty result for empty input")
}

// TestKClosestPoints_LargeCoordinates tests with large coordinate values
func TestKClosestPoints_LargeCoordinates(t *testing.T) {
    points := [][]int{
        {1000, 1000}, // distance: 2000000
        {1, 1},       // distance: 2
        {100, 100},   // distance: 20000
    }
    k := 1

    result := kClosestPoints(points, k)

    assert.Equal(t, 1, len(result), "Should return 1 point")
    assert.Equal(t, []int{1, 1}, result[0], "Smallest distance point should be returned")
}

// TestKClosestPoints_MixedQuadrants tests points in all four quadrants
func TestKClosestPoints_MixedQuadrants(t *testing.T) {
    points := [][]int{
        {1, 1},   // Q1: distance 2
        {-1, 1},  // Q2: distance 2
        {-1, -1}, // Q3: distance 2
        {1, -1},  // Q4: distance 2
        {5, 5},   // Q1: distance 50
    }
    k := 4

    result := kClosestPoints(points, k)

    assert.Equal(t, 4, len(result), "Should return 4 points")

    // All returned points should have distance 2
    for _, point := range result {
        distance := point[0]*point[0] + point[1]*point[1]
        assert.Equal(t, 2, distance, "Each point should have distance 2")
    }
}

// Helper function to sort points for consistent comparison
func sortPoints(points [][]int) {
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] != points[j][0] {
            return points[i][0] < points[j][0]
        }
        return points[i][1] < points[j][1]
    })
}
