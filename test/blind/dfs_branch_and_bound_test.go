package blind_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/blind"
	tracer "github.com/christat/search"
)

func TestBranchAndBound(t *testing.T) {
	fileName := "weighted_graph_test.dot"
	ok, graph := dot.ParseFile("../test_dot_files/" + fileName, false)
	graph.CostKey = "weight"
	if !ok {
		t.Errorf("Failed to parse file %v", fileName)
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["q1"]
	target := vertexMap["q11"]

	path, found, cost := search.DFSBranchAndBound(origin, target, 0)
	if !found {
		t.Errorf("DFSBnB failed to find valid path")
	}
	if cost != 11 {
		t.Errorf("DFSBnB cost computation is incorrect")
	}

	/*benchPath, found, cost, bench := search.BenchmarkDFSBranchAndBound(origin, target)
	if !found {
		t.Errorf("Failed to find valid path in Benchmark_Djikstra test file")
	}
	if cost <= 1 {
		t.Errorf("Cost computation is incorrect for Benchmark_Djikstra test file")
	}
	if bench.TotalExpansions <= 1 {
		t.Errorf("Failed to correctly compute Benchmark_Djikstra node expansions")
	}*/

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	//benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "q1 -> q3 -> q7 -> q6 -> q11"
	foundSolution := res.String()
	//benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in DFSBnB test.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	/*if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_Djikstra test file.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}*/
}