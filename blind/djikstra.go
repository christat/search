package search

import (
	"github.com/christat/search"
	"github.com/christat/gost/queue"
)

func Djikstra(origin, target search.WeightedState) (path map[search.State]search.State, found bool) {
	path = make(map[search.State]search.State)

	open := new(gost.PriorityQueue)
	open.Enqueue(origin, 1)

	cumulativeCost := make(map[search.State]float64)
	cumulativeCost[origin] = 0

	for open.Size() > 0 {
		vertex := open.Dequeue().(search.WeightedState)
		if vertex == target {
			found = true
			break
		}
		for _, neighbor := range vertex.Neighbors() {
			cost := cumulativeCost[vertex] + vertex.Cost(neighbor)
			lowestCost, visited := cumulativeCost[neighbor]
			if !visited || cost < lowestCost {
				cumulativeCost[neighbor] = cost
				path[neighbor] = vertex
				open.Enqueue(neighbor, cost)
			}
		}
	}
	return path, found
}
