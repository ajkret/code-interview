package graph

// Graph represents a generic graph structure that can support
// various graph algorithms like BFS, Dijkstra, Bellman-Ford, and A*.
type Graph[V comparable, W any] interface {
    // AddVertex adds a vertex to the graph
    AddVertex(vertex V)

    // AddEdge adds an edge between two vertices with an optional weight
    // For unweighted graphs, weight can be ignored or set to a default value
    AddEdge(from, to V, weight W)

    // Vertices returns all vertices in the graph
    Vertices() []V

    // Neighbors returns all adjacent vertices for a given vertex
    Neighbors(vertex V) []V

    // Weight returns the weight of the edge between two vertices
    // Returns the weight and a boolean indicating if the edge exists
    Weight(from, to V) (W, bool)

    // HasVertex checks if a vertex exists in the graph
    HasVertex(vertex V) bool

    // HasEdge checks if an edge exists between two vertices
    HasEdge(from, to V) bool
}
