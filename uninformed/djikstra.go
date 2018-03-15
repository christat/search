package search

import (
	"fmt"
	"github.com/christat/gost/queue"
	"github.com/christat/search"
)

// Djikstra performs heuristic search over Domains with arbitrary costs.
func Djikstra(domain search.Domain, origin, target interface{}) (path map[interface{}]interface{}, found bool) {
	open := new(gost.Queue)
	open.Enqueue(origin)

	for open.Size() > 0 {
		vertex := open.Dequeue()
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
				open.Enqueue(neighbor)
				path[neighbor] = vertex
			}
		}
	}
	return path, found
}
