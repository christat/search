package informed_test

import (
	"testing"
	"github.com/christat/dot"
	tracer "github.com/christat/search"
	"github.com/christat/search/informed"
)

func TestIterativeDeepeningAStar(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/heuristic_graph_test.dot", false)
	graph.CostKey = "w"
	graph.HeuristicKey = "h"
	if !ok {
		t.Errorf("Failed to parse IDA* test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["a"]
	target := vertexMap["f"]

	path, found, cost := search.IterativeDeepeningAStar(origin, target)
	if !found {
		t.Errorf("IDA* failed to find valid path")
	}
	if cost != 9 {
		t.Errorf("IDA* failed to compute costs correctly")
	}
	benchPath, found, cost, bench := search.BenchmarkIterativeDeepeningAStar(origin, target)
	if !found {
		t.Errorf("Benchmark_IDA* failed to find valid path")
	}
	if bench.TotalExpansions != 16 {
		t.Errorf("Benchmark_IDA* expansions calculation is incorrect")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "a -> b -> e -> f"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in IDA*.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_IDA*.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}
}

