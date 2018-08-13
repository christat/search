package blind_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/blind"
	tracer "github.com/christat/search"
)

func TestDjikstra(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/weighted_cyclic_graph_test.dot", false)
	graph.CostKey = "w"
	if !ok {
		t.Errorf("Failed to parse Djikstra test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["a"]
	target := vertexMap["e"]

	path, found, cost := search.Djikstra(origin, target)
	if !found {
		t.Errorf("Djikstra failed to find valid path")
	}
	if cost != 9 {
		t.Errorf("Djikstra cost computation is incorrect")
	}

	benchPath, found, cost, bench := search.BenchmarkDjikstra(origin, target)
	if !found {
		t.Errorf("Benchmark_Djikstra failed to find valid path")
	}
	if cost != 9 {
		t.Errorf("Benchmark_Djikstra cost computation is incorrect")
	}
	if bench.TotalExpansions != 5 {
		t.Errorf("Benchmark_Djikstra expansions calculation is incorrect")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "a -> c -> d -> e"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Djikstra.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_Djikstra.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}
}