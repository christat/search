package search

import (
"github.com/christat/gost/stack"
"github.com/christat/search"
	"fmt"
)

func IterativeDeepeningDepthFirst(domain search.Domain, origin, target interface{}, depth int) (path map[interface{}]interface{}, found bool) {
	currentDepth := 0
	open := new(gost.Stack)
	open.Push(origin)

	for open.Size() > 0 {
		vertex := open.Pop()
		currentDepth++ //TODO should depth be increased and fn recalled?
		if currentDepth > depth {
			return nil, false
		}
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
	return path, found
}
