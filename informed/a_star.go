package search

import (
	"github.com/christat/search"
	"github.com/christat/gost/queue"
)

// AStar implements the A* algorithm.
// Even though the function returns the shortest path between two vertices in a graph,
// an optimal traversal between any two points can be built by using the return path map and func TraceSolutionPath.
// If a path exists, it will be indicated by found and the traversal cost returned.
func AStar(origin, target search.HeuristicState) (path map[search.State]search.State, found bool, cost float64) {
	return search.BestFirst(origin, target, enqueueAStarNeighbors)
}

// Benchmark variant of AStar.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkAStar(origin, target search.HeuristicState) (path map[search.State]search.State, found bool, cost float64, bench search.AlgorithmBenchmark) {
	return search.BenchmarkBestFirst(origin, target, enqueueAStarNeighbors)
}

func enqueueAStarNeighbors(vertex search.HeuristicState, cost float64, queue *gost.MinPriorityQueue, open map[string]bool) {
	// instead of a decrease-priority operation on existent nodes within the pq, we simply re-insert them with a different priority.
	// See details at https://www.redblobgames.com/pathfinding/a-star/implementation.html#algorithm
	queue.Enqueue(vertex, cost + vertex.Heuristic())
	open[vertex.Name()] = true
}