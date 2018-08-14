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

	path, found := search.IterativeDeepeningAStar(origin, target)
	if !found {
		t.Errorf("IDA* failed to find valid path")
	}
	/*benchPath, found, bench := search.BenchmarkIterativeDeepening(origin, target, 10)
	if !found {
		t.Errorf("Benchmark_IDS failed to find valid path")
	}
	if bench.TotalExpansions != 7 {
		t.Errorf("Benchmark_IDS expansions calculation is incorrect")
	}*/

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	//benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "a -> b -> e -> f"
	foundSolution := res.String()
	//benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in IDA*.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	/*if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_IDS.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}*/
}

