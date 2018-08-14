package search

import (
	"github.com/christat/search"
	"math"
	"time"
)

// BranchAndBound performs depth search iteratively. An upper bound is set every time a solution is found,
// pruning costlier descendants and stopping once no better solution was found. Because of its nature (minimization of positive costs)
// it is not expected to work correctly with negative costs. Maximization problems should be redefined accordingly.
// Paramter bound can be left as default float64 (0); the algorithm will assume an initial bound of plus infinity.
func DFSBranchAndBound(origin, target search.WeightedState, bound float64) (path map[search.State]search.State, found bool, cost float64) {
	path, bound ,cost = initBnBVariables(bound)

	var solutionPath map[search.State]search.State
	// The expansion order in root dictates the order of branch expansions;
	// hence the algorithm follows a left-to-right DFS "scanning" pattern.
	for _, neighbor := range origin.Neighbors() {
		// because of Go's inflexible type system, neighbor must be coerced to allow access to cost function
		solutionFound, currentCost := costBoundSearch(origin, origin, neighbor.(search.HeuristicState), target, cost, bound, path)
		if solutionFound && currentCost < bound {
			found = true
			bound = currentCost
			solutionPath = copyPath(path)
		}
	}
	return solutionPath, found, bound
}

func costBoundSearch(origin, from, to, target search.WeightedState, branchCost, bound float64, path map[search.State]search.State) (found bool, cost float64) {
	expansionCost := from.Cost(to)
	if branchCost + expansionCost < bound {
		path[to] = from
		cost = branchCost + expansionCost
		if to.Equals(target) {
			bound = cost
			found = true
		} else {
			oldCost := cost
			for _, neighbor := range to.Neighbors() {
				// because of Go's inflexible type system, neighbor must be coerced to allow access to cost/heuristic functions
				solutionBranch, branchCost := costBoundSearch(origin, to, neighbor.(search.HeuristicState), target, oldCost, bound, path)
				if solutionBranch && branchCost < bound {
					found = true
					bound = branchCost
					cost = bound
				}
			}
		}
	}
	return
}

// Benchmark variant of DFSBranchAndBound.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
func BenchmarkDFSBranchAndBound(origin, target search.WeightedState, bound float64) (path map[search.State]search.State, found bool, cost float64, bench search.AlgorithmBenchmark) {
	path, bound ,cost = initBnBVariables(bound)

	start := time.Now()
	var expansions uint = 0

	var solutionPath map[search.State]search.State
	// The expansion order in root dictates the order of branch expansions;
	// hence the algorithm follows a left-to-right DFS "scanning" pattern.
	for _, neighbor := range origin.Neighbors() {
		// because of Go's inflexible type system, neighbor must be coerced to allow access to cost/heuristic functions
		solutionFound, currentCost := benchmarkCostBoundSearch(origin, origin, neighbor.(search.HeuristicState), target, cost, bound, path, &expansions)
		if solutionFound && currentCost < bound {
			found = true
			bound = currentCost
			solutionPath = copyPath(path)
		}
	}
	elapsed := time.Since(start)
	return solutionPath, found, bound, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}

func benchmarkCostBoundSearch(origin, from, to, target search.WeightedState, branchCost, bound float64, path map[search.State]search.State, expansions *uint) (found bool, cost float64) {
	expansionCost := from.Cost(to)
	if branchCost + expansionCost < bound {
		path[to] = from
		cost = branchCost + expansionCost
		if to.Equals(target) {
			bound = cost
			found = true
		} else {
			oldCost := cost
			for _, neighbor := range to.Neighbors() {
				*expansions++
				// because of Go's inflexible type system, neighbor must be coerced to allow access to cost/heuristic functions
				solutionBranch, branchCost := benchmarkCostBoundSearch(origin, to, neighbor.(search.HeuristicState), target, oldCost, bound, path, expansions)
				if solutionBranch && branchCost < bound {
					found = true
					bound = branchCost
					cost = bound
				}
			}
		}
	}
	return
}

func initBnBVariables(initialBound float64) (path map[search.State]search.State, bound, cost float64) {
	if bound == 0 {
		bound = math.Inf(0)
	} else {
		bound = initialBound
	}
	path = make(map[search.State]search.State)
	cost = 0
	return
}