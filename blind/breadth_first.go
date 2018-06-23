package search

import (
	"github.com/christat/search"
	"time"
)

func BreadthFirst(origin, target search.State, useNodeQueue ...bool) (path map[search.State]search.State, found bool) {
	path = make(map[search.State]search.State)
	open := selectQueueImplementation(useNodeQueue...)
	open.Enqueue(origin)

	for open.Size() > 0 {
		vertex := open.Dequeue().(search.State)
		found = checkVertexAndEnqueueNeighbors(vertex, target, open, path)
		if found {
			break
		}
	}
	return path, found
}

func BenchmarkBreadthFirst(origin, target search.State, useNodeQueue ...bool) (path map[search.State]search.State, found bool, bench search.AlgorithmBenchmark) {
	path = make(map[search.State]search.State)
	open := selectQueueImplementation(useNodeQueue...)
	start := time.Now()
	var expansions uint = 0

	open.Enqueue(origin)
	for open.Size() > 0 {
		vertex := open.Dequeue().(search.State)
		expansions++
		found = checkVertexAndEnqueueNeighbors(vertex, target, open, path)
		if found {
			break
		}
	}
	elapsed := time.Since(start)
	return path, found, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}
