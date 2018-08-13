package search

import (
	"github.com/christat/search"
	"math"
	"fmt"
)

// BranchAndBound performs depth search iteratively. An upper bound is set every time a solution is found,
// pruning costlier descendants and stopping once no better solution was found. Because of its nature (minimization of positive costs)
// it is not expected to work correctly with negative costs. Maximization problems should be redefined accordingly.
// Paramter bound can be left as default float64 (0); the algorithm will assume an initial bound of plus infinity.
func DFSBranchAndBound(origin, target search.WeightedState, bound float64) (path map[search.State]search.State, found bool, cost float64) {
	path, bound ,cost = initBnBVariables(bound)
	// The expansion order in root dictates the order of branch expansions;
	// hence the algorithm follows a left-to-right DFS "scanning" pattern.

	neighbors := origin.Neighbors()

	solutionCost := math.Inf(0)
	for _, neighbor := range neighbors {
		// because of Go's inflexible type system, neighbor must be coerced to allow access to cost/heuristic functions
		solutionPath, solutionFound, newCost := costBoundSearch(origin, neighbor.(search.HeuristicState), target, cost, bound, path)
		if solutionFound && newCost < bound && newCost < solutionCost {
			found = true
			bound = newCost
			solutionCost = newCost
			path = solutionPath
		}
	}
	return path, found, solutionCost
}

func reverse(s []search.State) []search.State {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func costBoundSearch(from, to, target search.WeightedState, branchCost, bound float64, path map[search.State]search.State) (solutionPath map[search.State]search.State, found bool, cost float64) {
	expansionCost := from.Cost(to)
	fmt.Printf("expanding %v -> %v with cost %v\n", from.Name(), to.Name(), expansionCost)
	if branchCost + expansionCost < bound {
		path[to] = from
		cost = branchCost + expansionCost
		fmt.Printf("Below bound; cost so far is %v and bound %v\n", cost, bound)
		if to.Equals(target) {
			fmt.Printf("solution found! FROM %v with COST %v\n", from.Name(), cost)
			solutionPath = copyPath(path)
			bound = cost
			found = true
		}
		for _, neighbor := range to.Neighbors() {
			// because of Go's inflexible type system, neighbor must be coerced to allow access to cost/heuristic functions
			subPath, subFound, subCost := costBoundSearch(to, neighbor.(search.HeuristicState), target, cost, bound, path)
			if subFound && subCost < bound {
				found = true
				solutionPath = subPath
				bound = subCost
			}
		}
	} else {
		fmt.Printf("ABORT: cost %v and bound %v; pruning branch..\n", branchCost + expansionCost, bound)
	}
	return
}

// Benchmark variant of DFSBranchAndBound.
// It measures execution parameters (time, nodes expanded) them in a search.AlgorithmBenchmark entity.
/*func BenchmarkDFSBranchAndBound(origin, target search.WeightedState, useNodeStack ...bool) (path map[search.State]search.State, found bool, cost float64, bench search.AlgorithmBenchmark) {
	open, path, foundPath, bound, totalCost := initBnBVariables(useNodeStack)
	open.Push(origin)

	start := time.Now()
	var expansions uint = 0

	for open.Size() > 0 {
		vertex := open.Pop().(search.WeightedState)
		expansions++
		totalCost, bound, found = evaluateNeighborsAndUpdateCost(totalCost, bound, vertex, target, found, path, foundPath, open)
	}
	elapsed := time.Since(start)
	return foundPath, found, bound, search.AlgorithmBenchmark{ElapsedTime: elapsed, TotalExpansions: expansions}
}*/

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