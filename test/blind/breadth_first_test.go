package blind_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/blind"
	tracer "github.com/christat/search"
	"fmt"
)

func TestBreadthFirst(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/uniform_cost_graph_test.dot", false)
	if !ok {
		t.Errorf("Failed to parse BFS test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["1"]
	target := vertexMap["10"]
	someVertex := vertexMap["3"]

	path, found := search.BreadthFirst(origin, target)
	if !found {
		t.Errorf("Failed to find valid path in BFS test file")
	}
	benchPath, found, bench := search.BenchmarkBreadthFirst(origin, target)
	if !found {
		t.Errorf("Failed to find valid path in BFS test file")
	}

	fmt.Print(bench)

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "1 -> 2 -> 5 -> 10"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in BFS test file.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_BFS test file.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}

	path, found = search.BreadthFirst(someVertex, origin)
	if found {
		t.Errorf("Found path up the tree in directed graph")
	}

	path, found = search.BreadthFirst(someVertex, target)
	if found {
		t.Errorf("Found path to target from isolated origin")
	}
}