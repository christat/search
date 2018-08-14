package informed_test

import (
	"testing"
	"github.com/christat/dot"
	"github.com/christat/search/informed"
	tracer "github.com/christat/search"
)

func TestGreedyBestFirst(t *testing.T) {
	ok, graph := dot.ParseFile("../test_dot_files/romania_manhattan_heuristic.dot", false)
	graph.CostKey = "distance"
	graph.HeuristicKey = "manhattan"
	if !ok {
		t.Errorf("Failed to parse GreedyBestFirst test file")
	}
	vertexMap := graph.VertexMap()
	origin := vertexMap["Arad"]
	target := vertexMap["Bucharest"]

	path, found, cost := search.GreedyBestFirst(origin, target)
	if !found {
		t.Errorf("GreedyBestFirst failed to find valid path")
	}
	if cost != 450 {
		t.Errorf("GreedyBestFirst cost computation is incorrect")
	}

	benchPath, found, cost, bench := search.BenchmarkGreedyBestFirst(origin, target)
	if !found {
		t.Errorf("Benchmark_GreedyBestFirst failed to find valid path")
	}
	if cost != 450 {
		t.Errorf("Benchmark_GreedyBestFirst cost computation is incorrect")
	}
	if bench.TotalExpansions != 4 {
		t.Errorf("Benchmark_GreedyBestFirst expansions calculation is incorrect")
	}

	res, _ := tracer.TraceSolutionPath(origin, target, path)
	benchRes, _ := tracer.TraceSolutionPath(origin, target, benchPath)

	expectedSolution := "Arad -> Sibiu -> Fagaras -> Bucharest"
	foundSolution := res.String()
	benchFoundSolution := benchRes.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in GreedyBestFirst.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if benchFoundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path in Benchmark_GreedyBestFirst.\nExpected: %v\nFound: %v", expectedSolution, benchFoundSolution)
	}

	origin = vertexMap["Dobreta"]
	target = vertexMap["Fagaras"]
	path, found, cost = search.GreedyBestFirst(origin, target)
	res, _ = tracer.TraceSolutionPath(origin, target, path)
	expectedSolution = "Dobreta -> Craiova -> RimnicuValcea -> Pitesti -> Bucharest -> Fagaras"
	foundSolution = res.String()
	if foundSolution != expectedSolution {
		t.Errorf("Failed to find correct solution path with unadmissible Heuristic in GreedyBestFirst.\nExpected: %v\nFound: %v", expectedSolution, foundSolution)
	}
	if cost != 677 {
		t.Errorf("Benchmark_GreedyBestFirst with unadmissible Heuristic cost computation is incorrect")
	}
}