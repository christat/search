package search

import (
	"github.com/christat/search"
	"math"
	"time"
)

// IterativeDeepeningAStar implements the IDA* algorithm.
// It performs a series of depth searches bounded by the minimum f value (cost + heuristic).
// if no solution is found, the bound is updated with the minimum explored f value to continue the depth search.
func IterativeDeepeningAStar(origin, target search.HeuristicState) (path map[search.State]search.State, found bool, cost float64) {
	path = make(map[search.State]search.State)
	bound := origin.Heuristic()
	lowestCost := math.Inf(0)
	for {
		found, lowestCost = heuristicBoundSearch(origin, target, 0, bound, path)
		if found || lowestCost == math.Inf(0) {
			break
		}
		bound = lowestCost
	}
	return path, found, lowestCost
}

// Unfortunately the logic in IDA* seems to differ greatly enough from standard IDS to make the usage of a common helper function unfeasible (practically speaking).
func heuristicBoundSearch(vertex, target search.HeuristicState, branchCost, bound float64, path map[search.State]search.State) (found bool, lowestCost float64) {
	cost := branchCost + vertex.Heuristic()
	if cost > bound {
		return false, cost
	}
	if vertex.Equals(target) {
		return true, bound
	}
	lowestTotalCost := math.Inf(0)
	for _, neighbor := range vertex.Neighbors() {
		path[neighbor] = vertex
		cost = branchCost + vertex.Cost(neighbor)
		found, lowestBound := heuristicBoundSearch(neighbor.(search.HeuristicState), target, cost, bound, path)
		if found {
			return true, lowestBound
		}
		if lowestBound < lowestTotalCost {
			lowestTotalCost = lowestBound
		}
	}
	return false, lowestTotalCost
}

// Benchmark variant of IterativeDeepeningAStar.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkIterativeDeepeningAStar(origin, target search.HeuristicState) (path map[search.State]search.State, found bool, cost float64, bench search.AlgorithmBenchmark) {
	path = make(map[search.State]search.State)
	bound := origin.Heuristic()
	lowestCost := math.Inf(0)

	start := time.Now()
	var expansions uint = 0

	for {
		found, lowestCost = benchmarkHeuristicBoundSearch(origin, target, 0, bound, path, &expansions)
		if found || lowestCost == math.Inf(0) {
			break
		}
		bound = lowestCost
	}
	elapsed := time.Since(start)
	return path, found, lowestCost, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}

func benchmarkHeuristicBoundSearch(vertex, target search.HeuristicState, branchCost, bound float64, path map[search.State]search.State, expansions *uint) (bool, float64) {
	cost := branchCost + vertex.Heuristic()
	if cost > bound {
		return false, cost
	}
	if vertex.Equals(target) {
		return true, bound
	}
	lowestTotalCost := math.Inf(0)
	for _, neighbor := range vertex.Neighbors() {
		*expansions = *expansions + 1
		path[neighbor] = vertex
		cost = branchCost + vertex.Cost(neighbor)
		found, lowestBound := benchmarkHeuristicBoundSearch(neighbor.(search.HeuristicState), target, cost, bound, path, expansions)
		if found {
			return true, lowestBound
		}
		if lowestBound < lowestTotalCost {
			lowestTotalCost = lowestBound
		}
	}
	return false, lowestTotalCost
}

