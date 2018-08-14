package search

import (
	"github.com/christat/gost/queue"
	stack "github.com/christat/gost/stack"
	"github.com/christat/search"
)

func checkVertexAndEnqueueNeighbors(vertex, target search.State, open gost.Queue, path map[search.State]search.State) (found bool) {
	if vertex.Equals(target) {
		found = true
		return
	}
	for _, neighbor := range vertex.Neighbors() {
		_, visited := path[neighbor]
		if !visited {
			open.Enqueue(neighbor)
			path[neighbor] = vertex
		}
	}
	return
}

func checkVertexAndPushNeighbors(vertex, target search.State, open stack.Stack, path map[search.State]search.State) (found bool) {
	if vertex.Equals(target) {
		found = true
		return
	}
	for _, neighbor := range vertex.Neighbors() {
		_, visited := path[neighbor]
		if !visited {
			open.Push(neighbor)
			path[neighbor] = vertex
		}
	}
	return
}

func copyPath(original map[search.State]search.State) (copy map[search.State]search.State) {
	copy = make(map[search.State]search.State, len(original))
	for key, value := range original {
		copy[key] = value
	}
	return
}