package search

import (
	"github.com/christat/search"
	"github.com/christat/gost/queue"
	"time"
)

func Djikstra(origin, target search.WeightedState) (path map[search.State]search.State, found bool, cost float64) {
	path = make(map[search.State]search.State)
	cumulativeCost := make(map[search.State]float64)
	queue := new(gost.MinPriorityQueue) // Min as we need to obtain lowest costs first
	open := make(map[string]bool) // A separate open/closed map is needed to avoid re-insertion and re-inspection of vertices.
	closed := make(map[string]bool)

	queue.Enqueue(origin, 0)
	cumulativeCost[origin] = 0
	for queue.Size() > 0 {
		vertex := queue.Dequeue().(search.WeightedState)
		closed[vertex.Name()] = true
		found = enqueueUnvisitedLowerCostNeighbors(vertex, target, queue, open, closed, cumulativeCost, path)
		if found {
			break
		}
	}
	return path, found, cumulativeCost[target]
}

func BenchmarkDjikstra(origin, target search.WeightedState) (path map[search.State]search.State, found bool, cost float64, bench search.AlgorithmBenchmark) {
	path = make(map[search.State]search.State)
	cumulativeCost := make(map[search.State]float64)
	queue := new(gost.MinPriorityQueue) // Min as we need to obtain lowest costs first
	open := make(map[string]bool) // A separate open/closed map is needed to avoid re-insertion and re-inspection of vertices.
	closed := make(map[string]bool)

	start := time.Now()
	var expansions uint = 0

	queue.Enqueue(origin, 0)
	cumulativeCost[origin] = 0
	for queue.Size() > 0 {
		vertex := queue.Dequeue().(search.WeightedState)
		closed[vertex.Name()] = true
		expansions++
		found = enqueueUnvisitedLowerCostNeighbors(vertex, target, queue, open, closed, cumulativeCost, path)
		if found {
			break
		}
	}
	elapsed := time.Since(start)
	return path, found, cumulativeCost[target], search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}

func enqueueUnvisitedLowerCostNeighbors(vertex, target search.WeightedState, queue *gost.MinPriorityQueue, open map[string]bool, closed map[string]bool, cumulativeCost map[search.State]float64, path map[search.State]search.State) (found bool) {
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
		adjacencyCost := vertex.Cost(neighbor)
		cost := cumulativeVertexCost + adjacencyCost
		lowestCost, valueSet := cumulativeCost[neighbor]
		if !valueSet || cost < lowestCost {
			cumulativeCost[neighbor] = cost
			path[neighbor] = vertex
			_, queued := open[neighbor.Name()]
			if !queued {
				queue.Enqueue(neighbor, cost)
				open[neighbor.Name()] = true
			}
		}
	}
	return
}