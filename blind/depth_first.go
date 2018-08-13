package search

import (
	"github.com/christat/search"
	"time"
)


// DepthFirst implements the Depth First Search algorithm.
// The origin and target states must fulfill the State interface.
// Optionally, a single-linked list backed stack can be enforced with useNodeStack.
func DepthFirst(origin, target search.State, useNodeStack ...bool) (path map[search.State]search.State, found bool) {
	path = make(map[search.State]search.State)

	open := selectStackImplementation(useNodeStack...)
	open.Push(origin)

	for open.Size() > 0 {
		vertex := open.Pop().(search.State)
		found = checkVertexAndPushNeighbors(vertex, target, open, path)
		if found {
			break
		}
	}
	return path, found
}

// Benchmark variant of DepthFirst.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkDepthFirst(origin, target search.State, useNodeStack ...bool) (path map[search.State]search.State, found bool, bench search.AlgorithmBenchmark) {
	path = make(map[search.State]search.State)
	open := selectStackImplementation(useNodeStack...)
	start := time.Now()
	var expansions uint = 0

	open.Push(origin)
	for open.Size() > 0 {
		vertex := open.Pop().(search.State)
		expansions++
		found = checkVertexAndPushNeighbors(vertex, target, open, path)
		if found {
			break
		}
	}
	elapsed := time.Since(start)
	return path, found, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}
