package search

import (
	"github.com/christat/search"
	"time"
	"github.com/christat/gost/queue"
	"sort"
)

// Beam implements Beam Search.
// On each execution step, the most promising set of descendants (i.e. with the lowest Heuristic value) are enqueued, discarding the others (hence not keeping them in the queue).
// In practice, Beam Search behaves like a pruning-enabled Breadth-First Search, retaining at each expansion a maximum of descendants marked by beamSize.
func Beam(origin, target search.HeuristicState, beamSize uint, useNodeQueue ...bool) (path map[search.State]search.State, found bool) {
	if beamSize == 0 {
		beamSize = 1
	}
	filterSize := int(beamSize) // acts as an invalidator for negative values

	path = make(map[search.State]search.State)
	open := search.SelectQueueImplementation(useNodeQueue...)

	open.Enqueue(origin)
	for open.Size() > 0 {
		vertex := open.Dequeue().(search.HeuristicState)
		if vertex.Equals(target) {
			found = true
			break
		}
		enqueueBeamedDescendants(vertex, filterSize, open, path)
	}
	return
}

// Benchmark variant of Beam.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkBeam(origin, target search.HeuristicState, beamSize uint, useNodeQueue ...bool) (path map[search.State]search.State, found bool, bench search.AlgorithmBenchmark) {
	if beamSize == 0 {
		beamSize = 1
	}
	filterSize := int(beamSize) // acts as an invalidator for negative values
	path = make(map[search.State]search.State)
	open := search.SelectQueueImplementation(useNodeQueue...)

	var expansions uint = 0
	start := time.Now()

	open.Enqueue(origin)
	for open.Size() > 0 {
		vertex := open.Dequeue().(search.HeuristicState)
		expansions++
		if vertex.Equals(target) {
			found = true
			break
		}
		enqueueBeamedDescendants(vertex, filterSize, open, path)
	}
	elapsed := time.Since(start)

	return path, found, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}

func enqueueBeamedDescendants(vertex search.HeuristicState, filterSize int, open gost.Queue, path map[search.State]search.State) {
	neighbors := vertex.Neighbors()
	// children are sorted in ascending order according to their heuristic value
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].(search.HeuristicState).Heuristic() < neighbors[j].(search.HeuristicState).Heuristic()
	})
	var beamResult []search.State
	if len(neighbors) >= filterSize {
		beamResult = neighbors[0:filterSize]
	} else {
		beamResult = neighbors
	}
	for _, neighbor := range beamResult {
		path[neighbor] = vertex
		open.Enqueue(neighbor)
	}
}