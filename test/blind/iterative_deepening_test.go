package blind_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/blind"
	tracer "github.com/christat/search"
)

func TestIterativeDeepening(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/cyclic_graph_test.dot", false)
	if !ok {
		t.Errorf("Failed to parse IterativeDeepening test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["1"]
	target := vertexMap["4"]
	someVertex := vertexMap["3"]
	bidirectionalAccessVertex := vertexMap["5"]

	path, found := search.IterativeDeepening(origin, target, 3)
	if !found {
		t.Errorf("IDS failed to find valid path")
	}
	benchPath, found, bench := search.BenchmarkIterativeDeepening(origin, target, 10)
	if !found {
		t.Errorf("Benchmark_IDS failed to find valid path")
	}
	if bench.TotalExpansions != 4 {
		t.Errorf("Benchmark_IDS expansions calculation is incorrect")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "1 -> 2 -> 4"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in IDS.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_IDS.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}

	path, found = search.IterativeDeepening(origin, target, 1)
	if found {
		t.Errorf("IDS found path with not enough depth assigned")
	}

	path, found, bench = search.BenchmarkIterativeDeepening(someVertex, bidirectionalAccessVertex, 4)
	res, _ = tracer.TraceSolutionPath(someVertex, bidirectionalAccessVertex, path)
	if !found || res.String() != "3 -> 6 -> 10 -> 5" {
		t.Errorf("IDS failed to find cyclic path to target")
	}
}

func TestBenchmarkIterations(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/ids_test.dot", false)
	if !ok {
		t.Errorf("Failed to parse DFS test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["1"]
	target := vertexMap["6"]
	leafTarget := vertexMap["9"]

	_, found, bench := search.BenchmarkIterativeDeepening(origin, target, 10)
	if !found {
		t.Errorf("Failed to find valid path in Benchmark_IDS test file")
	}
	var expectedExpansions uint = 7
	if bench.TotalExpansions != expectedExpansions {
		t.Errorf("Failed to correctly count expansions in Benchmark_IDS test file.\nExpected: %v\nFound: %v", expectedExpansions, bench.TotalExpansions)
	}

	_, found, bench = search.BenchmarkIterativeDeepening(origin, leafTarget, 10)
	if !found {
		t.Errorf("Failed to find valid path in Benchmark_IDS test file")
	}
	expectedExpansions = 15
	if bench.TotalExpansions != expectedExpansions {
		t.Errorf("Failed to correctly count expansions in Benchmark_IDS test file.\nExpected: %v\nFound: %v", expectedExpansions, bench.TotalExpansions)
	}
}

