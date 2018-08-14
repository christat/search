package search

import (
	"github.com/christat/search"
	"time"
)

// BreadthFirst implements the Breadth First Search algorithm.
// The origin and target states must fulfill the State interface.
// Optionally, a single-linked list backed stack can be enforced with useNodeQueue.
func BreadthFirst(origin, target search.State, useNodeQueue ...bool) (path map[search.State]search.State, found bool) {
	path = make(map[search.State]search.State)
	open := search.SelectQueueImplementation(useNodeQueue...)
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

// Benchmark variant of BreadthFirst.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkBreadthFirst(origin, target search.State, useNodeQueue ...bool) (path map[search.State]search.State, found bool, bench search.AlgorithmBenchmark) {
	path = make(map[search.State]search.State)
	open := search.SelectQueueImplementation(useNodeQueue...)
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
