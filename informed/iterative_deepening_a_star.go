package search

import (
	"github.com/christat/search"
	"math"
)

// IterativeDeepeningAStar implements the IDA* algorithm.
// It is essentially a series of depth searches bounded by the minimum f value (cost + heuristic).
// if no solution is found, the bound is updated with the minimum explored f value to continue the depth search.
func IterativeDeepeningAStar(origin, target search.HeuristicState) (path map[search.State]search.State, found bool) {
	path = make(map[search.State]search.State)

	bound := origin.Heuristic()

	for {
		found, lowestTotalCost := heuristicBoundSearch(origin, target, 0, bound, path)
		if found || lowestTotalCost == math.Inf(0) {
			break
		}
		bound = lowestTotalCost
	}
	return path, found
}

func heuristicBoundSearch(vertex, target search.HeuristicState, branchCost, bound float64, path map[search.State]search.State) (found bool, newBound float64) {
	f := branchCost + vertex.Heuristic()
	if f > bound {
		return false, f
	}
	if vertex.Equals(target) {
		return true, bound
	}
	lowestTotalCost := math.Inf(0)
	for _, neighbor := range vertex.Neighbors() {
		path[neighbor] = vertex
		found, lowestBound := heuristicBoundSearch(neighbor.(search.HeuristicState), target, branchCost + vertex.Cost(neighbor), bound, path)
		if found {
			break
		}
		if lowestBound < lowestTotalCost {
			lowestTotalCost = lowestBound
		}
	}
	return false, lowestTotalCost
}

// Benchmark variant of IterativeDeepeningAStar.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
/*func BenchmarkIterativeDeepeningAStar(origin, target search.State) (path map[search.State]search.State, found bool, bench search.AlgorithmBenchmark)  {
	path = make(map[search.State]search.State)

	start := time.Now()
	var expansions uint = 0

	for i := 0; i < maxDepth; i++ {
		path, found, expansions = benchmarkDepthBoundSearch(origin, target, i, path, expansions)
		if found {
			break
		}
	}
	elapsed := time.Since(start)
	return path, found, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}

func benchmarkHeuristicBoundSearch(vertex, target search.State, bound int, path map[search.State]search.State, expansions uint) (map[search.State]search.State, bool, uint) {
	expansions++
	var found bool
	if vertex.Equals(target) {
		return path, true, expansions
	} else if bound > 0 {
		for _, neighbor := range vertex.Neighbors() {
			path[neighbor] = vertex
			path, found, expansions = benchmarkDepthBoundSearch(neighbor, target, bound - 1, path, expansions)
			if found {
				break
			}
		}
	}
	return path, found, expansions
}*/

