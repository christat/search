package blind_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/blind"
	tracer "github.com/christat/search"
)

func TestBreadthFirst(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/simple_graph_test.dot", false)
	if !ok {
		t.Errorf("Failed to parse BFS test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["1"]
	target := vertexMap["10"]
	someVertex := vertexMap["3"]

	path, found := search.BreadthFirst(origin, target)
	if !found {
		t.Errorf("BFS failed to find valid path")
	}
	benchPath, found, bench := search.BenchmarkBreadthFirst(origin, target)
	if !found {
		t.Errorf("Benchmark_BFS failed to find valid path")
	}
	if bench.TotalExpansions != 10 {
		t.Errorf("Benchmark_BFS expansions calculation is incorrect")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "1 -> 2 -> 5 -> 10"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in BFS.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_BFS.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}

	path, found = search.BreadthFirst(someVertex, origin)
	if found {
		t.Errorf("BFS Found path up the tree in directed graph")
	}

	path, found = search.BreadthFirst(someVertex, target)
	if found {
		t.Errorf("BFS Found path to target from isolated origin")
	}
}