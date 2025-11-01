package graph

import (
    "sort"
    "testing"

    "github.com/stretchr/testify/assert"
)

// TestNewGraph tests graph initialization
func TestNewGraph(t *testing.T) {
    g := NewGraph[string, int]()

    assert.NotNil(t, g, "Graph should not be nil")
    assert.NotNil(t, g.Edges, "Edges map should be initialized")
    assert.NotNil(t, g.Vertices, "Vertices map should be initialized")
    assert.Equal(t, 0, len(g.Vertices), "New graph should have no vertices")
    assert.Equal(t, 0, len(g.Edges), "New graph should have no edges")
}

// TestNewGraph_DifferentTypes tests graph with different generic types
func TestNewGraph_DifferentTypes(t *testing.T) {
    t.Run("IntVertex_FloatWeight", func(t *testing.T) {
        g := NewGraph[int, float64]()
        assert.NotNil(t, g)
    })

    t.Run("StringVertex_IntWeight", func(t *testing.T) {
        g := NewGraph[string, int]()
        assert.NotNil(t, g)
    })

    t.Run("IntVertex_IntWeight", func(t *testing.T) {
        g := NewGraph[int, int]()
        assert.NotNil(t, g)
    })
}

// TestAddVertex_SingleVertex tests adding a single vertex
func TestAddVertex_SingleVertex(t *testing.T) {
    g := NewGraph[string, int]()

    g.AddVertex("A")

    assert.Equal(t, 1, len(g.Vertices), "Should have 1 vertex")
    assert.True(t, g.Vertices["A"], "Vertex A should exist")
    assert.NotNil(t, g.Edges["A"], "Edge entry for A should exist")
    assert.Equal(t, 0, len(g.Edges["A"]), "Vertex A should have no outgoing edges")
}

// TestAddVertex_MultipleVertices tests adding multiple vertices
func TestAddVertex_MultipleVertices(t *testing.T) {
    g := NewGraph[string, int]()

    g.AddVertex("A")
    g.AddVertex("B")
    g.AddVertex("C")

    assert.Equal(t, 3, len(g.Vertices), "Should have 3 vertices")
    assert.True(t, g.Vertices["A"], "Vertex A should exist")
    assert.True(t, g.Vertices["B"], "Vertex B should exist")
    assert.True(t, g.Vertices["C"], "Vertex C should exist")
}

// TestAddVertex_Duplicate tests adding duplicate vertices
func TestAddVertex_Duplicate(t *testing.T) {
    g := NewGraph[string, int]()

    g.AddVertex("A")
    g.AddVertex("A")
    g.AddVertex("A")

    assert.Equal(t, 1, len(g.Vertices), "Should have only 1 vertex")
    assert.True(t, g.Vertices["A"], "Vertex A should exist")
}

// TestAddVertex_IntegerVertices tests with integer vertex IDs
func TestAddVertex_IntegerVertices(t *testing.T) {
    g := NewGraph[int, int]()

    g.AddVertex(1)
    g.AddVertex(2)
    g.AddVertex(100)

    assert.Equal(t, 3, len(g.Vertices), "Should have 3 vertices")
    assert.True(t, g.Vertices[1], "Vertex 1 should exist")
    assert.True(t, g.Vertices[2], "Vertex 2 should exist")
    assert.True(t, g.Vertices[100], "Vertex 100 should exist")
}

// TestAddEdge_SingleEdge tests adding a single edge
func TestAddEdge_SingleEdge(t *testing.T) {
    g := NewGraph[string, int]()

    g.AddEdge("A", "B", 5)

    assert.Equal(t, 2, len(g.Vertices), "Should have 2 vertices")
    assert.True(t, g.Vertices["A"], "Vertex A should exist")
    assert.True(t, g.Vertices["B"], "Vertex B should exist")
    assert.Equal(t, 1, len(g.Edges["A"]), "A should have 1 outgoing edge")
    assert.Equal(t, "B", g.Edges["A"][0].To, "Edge should point to B")
    assert.Equal(t, 5, g.Edges["A"][0].Weight, "Edge weight should be 5")
}

// TestAddEdge_MultipleEdges tests adding multiple edges from same vertex
func TestAddEdge_MultipleEdges(t *testing.T) {
    g := NewGraph[string, int]()

    g.AddEdge("A", "B", 5)
    g.AddEdge("A", "C", 3)
    g.AddEdge("A", "D", 7)

    assert.Equal(t, 4, len(g.Vertices), "Should have 4 vertices")
    assert.Equal(t, 3, len(g.Edges["A"]), "A should have 3 outgoing edges")

    // Verify all edges exist
    edges := g.Edges["A"]
    destinations := make(map[string]int)
    for _, edge := range edges {
        destinations[edge.To] = edge.Weight
    }

    assert.Equal(t, 5, destinations["B"], "Edge A->B should have weight 5")
    assert.Equal(t, 3, destinations["C"], "Edge A->C should have weight 3")
    assert.Equal(t, 7, destinations["D"], "Edge A->D should have weight 7")
}

// TestAddEdge_ComplexGraph tests building a more complex graph
func TestAddEdge_ComplexGraph(t *testing.T) {
    g := NewGraph[string, int]()

    // Build a simple directed graph:
    // A -> B (weight: 4)
    // A -> C (weight: 2)
    // B -> C (weight: 1)
    // B -> D (weight: 5)
    // C -> D (weight: 3)
    g.AddEdge("A", "B", 4)
    g.AddEdge("A", "C", 2)
    g.AddEdge("B", "C", 1)
    g.AddEdge("B", "D", 5)
    g.AddEdge("C", "D", 3)

    assert.Equal(t, 4, len(g.Vertices), "Should have 4 vertices")
    assert.Equal(t, 2, len(g.Edges["A"]), "A should have 2 outgoing edges")
    assert.Equal(t, 2, len(g.Edges["B"]), "B should have 2 outgoing edges")
    assert.Equal(t, 1, len(g.Edges["C"]), "C should have 1 outgoing edge")
    assert.Equal(t, 0, len(g.Edges["D"]), "D should have 0 outgoing edges")
}

// TestAddEdge_DuplicateEdges tests adding duplicate edges
func TestAddEdge_DuplicateEdges(t *testing.T) {
    g := NewGraph[string, int]()

    g.AddEdge("A", "B", 5)
    g.AddEdge("A", "B", 5)
    g.AddEdge("A", "B", 10)

    // Duplicate edges are allowed (multi-graph)
    assert.Equal(t, 3, len(g.Edges["A"]), "Should have 3 edges (duplicates allowed)")
}

// TestAddEdge_SelfLoop tests adding self-loop edges
func TestAddEdge_SelfLoop(t *testing.T) {
    g := NewGraph[string, int]()

    g.AddEdge("A", "A", 1)

    assert.Equal(t, 1, len(g.Vertices), "Should have 1 vertex")
    assert.Equal(t, 1, len(g.Edges["A"]), "A should have 1 self-loop edge")
    assert.Equal(t, "A", g.Edges["A"][0].To, "Edge should point to itself")
}

// TestAddEdge_IntegerGraph tests graph with integer vertices and weights
func TestAddEdge_IntegerGraph(t *testing.T) {
    g := NewGraph[int, int]()

    g.AddEdge(1, 2, 10)
    g.AddEdge(2, 3, 20)
    g.AddEdge(3, 1, 30)

    assert.Equal(t, 3, len(g.Vertices), "Should have 3 vertices")
    assert.Equal(t, 10, g.Edges[1][0].Weight, "Edge 1->2 should have weight 10")
    assert.Equal(t, 20, g.Edges[2][0].Weight, "Edge 2->3 should have weight 20")
    assert.Equal(t, 30, g.Edges[3][0].Weight, "Edge 3->1 should have weight 30")
}

// TestAddEdge_FloatWeights tests graph with float weights
func TestAddEdge_FloatWeights(t *testing.T) {
    g := NewGraph[string, float64]()

    g.AddEdge("A", "B", 1.5)
    g.AddEdge("B", "C", 2.7)
    g.AddEdge("C", "A", 0.3)

    assert.Equal(t, 3, len(g.Vertices), "Should have 3 vertices")
    assert.InDelta(t, 1.5, g.Edges["A"][0].Weight, 0.001, "Weight should be 1.5")
    assert.InDelta(t, 2.7, g.Edges["B"][0].Weight, 0.001, "Weight should be 2.7")
    assert.InDelta(t, 0.3, g.Edges["C"][0].Weight, 0.001, "Weight should be 0.3")
}

// TestAddEdge_NegativeWeights tests graph with negative weights
func TestAddEdge_NegativeWeights(t *testing.T) {
    g := NewGraph[string, int]()

    g.AddEdge("A", "B", -5)
    g.AddEdge("B", "C", -10)

    assert.Equal(t, -5, g.Edges["A"][0].Weight, "Should support negative weight -5")
    assert.Equal(t, -10, g.Edges["B"][0].Weight, "Should support negative weight -10")
}

// TestAddEdge_ZeroWeight tests edges with zero weight
func TestAddEdge_ZeroWeight(t *testing.T) {
    g := NewGraph[string, int]()

    g.AddEdge("A", "B", 0)

    assert.Equal(t, 1, len(g.Edges["A"]), "Should have 1 edge")
    assert.Equal(t, 0, g.Edges["A"][0].Weight, "Edge weight should be 0")
}

// TestGetVertices_EmptyGraph tests getting vertices from empty graph
func TestGetVertices_EmptyGraph(t *testing.T) {
    g := NewGraph[string, int]()

    vertices := g.GetVertices()

    assert.NotNil(t, vertices, "Should return non-nil slice")
    assert.Equal(t, 0, len(vertices), "Empty graph should have no vertices")
}

// TestGetVertices_SingleVertex tests getting single vertex
func TestGetVertices_SingleVertex(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddVertex("A")

    vertices := g.GetVertices()

    assert.Equal(t, 1, len(vertices), "Should return 1 vertex")
    assert.Contains(t, vertices, "A", "Should contain vertex A")
}

// TestGetVertices_MultipleVertices tests getting multiple vertices
func TestGetVertices_MultipleVertices(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddVertex("A")
    g.AddVertex("B")
    g.AddVertex("C")

    vertices := g.GetVertices()

    assert.Equal(t, 3, len(vertices), "Should return 3 vertices")
    assert.Contains(t, vertices, "A", "Should contain vertex A")
    assert.Contains(t, vertices, "B", "Should contain vertex B")
    assert.Contains(t, vertices, "C", "Should contain vertex C")
}

// TestGetVertices_AfterAddingEdges tests getting vertices after adding edges
func TestGetVertices_AfterAddingEdges(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddEdge("A", "B", 1)
    g.AddEdge("B", "C", 2)

    vertices := g.GetVertices()

    assert.Equal(t, 3, len(vertices), "Should have 3 vertices")
    assert.Contains(t, vertices, "A", "Should contain vertex A")
    assert.Contains(t, vertices, "B", "Should contain vertex B")
    assert.Contains(t, vertices, "C", "Should contain vertex C")
}

// TestGetVertices_OrderIndependent tests that vertex order doesn't matter
func TestGetVertices_OrderIndependent(t *testing.T) {
    g := NewGraph[int, int]()
    g.AddVertex(3)
    g.AddVertex(1)
    g.AddVertex(2)

    vertices := g.GetVertices()
    sort.Ints(vertices)

    expected := []int{1, 2, 3}
    assert.Equal(t, expected, vertices, "Should contain all vertices regardless of insertion order")
}

// TestGetEdgesFrom_NonExistentVertex tests getting edges from non-existent vertex
func TestGetEdgesFrom_NonExistentVertex(t *testing.T) {
    g := NewGraph[string, int]()

    edges := g.GetEdgesFrom("A")

    assert.Nil(t, edges, "Should return nil for non-existent vertex")
}

// TestGetEdgesFrom_VertexWithNoEdges tests getting edges from vertex with no outgoing edges
func TestGetEdgesFrom_VertexWithNoEdges(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddVertex("A")

    edges := g.GetEdgesFrom("A")

    assert.NotNil(t, edges, "Should return non-nil slice")
    assert.Equal(t, 0, len(edges), "Vertex with no edges should return empty slice")
}

// TestGetEdgesFrom_SingleEdge tests getting single edge
func TestGetEdgesFrom_SingleEdge(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddEdge("A", "B", 5)

    edges := g.GetEdgesFrom("A")

    assert.Equal(t, 1, len(edges), "Should return 1 edge")
    assert.Equal(t, "B", edges[0].To, "Edge should point to B")
    assert.Equal(t, 5, edges[0].Weight, "Edge weight should be 5")
}

// TestGetEdgesFrom_MultipleEdges tests getting multiple edges
func TestGetEdgesFrom_MultipleEdges(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddEdge("A", "B", 5)
    g.AddEdge("A", "C", 3)
    g.AddEdge("A", "D", 7)

    edges := g.GetEdgesFrom("A")

    assert.Equal(t, 3, len(edges), "Should return 3 edges")

    // Verify all edges
    destinations := make(map[string]int)
    for _, edge := range edges {
        destinations[edge.To] = edge.Weight
    }

    assert.Equal(t, 5, destinations["B"], "Should have edge to B with weight 5")
    assert.Equal(t, 3, destinations["C"], "Should have edge to C with weight 3")
    assert.Equal(t, 7, destinations["D"], "Should have edge to D with weight 7")
}

// TestGetEdgesFrom_DestinationVertex tests getting edges from destination-only vertex
func TestGetEdgesFrom_DestinationVertex(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddEdge("A", "B", 5)

    // B is only a destination, not a source
    edges := g.GetEdgesFrom("B")

    assert.NotNil(t, edges, "Should return non-nil slice")
    assert.Equal(t, 0, len(edges), "Destination vertex should have no outgoing edges")
}

// TestGraph_IsolatedVertices tests graph with isolated vertices
func TestGraph_IsolatedVertices(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddVertex("A")
    g.AddVertex("B")
    g.AddVertex("C")
    g.AddEdge("D", "E", 1)

    vertices := g.GetVertices()

    assert.Equal(t, 5, len(vertices), "Should have 5 vertices")
    assert.Equal(t, 0, len(g.GetEdgesFrom("A")), "A should have no edges")
    assert.Equal(t, 0, len(g.GetEdgesFrom("B")), "B should have no edges")
    assert.Equal(t, 0, len(g.GetEdgesFrom("C")), "C should have no edges")
}

// TestGraph_DirectedNature tests that graph is directed (A->B doesn't create B->A)
func TestGraph_DirectedNature(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddEdge("A", "B", 5)

    edgesFromA := g.GetEdgesFrom("A")
    edgesFromB := g.GetEdgesFrom("B")

    assert.Equal(t, 1, len(edgesFromA), "A should have 1 outgoing edge")
    assert.Equal(t, 0, len(edgesFromB), "B should have no outgoing edges (directed graph)")
}

// TestGraph_BidirectionalEdges tests creating bidirectional edges manually
func TestGraph_BidirectionalEdges(t *testing.T) {
    g := NewGraph[string, int]()
    g.AddEdge("A", "B", 5)
    g.AddEdge("B", "A", 5)

    edgesFromA := g.GetEdgesFrom("A")
    edgesFromB := g.GetEdgesFrom("B")

    assert.Equal(t, 1, len(edgesFromA), "A should have 1 outgoing edge")
    assert.Equal(t, "B", edgesFromA[0].To, "A's edge should point to B")
    assert.Equal(t, 1, len(edgesFromB), "B should have 1 outgoing edge")
    assert.Equal(t, "A", edgesFromB[0].To, "B's edge should point to A")
}

// TestGraph_CompleteGraph tests creating a complete graph (all vertices connected)
func TestGraph_CompleteGraph(t *testing.T) {
    g := NewGraph[int, int]()

    // Create a complete graph with 4 vertices
    vertices := []int{1, 2, 3, 4}
    for _, from := range vertices {
        for _, to := range vertices {
            if from != to {
                g.AddEdge(from, to, from*10+to)
            }
        }
    }

    assert.Equal(t, 4, len(g.GetVertices()), "Should have 4 vertices")
    for _, v := range vertices {
        edges := g.GetEdgesFrom(v)
        assert.Equal(t, 3, len(edges), "Each vertex should have 3 outgoing edges")
    }
}

// TestGraph_LargeGraph tests graph with many vertices and edges
func TestGraph_LargeGraph(t *testing.T) {
    g := NewGraph[int, int]()

    // Add 100 vertices
    for i := 0; i < 100; i++ {
        g.AddVertex(i)
    }

    // Add 200 edges
    for i := 0; i < 100; i++ {
        g.AddEdge(i, (i+1)%100, i)
        g.AddEdge(i, (i+2)%100, i*2)
    }

    vertices := g.GetVertices()
    assert.Equal(t, 100, len(vertices), "Should have 100 vertices")

    totalEdges := 0
    for i := 0; i < 100; i++ {
        totalEdges += len(g.GetEdgesFrom(i))
    }
    assert.Equal(t, 200, totalEdges, "Should have 200 edges total")
}

// TestGraph_ChainGraph tests linear chain graph structure
func TestGraph_ChainGraph(t *testing.T) {
    g := NewGraph[string, int]()

    // Create chain: A -> B -> C -> D -> E
    g.AddEdge("A", "B", 1)
    g.AddEdge("B", "C", 2)
    g.AddEdge("C", "D", 3)
    g.AddEdge("D", "E", 4)

    assert.Equal(t, 5, len(g.GetVertices()), "Should have 5 vertices")
    assert.Equal(t, 1, len(g.GetEdgesFrom("A")), "A should have 1 edge")
    assert.Equal(t, 1, len(g.GetEdgesFrom("B")), "B should have 1 edge")
    assert.Equal(t, 1, len(g.GetEdgesFrom("C")), "C should have 1 edge")
    assert.Equal(t, 1, len(g.GetEdgesFrom("D")), "D should have 1 edge")
    assert.Equal(t, 0, len(g.GetEdgesFrom("E")), "E should have no edges")
}

// TestGraph_CyclicGraph tests graph with cycles
func TestGraph_CyclicGraph(t *testing.T) {
    g := NewGraph[string, int]()

    // Create cycle: A -> B -> C -> A
    g.AddEdge("A", "B", 1)
    g.AddEdge("B", "C", 2)
    g.AddEdge("C", "A", 3)

    assert.Equal(t, 3, len(g.GetVertices()), "Should have 3 vertices")
    assert.Equal(t, 1, len(g.GetEdgesFrom("A")), "A should have 1 edge")
    assert.Equal(t, 1, len(g.GetEdgesFrom("B")), "B should have 1 edge")
    assert.Equal(t, 1, len(g.GetEdgesFrom("C")), "C should have 1 edge")
}

// TestGraph_StarGraph tests star topology (one central node)
func TestGraph_StarGraph(t *testing.T) {
    g := NewGraph[string, int]()

    // Create star: Center -> A, B, C, D
    g.AddEdge("Center", "A", 1)
    g.AddEdge("Center", "B", 2)
    g.AddEdge("Center", "C", 3)
    g.AddEdge("Center", "D", 4)

    assert.Equal(t, 5, len(g.GetVertices()), "Should have 5 vertices")
    assert.Equal(t, 4, len(g.GetEdgesFrom("Center")), "Center should have 4 edges")
    assert.Equal(t, 0, len(g.GetEdgesFrom("A")), "Leaf nodes should have no edges")
}

// ... existing code ...

// City represents a complex object to be used as a vertex in the graph
type City struct {
    Name    string
    Country string
    ZipCode int
}

// TestGraph_ComplexObjectAsVertex tests using a complex struct as vertex type
func TestGraph_ComplexObjectAsVertex(t *testing.T) {
    g := NewGraph[City, float64]()

    // Define cities
    newYork := City{Name: "New York", Country: "USA", ZipCode: 10001}
    london := City{Name: "London", Country: "UK", ZipCode: 12345}
    tokyo := City{Name: "Tokyo", Country: "Japan", ZipCode: 54321}
    var _ = City{Name: "Paris", Country: "France", ZipCode: 75001}

    // Add vertices
    g.AddVertex(newYork)
    g.AddVertex(london)
    g.AddVertex(tokyo)

    vertices := g.GetVertices()
    assert.Equal(t, 3, len(vertices), "Should have 3 cities")
    assert.Contains(t, vertices, newYork, "Should contain New York")
    assert.Contains(t, vertices, london, "Should contain London")
    assert.Contains(t, vertices, tokyo, "Should contain Tokyo")
}

// TestGraph_ComplexObjectWithEdges tests adding edges between complex objects
func TestGraph_ComplexObjectWithEdges(t *testing.T) {
    g := NewGraph[City, float64]()

    // Define cities
    newYork := City{Name: "New York", Country: "USA", ZipCode: 10001}
    london := City{Name: "London", Country: "UK", ZipCode: 12345}
    tokyo := City{Name: "Tokyo", Country: "Japan", ZipCode: 54321}
    paris := City{Name: "Paris", Country: "France", ZipCode: 75001}

    // Add edges with distances in thousands of kilometers
    g.AddEdge(newYork, london, 5.585) // ~5,585 km
    g.AddEdge(newYork, tokyo, 10.838) // ~10,838 km
    g.AddEdge(london, paris, 0.344)   // ~344 km
    g.AddEdge(paris, tokyo, 9.715)    // ~9,715 km
    g.AddEdge(tokyo, newYork, 10.838) // ~10,838 km (return flight)

    // Verify vertices were created
    assert.Equal(t, 4, len(g.GetVertices()), "Should have 4 cities")

    // Verify edges from New York
    edgesFromNY := g.GetEdgesFrom(newYork)
    assert.Equal(t, 2, len(edgesFromNY), "New York should have 2 outgoing edges")

    destinations := make(map[City]float64)
    for _, edge := range edgesFromNY {
        destinations[edge.To] = edge.Weight
    }

    assert.InDelta(t, 5.585, destinations[london], 0.001, "Distance NY->London should be 5.585")
    assert.InDelta(t, 10.838, destinations[tokyo], 0.001, "Distance NY->Tokyo should be 10.838")

    // Verify edges from London
    edgesFromLondon := g.GetEdgesFrom(london)
    assert.Equal(t, 1, len(edgesFromLondon), "London should have 1 outgoing edge")
    assert.Equal(t, paris, edgesFromLondon[0].To, "London should connect to Paris")
    assert.InDelta(t, 0.344, edgesFromLondon[0].Weight, 0.001, "Distance London->Paris should be 0.344")

    // Verify edges from Paris
    edgesFromParis := g.GetEdgesFrom(paris)
    assert.Equal(t, 1, len(edgesFromParis), "Paris should have 1 outgoing edge")
    assert.Equal(t, tokyo, edgesFromParis[0].To, "Paris should connect to Tokyo")

    // Verify edges from Tokyo
    edgesFromTokyo := g.GetEdgesFrom(tokyo)
    assert.Equal(t, 1, len(edgesFromTokyo), "Tokyo should have 1 outgoing edge")
    assert.Equal(t, newYork, edgesFromTokyo[0].To, "Tokyo should connect back to New York")
}

// TestGraph_ComplexObjectEquality tests that struct equality works as expected
func TestGraph_ComplexObjectEquality(t *testing.T) {
    g := NewGraph[City, int]()

    city1 := City{Name: "Boston", Country: "USA", ZipCode: 2101}
    city2 := City{Name: "Boston", Country: "USA", ZipCode: 2101}
    city3 := City{Name: "Boston", Country: "USA", ZipCode: 2102} // Different zip code

    // city1 and city2 are equal (same values)
    g.AddVertex(city1)
    g.AddVertex(city2) // Should not add a new vertex

    assert.Equal(t, 1, len(g.GetVertices()), "city1 and city2 should be treated as the same vertex")

    // city3 is different
    g.AddVertex(city3)
    assert.Equal(t, 2, len(g.GetVertices()), "city3 should be a different vertex")
}

// TestGraph_ComplexObjectIsolatedVertex tests isolated vertices with complex objects
func TestGraph_ComplexObjectIsolatedVertex(t *testing.T) {
    g := NewGraph[City, float64]()

    newYork := City{Name: "New York", Country: "USA", ZipCode: 10001}
    london := City{Name: "London", Country: "UK", ZipCode: 12345}
    tokyo := City{Name: "Tokyo", Country: "Japan", ZipCode: 54321}

    // Add only New York with edges
    g.AddEdge(newYork, london, 5.585)

    // Add Tokyo as isolated vertex
    g.AddVertex(tokyo)

    assert.Equal(t, 3, len(g.GetVertices()), "Should have 3 cities")

    // Tokyo should have no outgoing edges
    edgesFromTokyo := g.GetEdgesFrom(tokyo)
    assert.NotNil(t, edgesFromTokyo, "Should return non-nil slice for Tokyo")
    assert.Equal(t, 0, len(edgesFromTokyo), "Tokyo should have no outgoing edges")

    // London should have no outgoing edges (only destination)
    edgesFromLondon := g.GetEdgesFrom(london)
    assert.NotNil(t, edgesFromLondon, "Should return non-nil slice for London")
    assert.Equal(t, 0, len(edgesFromLondon), "London should have no outgoing edges")
}

// Person represents another complex object with pointer field (still comparable)
type Person struct {
    ID   int
    Name string
    Age  int
}

// TestGraph_ComplexObjectPerson tests graph with Person objects
func TestGraph_ComplexObjectPerson(t *testing.T) {
    g := NewGraph[Person, int]()

    alice := Person{ID: 1, Name: "Alice", Age: 30}
    bob := Person{ID: 2, Name: "Bob", Age: 25}
    charlie := Person{ID: 3, Name: "Charlie", Age: 35}

    // Build a social network graph (trust scores as weights)
    g.AddEdge(alice, bob, 85)     // Alice trusts Bob with score 85
    g.AddEdge(alice, charlie, 70) // Alice trusts Charlie with score 70
    g.AddEdge(bob, charlie, 90)   // Bob trusts Charlie with score 90
    g.AddEdge(charlie, alice, 75) // Charlie trusts Alice with score 75

    // Verify the graph structure
    assert.Equal(t, 3, len(g.GetVertices()), "Should have 3 people")

    // Check Alice's connections
    aliceEdges := g.GetEdgesFrom(alice)
    assert.Equal(t, 2, len(aliceEdges), "Alice should trust 2 people")

    trustScores := make(map[Person]int)
    for _, edge := range aliceEdges {
        trustScores[edge.To] = edge.Weight
    }

    assert.Equal(t, 85, trustScores[bob], "Alice's trust in Bob should be 85")
    assert.Equal(t, 70, trustScores[charlie], "Alice's trust in Charlie should be 70")

    // Check Bob's connections
    bobEdges := g.GetEdgesFrom(bob)
    assert.Equal(t, 1, len(bobEdges), "Bob should trust 1 person")
    assert.Equal(t, charlie, bobEdges[0].To, "Bob should trust Charlie")
    assert.Equal(t, 90, bobEdges[0].Weight, "Bob's trust in Charlie should be 90")
}

// Coordinate represents a complex object with multiple numeric fields
type Coordinate struct {
    X float64
    Y float64
    Z float64
}

// TestGraph_ComplexObjectCoordinate tests graph with 3D coordinates
func TestGraph_ComplexObjectCoordinate(t *testing.T) {
    g := NewGraph[Coordinate, float64]()

    origin := Coordinate{X: 0, Y: 0, Z: 0}
    point1 := Coordinate{X: 1, Y: 0, Z: 0}
    point2 := Coordinate{X: 0, Y: 1, Z: 0}
    point3 := Coordinate{X: 0, Y: 0, Z: 1}
    point4 := Coordinate{X: 1, Y: 1, Z: 1}

    // Build a 3D coordinate graph with euclidean distances
    g.AddEdge(origin, point1, 1.0)
    g.AddEdge(origin, point2, 1.0)
    g.AddEdge(origin, point3, 1.0)
    g.AddEdge(point1, point4, 1.732) // sqrt(3)
    g.AddEdge(point2, point4, 1.732) // sqrt(3)
    g.AddEdge(point3, point4, 1.732) // sqrt(3)

    // Verify graph structure
    assert.Equal(t, 5, len(g.GetVertices()), "Should have 5 coordinate points")

    // Origin should connect to 3 points
    originEdges := g.GetEdgesFrom(origin)
    assert.Equal(t, 3, len(originEdges), "Origin should connect to 3 points")

    // point4 should have no outgoing edges
    point4Edges := g.GetEdgesFrom(point4)
    assert.Equal(t, 0, len(point4Edges), "point4 should have no outgoing edges")

    // Verify specific edge
    point1Edges := g.GetEdgesFrom(point1)
    assert.Equal(t, 1, len(point1Edges), "point1 should have 1 outgoing edge")
    assert.Equal(t, point4, point1Edges[0].To, "point1 should connect to point4")
    assert.InDelta(t, 1.732, point1Edges[0].Weight, 0.001, "Distance should be approximately sqrt(3)")
}

// TestGraph_ComplexObjectSelfLoop tests self-loops with complex objects
func TestGraph_ComplexObjectSelfLoop(t *testing.T) {
    g := NewGraph[City, int]()

    newYork := City{Name: "New York", Country: "USA", ZipCode: 10001}

    // Add a self-loop (e.g., internal city transit cost)
    g.AddEdge(newYork, newYork, 5)

    vertices := g.GetVertices()
    assert.Equal(t, 1, len(vertices), "Should have 1 city")

    edges := g.GetEdgesFrom(newYork)
    assert.Equal(t, 1, len(edges), "Should have 1 self-loop edge")
    assert.Equal(t, newYork, edges[0].To, "Edge should point to itself")
    assert.Equal(t, 5, edges[0].Weight, "Self-loop weight should be 5")
}

// TestGraph_ComplexObjectMultipleEdges tests multiple edges between same complex objects
func TestGraph_ComplexObjectMultipleEdges(t *testing.T) {
    g := NewGraph[City, float64]()

    newYork := City{Name: "New York", Country: "USA", ZipCode: 10001}
    london := City{Name: "London", Country: "UK", ZipCode: 12345}

    // Add multiple edges (e.g., different airlines with different prices)
    g.AddEdge(newYork, london, 500.0) // Airline A
    g.AddEdge(newYork, london, 450.0) // Airline B
    g.AddEdge(newYork, london, 600.0) // Airline C

    edges := g.GetEdgesFrom(newYork)
    assert.Equal(t, 3, len(edges), "Should have 3 edges from NY to London")

    // Verify all weights exist
    weights := make([]float64, 0)
    for _, edge := range edges {
        assert.Equal(t, london, edge.To, "All edges should point to London")
        weights = append(weights, edge.Weight)
    }

    assert.Contains(t, weights, 500.0, "Should contain weight 500.0")
    assert.Contains(t, weights, 450.0, "Should contain weight 450.0")
    assert.Contains(t, weights, 600.0, "Should contain weight 600.0")
}
