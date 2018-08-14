package informed_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/informed"
	base "github.com/christat/search"
	tracer "github.com/christat/search"
)

func TestBeam(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/romania_manhattan_heuristic.dot", false)
	graph.CostKey = "distance"
	graph.HeuristicKey = "manhattan"
	if !ok {
		t.Errorf("Failed to parse Beam test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["Arad"]
	target := vertexMap["Bucharest"]

	path, found := search.Beam(origin, target, 1)
	if !found {
		t.Errorf("Beam failed to find valid path")
	}

	benchPath, found, bench := search.BenchmarkBeam(origin, target, 1)
	if !found {
		t.Errorf("Benchmark_Beam failed to find valid path")
	}
	if bench.TotalExpansions != 4 {
		t.Errorf("Benchmark_Beam expansions calculation is incorrect")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "Arad -> Sibiu -> Fagaras -> Bucharest"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Beam.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_Beam.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}
}

func TestBeamBranching(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/beam_test.dot", false)
	graph.HeuristicKey = "h"
	if !ok {
		t.Errorf("Failed to parse Beam Branching test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["a"]
	target := vertexMap["m"]

	expectedResults := []expectedBeamResults{
		{false, "", base.AlgorithmBenchmark{ElapsedTime: 0, TotalExpansions: 3}},
		{true, "a -> d -> i -> m", base.AlgorithmBenchmark{ElapsedTime: 0, TotalExpansions: 9}},
		{true, "a -> d -> i -> m", base.AlgorithmBenchmark{ElapsedTime: 0, TotalExpansions: 11}},
	}
	for i := 1; i < 4; i++ {
		executeBeamBranchingTest(origin, target, uint(i), expectedResults[i - 1], t)
	}

}

type expectedBeamResults struct {
	found     bool
	path      string
	benchmark base.AlgorithmBenchmark
}

func executeBeamBranchingTest(origin, target base.HeuristicState, branching uint, expectedResults expectedBeamResults, t *testing.T) {
	path, found := search.Beam(origin, target, branching)
	if found != expectedResults.found {
		t.Errorf("Beam branching solution found expected: %v, actual: %v\n", expectedResults.found, found)
	}

	benchPath, found, bench := search.BenchmarkBeam(origin, target, branching)
	if found != expectedResults.found {
		t.Errorf("Benchmark_Beam branching solution found expected: %v, actual: %v\n", expectedResults.found, found)
	}
	if bench.TotalExpansions != expectedResults.benchmark.TotalExpansions {
		t.Errorf("Benchmark_Beam expansions calculation expected: %v, actual: %v\n", expectedResults.benchmark.TotalExpansions, bench.TotalExpansions)
	}

	if expectedResults.found {
		res, _ := tracer.TraceSolutionPath(origin, target, path)
		benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

		expectedSolution := expectedResults.path
		foundSolution := res.String()
		benchFoundSolution := benchRes.String()
		if foundSolution != expectedSolution {
			t.Errorf("Failed to find correct solution path in Beam.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
		}
		if benchFoundSolution != expectedSolution {
			t.Errorf("Failed to find correct solution path in Benchmark_Beam.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
		}
	}

}