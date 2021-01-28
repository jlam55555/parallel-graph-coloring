package main

import (
	"math"
	"math/rand"
	"proj1/graph"
	"testing"
)

func countEdges(g graph.Graph) int {
	edges := 0

	for i := 0; i < len(g.Nodes); i++ {
		edges += len(g.Nodes[i].Adj)
	}

	return edges
}

// TestBranchingFactor verifies that the branching factor is
// approximately correct
func TestBranchingFactor(t *testing.T) {
	const maxGraphSize, maxBfRatio, threshold = 1000, 0.75, 0.20

	// try some random large graph sizes and branching factors
	for i := 0; i < 20; i++ {
		N := maxGraphSize / 2 + rand.Int() % (maxGraphSize / 2)
		desiredBf := rand.Float64() * float64(N) * maxBfRatio
		actualBf := float64(countEdges(
			graph.NewRandomGraph(N, float32(desiredBf)))) / float64(N)

		t.Logf("Test: NewRandomGraph(%d, %f)", N, desiredBf)

		if math.Abs(desiredBf-actualBf) > threshold*desiredBf {
			t.Errorf("Branching factor error. Got %f, desired %f",
				actualBf, desiredBf)
		}
	}
}

func checkIndices(g *graph.Graph, n int) bool {
	if n != len(g.Nodes) {
		return false
	}

	for i, node := range g.Nodes {
		if i != node.Index {
			return false
		}
	}

	return true
}

// TestIndices verifies that the indices are correct
func TestIndices(t *testing.T) {
	N := 1000
	bf := float32(30)

	if g := graph.NewCompleteGraph(N); !checkIndices(&g, N) {
		t.Errorf("NewCompleteGraph creates wrong indices")
	}

	if g := graph.NewRingGraph(N); !checkIndices(&g, N) {
		t.Errorf("NewRingGraph creates wrong indices")
	}

	if g := graph.NewRandomGraph(N, bf); !checkIndices(&g, N) {
		t.Errorf("NewRandomGraph creates wrong indices")
	}
}

// check that a graph is appropriately colored
func checkColoring(g *graph.Graph) bool {
	for _, node := range g.Nodes {
		for _, neighbor := range node.Adj {
			if node.Value == neighbor.Value {
				return false
			}
		}
	}
	return true
}

// TestSequential checks that the sequential coloring works
func TestSequential(t *testing.T) {
	N := 1000
	bf := float32(30)
	maxColor := 1000

	t.Logf("Test: NewCompleteGraph(%d)", N)
	g := graph.NewCompleteGraph(N)
	colorSequential(&g, maxColor)
	if !checkColoring(&g) {
		t.Errorf("NewCompleteGraph is improperly colored")
	}

	t.Logf("Test: NewCompleteGraph(%d)", N)
	g = graph.NewRingGraph(N)
	colorSequential(&g, maxColor)
	if !checkColoring(&g) {
		t.Errorf("NewRingGraph is improperly colored")
	}

	t.Logf("Test: NewRandomGraph(%d, %f)", N, bf)
	g = graph.NewRandomGraph(N, bf)
	colorSequential(&g, maxColor)
	if !checkColoring(&g) {
		t.Errorf("NewRandomGraph is improperly colored")
	}
}