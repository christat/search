package search

import (
	"github.com/christat/search"
	"github.com/christat/gost/queue"
)

// Djikstra implements the well known algorithm of said name.
// Even though the function returns the shortest path between two vertices in a graph,
// an optimal traversal between any two points can be built by using the return path map and func TraceSolutionPath.
// If a path exists, it will be indicated by found and the traversal cost returned.
func Djikstra(origin, target search.WeightedState) (path map[search.State]search.State, found bool, cost float64) {
	path, found, cost, _ = search.BestFirst(origin.(search.HeuristicState), target.(search.HeuristicState), enqueueDjikstraNeighbors)
	return
}

// Benchmark variant of Djikstra.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkDjikstra(origin, target search.WeightedState) (path map[search.State]search.State, found bool, cost float64, bench search.AlgorithmBenchmark) {
	path, found, cost, bench, _ = search.BenchmarkBestFirst(origin.(search.HeuristicState), target.(search.HeuristicState), enqueueDjikstraNeighbors)
	return
}

func enqueueDjikstraNeighbors(vertex search.HeuristicState, cost float64, queue *gost.MinPriorityQueue, open map[string]bool) {
	_, queued := open[vertex.Name()]
	if !queued {
		queue.Enqueue(vertex, cost)
		open[vertex.Name()] = true
	}
}