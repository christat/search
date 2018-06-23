package search

import (
	"github.com/christat/search"
)

func IterativeDeepening(origin, target search.State, maxDepth int, useNodeStack ...bool) (path map[interface{}]interface{}, found bool) {
	open := selectStackImplementation(useNodeStack...)
	open.Push(origin)

	/*currentDepth := 0
	for open.Size() > 0 {
		vertex := open.Pop().(search.State)
		currentDepth++ //TODO should depth be increased and fn recalled?
		if currentDepth > maxDepth {
			return nil, false
		}
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
	}*/
	return path, found
}

/*func DepthLimitedSearch(origin, target search.State, limit int) (state search.State) {

}*/
