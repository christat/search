package search

import (
	"math"

	"fmt"
	"github.com/christat/gost/stack"
	"github.com/christat/search"
)

func upperBoundBnB(domain search.Domain, origin, target interface{}, bound float64) (path map[interface{}]interface{}, found bool) {
	var totalCost float64 = 0
	open := new(gost.Stack)
	open.Push(origin)

	for open.Size() > 0 {
		vertex := open.Pop()
		cost, err := domain.G(vertex)
		if err != nil {
			fmt.Print(err)
			return nil, false
		}
		totalCost += cost
		if totalCost < bound {
			if vertex == target {
				found = true
				break
			}
			neighbors, err := domain.Neighbors(vertex)
			if err != nil {
				fmt.Print(err)
				return nil, false
			}
			for neighbor := range neighbors {
				_, visited := path[neighbor]
				if !visited {
					open.Push(neighbor)
					path[neighbor] = vertex
				}
			}
		}
	}
	return path, found
}

// BranchAndBound performs depth search iteratively. An upper bound is set every time a solution is found,
// pruning costlier descendants and stopping once no better solution was found. Because of its nature (minimization of positive costs)
// it is not expected to work correctly with negative costs. Equivalently, maximization problems should be transformed accordingly.
func BranchAndBound(domain search.Domain, origin, target interface{}) (path map[interface{}]interface{}, found bool) {
	return upperBoundBnB(domain, origin, target, math.Inf(0))
}
