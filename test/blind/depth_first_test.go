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
		t.Errorf("DFS failed to find valid path")
	}
	benchPath, found, bench := search.BenchmarkDepthFirst(origin, target)
	if !found {
		t.Errorf("Benchmark_DFS failed to find valid path")
	}
	if bench.TotalExpansions != 9 {
		t.Errorf("Benchmark_BFS expansions calculation incorrect")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "1 -> 2 -> 4 -> 9"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in DFS.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_DFS.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}

	path, found = search.DepthFirst(someVertex, origin)
	if found {
		t.Errorf("DFS found path up the tree in directed graph")
	}

	path, found = search.DepthFirst(isolatedVertex, target)
	if found {
		t.Errorf("DFS found path to target from isolated origin")
	}
}
