package search

import (
	"github.com/christat/search"
	"github.com/christat/gost/queue"
	"time"
)

// AStar implements the A* algorithm.
// Even though the function returns the shortest path between two vertices in a graph,
// an optimal traversal between any two points can be built by using the return path map and func TraceSolutionPath.
// If a path exists, it will be indicated by found and the traversal cost returned.
func AStar(origin, target search.HeuristicState) (path map[search.State]search.State, found bool, cost float64) {
	path, cumulativeCost, queue, open, closed := initAStarVariables()

	queue.Enqueue(origin, 0)
	cumulativeCost[origin] = 0
	for queue.Size() > 0 {
		vertex := queue.Dequeue().(search.HeuristicState)
		closed[vertex.Name()] = true
		found = enqueueAStarNeighbors(vertex, target, queue, open, closed, cumulativeCost, path)
		if found {
			break
		}
	}
	return path, found, cumulativeCost[target]
}

// Benchmark variant of AStar.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkAStar(origin, target search.HeuristicState) (path map[search.State]search.State, found bool, cost float64, bench search.AlgorithmBenchmark) {
	path, cumulativeCost, queue, open, closed := initAStarVariables()

	start := time.Now()
	var expansions uint = 0

	queue.Enqueue(origin, 0)
	cumulativeCost[origin] = 0
	for queue.Size() > 0 {
		vertex := queue.Dequeue().(search.HeuristicState)
		closed[vertex.Name()] = true
		expansions++
		found = enqueueAStarNeighbors(vertex, target, queue, open, closed, cumulativeCost, path)
		if found {
			break
		}
	}
	elapsed := time.Since(start)
	return path, found, cumulativeCost[target], search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}

func initAStarVariables() (path map[search.State]search.State, cumulativeCost map[search.State]float64, queue *gost.MinPriorityQueue, open, closed map[string]bool) {
	path = make(map[search.State]search.State)
	cumulativeCost = make(map[search.State]float64)
	queue = new(gost.MinPriorityQueue) // Min as we need to obtain lowest costs first
	open = make(map[string]bool) // A separate open/closed map is needed to avoid re-insertion and re-inspection of vertices.
	closed = make(map[string]bool)
	return
}

func enqueueAStarNeighbors(vertex, target search.HeuristicState, queue *gost.MinPriorityQueue, open map[string]bool, closed map[string]bool, cumulativeCost map[search.State]float64, path map[search.State]search.State) (found bool) {
	if vertex.Equals(target) {
		found = true
		return
	}
	for _, neighbor := range vertex.Neighbors() {
		_, visited := closed[neighbor.Name()]
		if visited {
			continue
		}

		cumulativeVertexCost := cumulativeCost[vertex]
		cost := cumulativeVertexCost + vertex.Cost(neighbor)
		lowestCost, valueSet := cumulativeCost[neighbor]
		if !valueSet || cost < lowestCost {
			cumulativeCost[neighbor] = cost
			path[neighbor] = vertex
			priority := cost + neighbor.(search.HeuristicState).Heuristic()
			// instead of a decrease-priority operation on existent nodes within the pq, we simply re-insert them with a different priority.
			// See details at https://www.redblobgames.com/pathfinding/a-star/implementation.html#algorithm
			queue.Enqueue(neighbor, priority)
			open[neighbor.Name()] = true
		}
	}
	return
}