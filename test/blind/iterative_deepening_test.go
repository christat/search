package blind_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/blind"
	tracer "github.com/christat/search"
	"fmt"
)

func TestIterativeDeepening(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/uniform_cost_cyclic_graph_test.dot", false)
	if !ok {
		t.Errorf("Failed to parse DFS test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["1"]
	target := vertexMap["4"]
	someVertex := vertexMap["3"]
	bidirectionalAccessVertex := vertexMap["5"]

	path, found := search.IterativeDeepening(origin, target, 3)
	if !found {
		t.Errorf("Failed to find valid path in IDS test file")
	}
	benchPath, found, bench := search.BenchmarkIterativeDeepening(origin, target, 10)
	if !found {
		t.Errorf("Failed to find valid path in Benchmark_IDS test file")
	}
	fmt.Print(bench)

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "1 -> 2 -> 4"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in IDS test file.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_IDS test file.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}

	path, found = search.IterativeDeepening(origin, target, 1)
	if found {
		t.Errorf("Found path with not enough depth assigned")
	}

	path, found = search.IterativeDeepening(someVertex, bidirectionalAccessVertex, 4)
	res, _ = tracer.TraceSolutionPath(someVertex, bidirectionalAccessVertex, path)
	if !found || res.String() != "3 -> 6 -> 10 -> 5" {
		t.Errorf("Failed to find cyclic path to target")
	}
}
