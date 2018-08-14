package search

import (
	stack "github.com/christat/gost/stack"
	"github.com/christat/gost/queue"
)

// SelectStackImplementation is a shared util to return a data structure implementing the gost.Stack interface.
// By default, the slice-backed stack is used. When useNodeStack is set to true, a single linked list variant is returned.
func SelectStackImplementation(useNodeStack ...bool) (stack.Stack) {
	if len(useNodeStack) > 0 && useNodeStack[0] == true {
		return new(stack.NodeStack)
	} else {
		return new(stack.SliceStack)
	}
}

// SelectQueueImplementation is a shared util to return a data structure implementing the gost.Queue interface.
// By default, the slice-backed queue is used. When useNodeQueue is set to true, a single linked list variant is returned.
func SelectQueueImplementation(useNodeQueue ...bool) (gost.Queue) {
	if len(useNodeQueue) > 0 && useNodeQueue[0] == true {
		return new(gost.NodeQueue)
	} else {
		return new(gost.SliceQueue)
	}
}