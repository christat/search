package search

import (
	"github.com/christat/search"
	"time"
)

// IterativeDeepening implements recursive IDS.
// It will look for optimal solutions reaching target from origin.
// The depth bound is slowly increased until reaching maxDepth.
func IterativeDeepening(origin, target search.State, maxDepth int) (path map[search.State]search.State, found bool) {
	path = make(map[search.State]search.State)

	for i := 0; i <= maxDepth; i++ {
		path, found = depthBoundSearch(origin, target, i, path)
		if found {
			break
		}
	}
	return path, found
}

// Recursive DFS search with an applied depth bound
func depthBoundSearch(vertex, target search.State, bound int, path map[search.State]search.State) (map[search.State]search.State, bool) {
	var found bool
	if vertex.Equals(target) {
		return path, true
	} else if bound > 0 {
		for _, neighbor := range vertex.Neighbors() {
			path[neighbor] = vertex
			path, found = depthBoundSearch(neighbor, target, bound - 1, path)
			if found {
				break
			}
		}
	}
	return path, found
}

// Benchmark variant of IterativeDeepening.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkIterativeDeepening(origin, target search.State, maxDepth int) (path map[search.State]search.State, found bool, bench search.AlgorithmBenchmark)  {
	path = make(map[search.State]search.State)

	start := time.Now()
	var expansions uint = 0

	for i := 0; i < maxDepth; i++ {
		found, expansions = benchmarkDepthBoundSearch(origin, target, i, path, expansions)
		if found {
			break
		}
	}
	elapsed := time.Since(start)
	return path, found, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}

func benchmarkDepthBoundSearch(vertex, target search.State, bound int, path map[search.State]search.State, expansions uint) (bool, uint) {
	expansions++
	var found bool
	if vertex.Equals(target) {
		return true, expansions
	} else if bound > 0 {
		for _, neighbor := range vertex.Neighbors() {
			path[neighbor] = vertex
			found, expansions = benchmarkDepthBoundSearch(neighbor, target, bound - 1, path, expansions)
			if found {
				break
			}
		}
	}
	return found, expansions
}
