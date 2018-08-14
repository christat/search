package informed_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/informed"
	tracer "github.com/christat/search"
)

func TestHillClimbing(t *testing.T) {
	// Hill Climbing is a special case of Beam Search. Both test files should behave exactly equally in test
	// with the precondition that Beam Search's beamSize is set to 1.
	ok, graph := dot.ParseFile("../test_dot_files/romania_manhattan_heuristic.dot", false)
	graph.CostKey = "distance"
	graph.HeuristicKey = "manhattan"
	if !ok {
		t.Errorf("Failed to parse HillClimbing test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["Arad"]
	target := vertexMap["Bucharest"]

	path, found := search.HillClimbing(origin, target)
	if !found {
		t.Errorf("HillClimbing failed to find valid path")
	}

	benchPath, found, bench := search.BenchmarkHillClimbing(origin, target)
	if !found {
		t.Errorf("Benchmark_HillClimbing failed to find valid path")
	}
	if bench.TotalExpansions != 4 {
		t.Errorf("Benchmark_HillClimbing expansions calculation is incorrect")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "Arad -> Sibiu -> Fagaras -> Bucharest"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in HillClimbing.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_HillClimbing.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}
}