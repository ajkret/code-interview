package graph

import (
    "golang.org/x/exp/constraints"
)

// NodeID is defined as a type alias for any comparable type,
// ensuring it can be used as a key in maps.
type NodeID comparable

// Edge represents a directed edge connecting two nodes.
// T: Type of the Vertex ID, W: Type of the Weight.
type Edge[T NodeID, W constraints.Ordered] struct {
    To     T // The destination node (type T)
    Weight W // The cost/distance of the edge (type W)
}

// Graph [T, W] is the concrete implementation that uses an adjacency map.
// T: Type of the Vertex ID (must be comparable for map keys).
// W: Type of the Edge Weight (must be ordered).
type Graph[T NodeID, W constraints.Ordered] struct {
    // Edges: Maps the source node (T) to a slice of outgoing edges (Edges).
    Edges map[T][]Edge[T, W]
    // Vertices: A map used to store the set of all vertices (for easy lookup and iteration).
    Vertices map[T]bool
}

// NewGraph creates and initializes a new Graph.
func NewGraph[T NodeID, W constraints.Ordered]() *Graph[T, W] {
    return &Graph[T, W]{
        Edges:    make(map[T][]Edge[T, W]),
        Vertices: make(map[T]bool),
    }
}

// AddVertex ensures a vertex ID exists in the graph's set of vertices
// and prepares its entry in the edges map.
func (g *Graph[T, W]) AddVertex(id T) {
    g.Vertices[id] = true

    // Ensure that even vertices with no outgoing edges exist in the edges map
    if _, exists := g.Edges[id]; !exists {
        g.Edges[id] = []Edge[T, W]{}
    }
}

// AddEdge adds a directed edge from 'from' to 'to' with a specific 'weight'.
func (g *Graph[T, W]) AddEdge(from T, to T, weight W) {
    // 1. Ensure vertices exist (including the 'to' destination)
    g.AddVertex(from)
    g.AddVertex(to)

    // 2. Add the edge to the list of edges for the 'from' node
    g.Edges[from] = append(g.Edges[from], Edge[T, W]{To: to, Weight: weight})
}

// GetVertices returns a slice containing all vertex IDs in the graph.
func (g *Graph[T, W]) GetVertices() []T {
    ids := make([]T, 0, len(g.Vertices))
    for id := range g.Vertices {
        ids = append(ids, id)
    }
    return ids
}

// GetEdgesFrom returns all outgoing edges from a specific node 'from'.
// Essential for all search algorithms (BFS, Dijkstra, Bellman-Ford).
func (g *Graph[T, W]) GetEdgesFrom(from T) []Edge[T, W] {
    if edges, ok := g.Edges[from]; ok {
        return edges
    }
    return nil
}
