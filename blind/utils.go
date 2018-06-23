package search

import (
	"github.com/christat/gost/queue"
	stack "github.com/christat/gost/stack"
	"github.com/christat/search"
)

func selectQueueImplementation(useNodeQueue ...bool) (gost.Queue) {
	if len(useNodeQueue) > 0 && useNodeQueue[0] == true {
		return new(gost.NodeQueue)
	} else {
		return new(gost.SliceQueue)
	}
}

func checkVertexAndEnqueueNeighbors(vertex, target search.State, open gost.Queue, path map[search.State]search.State) (found bool) {
	if vertex == target {
		found = true
		return
	}
	neighbors := vertex.Neighbors()
	for _, neighbor := range neighbors {
		_, visited := path[neighbor]
		if !visited {
			open.Enqueue(neighbor)
			path[neighbor] = vertex
		}
	}
	return
}

func selectStackImplementation(useNodeStack ...bool) (stack.Stack) {
	if len(useNodeStack) > 0 && useNodeStack[0] == true {
		return new(stack.NodeStack)
	} else {
		return new(stack.SliceStack)
	}
}

func checkVertexAndPushNeighbors(vertex, target search.State, open stack.Stack, path map[search.State]search.State) (found bool) {
	if vertex == target {
		found = true
		return
	}
	neighbors := vertex.Neighbors()
	for _, neighbor := range neighbors {
		_, visited := path[neighbor]
		if !visited {
			open.Push(neighbor)
			path[neighbor] = vertex
		}
	}
	return
}