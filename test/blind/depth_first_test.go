package blind_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/blind"
	tracer "github.com/christat/search"
)

func TestDepthFirst(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/simple_graph_test.dot", false)
	if !ok {
		t.Errorf("Failed to parse DFS test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["1"]
	target := vertexMap["9"]
	someVertex := vertexMap["3"]
	isolatedVertex := vertexMap["11"]

	path, found := search.DepthFirst(origin, target)
	if !found {
		t.Errorf("Failed to find valid path in DFS test file")
	}
	benchPath, found, bench := search.BenchmarkDepthFirst(origin, target)
	if !found {
		t.Errorf("Failed to find valid path in Benchmark_DFS test file")
	}
	if bench.TotalExpansions != 9 {
		t.Errorf("Failed to compute node expansions for Benchmark_BFS")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "1 -> 2 -> 4 -> 9"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in DFS test file.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_DFS test file.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}

	path, found = search.DepthFirst(someVertex, origin)
	if found {
		t.Errorf("Found path up the tree in directed graph")
	}

	path, found = search.DepthFirst(isolatedVertex, target)
	if found {
		t.Errorf("Found path to target from isolated origin")
	}
}
