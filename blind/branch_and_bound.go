package search

/*
func upperBoundBnB(origin, target search.WeightedState, bound float64, useNodeStack bool) (path map[interface{}]interface{}, found bool) {
	open := selectStackImplementation(useNodeStack...)
	open.Push(origin)

	var totalCost float64 = 0
	for open.Size() > 0 {
		vertex := open.Pop().(search.WeightedState)
		totalCost += vertex.Cost()
		if totalCost < bound {
			if vertex == target {
				found = true
				break
			}
			neighbors := vertex.Neighbors()
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
// it is not expected to work correctly with negative costs. Therefore, maximization problems should be defined accordingly.
func BranchAndBound(origin, target search.WeightedState, useNodeStack ...bool) (path map[interface{}]interface{}, found bool) {
	var doUseNodeStack bool
	if len(useNodeStack) > 0 && useNodeStack[0] == true {
		doUseNodeStack = true
	}
	return upperBoundBnB(origin, target, math.Inf(0), doUseNodeStack)
}
*/