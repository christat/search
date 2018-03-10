package search

import (
	"github.com/christat/gost/stack"
	"github.com/christat/search"
	"fmt"
)

func DepthFirst(domain search.Domain, origin, target interface{}) (path map[interface{}]interface{}, found bool) {
	open := new(gost.Stack)
	open.Push(origin)

	for open.Size() > 0 {
		vertex := open.Pop()
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
