package graph

import (
	"slices"

	"github.com/harry713j/dsa_practice/constraints"
)

// Implementation of graph using adjacency list
type Graph[T constraints.Ordered] struct {
	adj      map[T][]T
	directed bool // indicate if the graph is a directed graph or not
	vertices int  // size of the vertices
}

func New[T constraints.Ordered](directed bool) *Graph[T] {
	return &Graph[T]{
		adj:      make(map[T][]T, 0),
		directed: directed,
	}
}

// Add a new vertex to the graph if it is not exists
func (g *Graph[T]) AddVertex(v T) {
	if _, ok := g.adj[v]; !ok {
		g.adj[v] = []T{}
		g.vertices++
	}
}

// Add a edge between two vertices
func (g *Graph[T]) AddEdge(from, to T) {
	g.AddVertex(from)
	g.AddVertex(to)

	if !slices.Contains(g.adj[from], to) {
		g.adj[from] = append(g.adj[from], to)
	}

	// if the graph is undirected
	if !g.directed {
		g.adj[to] = append(g.adj[to], from)
	}
}
