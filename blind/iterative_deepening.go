package search

import (
	"github.com/christat/search"
	"time"
)

func IterativeDeepening(origin, target search.State, maxDepth int) (path map[search.State]search.State, found bool) {
	path = make(map[search.State]search.State)

	for i := 0; i <= maxDepth; i++ {
		path, found = DepthBoundSearch(origin, target, i, path)
		if found {
			break
		}
	}
	return path, found
}

func DepthBoundSearch(vertex, target search.State, bound int, path map[search.State]search.State) (map[search.State]search.State, bool) {
	var found bool
	if vertex.Equals(target) {
		return path, true
	} else if bound > 0 {
		for _, neighbor := range vertex.Neighbors() {
			path[neighbor] = vertex
			path, found = DepthBoundSearch(neighbor, target, bound - 1, path)
			if found {
				break
			}
		}
	}
	return path, found
}

func BenchmarkIterativeDeepening(origin, target search.State, maxDepth int) (path map[search.State]search.State, found bool, bench search.AlgorithmBenchmark)  {
	path = make(map[search.State]search.State)
	start := time.Now()
	var expansions uint = 0

	for i := 0; i < maxDepth; i++ {
		path, found, expansions = BenchmarkDepthBoundSearch(origin, target, i, path, expansions)
		if found {
			break
		}
	}
	elapsed := time.Since(start)
	return path, found, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}

func BenchmarkDepthBoundSearch(vertex, target search.State, bound int, path map[search.State]search.State, expansions uint) (map[search.State]search.State, bool, uint) {
	expansions++
	var found bool
	if vertex.Equals(target) {
		return path, true, expansions
	} else if bound > 0 {
		for _, neighbor := range vertex.Neighbors() {
			path[neighbor] = vertex
			path, found, expansions = BenchmarkDepthBoundSearch(neighbor, target, bound - 1, path, expansions)
			if found {
				break
			}
		}
	}
	return path, found, expansions
}
