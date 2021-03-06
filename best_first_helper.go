package search

import (
	"github.com/christat/gost/queue"
	"time"
	"fmt"
)

// Best First Search underpins several algorithms, such as Greedy BFS or A*.
// The main difference comes in the enqueuing logic, which is specific to the algorithm itself.
func BestFirst(origin, target HeuristicState, callback BFSEnqueuingCallback) (path map[State]State, found bool, cost float64, err error) {
	path, cumulativeCost, queue, open, closed := initBestFirstVariables()

	if callback == nil {
		return path, found, 0, fmt.Errorf("enqueuing Callback not provided! Best First Search cannot be executed")
	}

	queue.Enqueue(origin, 0)
	cumulativeCost[origin] = 0
	for queue.Size() > 0 {
		vertex := queue.Dequeue().(HeuristicState)
		closed[vertex.Name()] = true
		found = enqueueBestFirstNeighbors(vertex, target, queue, open, closed, cumulativeCost, path, callback)
		if found {
			break
		}
	}
	return path, found, cumulativeCost[target], nil
}

func BenchmarkBestFirst(origin, target HeuristicState, callback BFSEnqueuingCallback) (path map[State]State, found bool, cost float64, bench AlgorithmBenchmark, err error) {
	path, cumulativeCost, queue, open, closed := initBestFirstVariables()

	if callback == nil {
		return path, found, 0, AlgorithmBenchmark{}, fmt.Errorf("enqueuing Callback not provided! Best First Search cannot be executed")
	}

	start := time.Now()
	var expansions uint = 0

	queue.Enqueue(origin, 0)
	cumulativeCost[origin] = 0
	for queue.Size() > 0 {
		vertex := queue.Dequeue().(HeuristicState)
		closed[vertex.Name()] = true
		expansions++
		found = enqueueBestFirstNeighbors(vertex, target, queue, open, closed, cumulativeCost, path, callback)
		if found {
			break
		}
	}
	elapsed := time.Since(start)
	return path, found, cumulativeCost[target], AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}, nil
}

func initBestFirstVariables() (path map[State]State, cumulativeCost map[State]float64, queue *gost.MinPriorityQueue, open, closed map[string]bool) {
	path = make(map[State]State)
	cumulativeCost = make(map[State]float64)
	queue = new(gost.MinPriorityQueue) // Min as we need to obtain lowest costs first
	open = make(map[string]bool)
	closed = make(map[string]bool)
	return
}

func enqueueBestFirstNeighbors(vertex, target HeuristicState, queue *gost.MinPriorityQueue, open map[string]bool, closed map[string]bool, cumulativeCost map[State]float64, path map[State]State, callback BFSEnqueuingCallback) (found bool) {
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
			if callback != nil {
				callback(neighbor.(HeuristicState), cost, queue, open)
			}
		}
	}
	return
}

// Each algorithm decides how to enqueue its nodes. The callback should provide any neccesary parameters.
type BFSEnqueuingCallback func(vertex HeuristicState, cost float64, queue *gost.MinPriorityQueue, open map[string]bool)
