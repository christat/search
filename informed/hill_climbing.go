package search

import (
	"github.com/christat/gost/stack"
	"github.com/christat/search"
	"fmt"
)

func HillClimbing(domain search.Domain, origin, target interface{}) (path map[interface{}]interface{}, found bool) {
	open := new(gost.Stack)
	open.Push(origin)

	var next interface{}
	for open.Size() > 0 {
		vertex := open.Pop() //TODO check stack push/pop cost
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
			var lowestValue float64
			_, visited := path[neighbor]
			if !visited {
				heuristic, err := domain.H(neighbor)
				if err != nil {
					fmt.Print(err)
					return nil, false
				}
				if heuristic < lowestValue {
					lowestValue = heuristic
					next = neighbor
				}
			}
		}
		open.Push(next)
		path[next] = vertex
	}
	return path, found
}
