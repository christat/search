package search

import (
	"github.com/christat/search"
	"math"
	"time"
)

// HillClimbing implements a heuristic-based Hill Climbing algorithm.
// On each execution step, the most promising descendant (i.e. with the lowest Heuristic value) is further expanded, discarding the others (hence not keeping them in any collection).
func HillClimbing(origin, target search.HeuristicState) (path map[search.State]search.State, found bool) {
	path = make(map[search.State]search.State)
	current := origin
	for {
		if current.Equals(target) {
			found = true
			break
		}
		chosenDescendant := chooseMostPromisingDescendant(current)
		if chosenDescendant == nil {
			break
		}
		path[chosenDescendant] = current
		current = chosenDescendant
	}
	return
}

// Benchmark variant of HillClimbing.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkHillClimbing(origin, target search.HeuristicState) (path map[search.State]search.State, found bool, bench search.AlgorithmBenchmark) {
	path = make(map[search.State]search.State)
	current := origin

	var expansions uint = 1 // current = origin is the first expansion
	start := time.Now()

	for {
		if current.Equals(target) {
			found = true
			break
		}
		chosenDescendant := chooseMostPromisingDescendant(current)
		if chosenDescendant == nil {
			break
		}
		expansions++
		path[chosenDescendant] = current
		current = chosenDescendant
	}
	elapsed := time.Since(start)

	return path, found, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}

func chooseMostPromisingDescendant(vertex search.HeuristicState) search.HeuristicState {
	mostPromisingHeuristic := math.Inf(0)
	var mostPromisingDescendant search.HeuristicState
	for _, descendant := range vertex.Neighbors() {
		heuristicDescendant := descendant.(search.HeuristicState)
		heuristic := heuristicDescendant.Heuristic()
		if heuristic < mostPromisingHeuristic {
			mostPromisingDescendant = heuristicDescendant
			mostPromisingHeuristic = heuristic
		}
	}
	return mostPromisingDescendant
}