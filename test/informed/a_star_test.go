package informed_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/informed"
	tracer "github.com/christat/search"
)

func TestAStar(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/heuristic_graph_test.dot", false)
	graph.CostKey = "w"
	graph.HeuristicKey = "h"
	if !ok {
		t.Errorf("Failed to parse AStar test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["a"]
	target := vertexMap["f"]

	path, found, cost := search.AStar(origin, target)
	if !found {
		t.Errorf("AStar failed to find valid path")
	}
	if cost != 9 {
		t.Errorf("AStar cost computation is incorrect")
	}

	benchPath, found, cost, bench := search.BenchmarkAStar(origin, target)
	if !found {
		t.Errorf("Benchmark_AStar failed to find valid path")
	}
	if cost != 9 {
		t.Errorf("Benchmark_AStar cost computation is incorrect")
	}
	if bench.TotalExpansions != 6 {
		t.Errorf("Benchmark_AStar expansions calculation is incorrect")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "a -> b -> e -> f"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in AStar.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_AStar.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}
}