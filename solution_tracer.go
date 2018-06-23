package search

import (
	"fmt"
	"bytes"
)

// TraceSolutionPath allows to invert/interpret the result given by a given search algorithm as a slice of state names.
func TraceSolutionPath(origin, target State, path map[State]State) (tracedPath SolutionPath, err error) {
	ancestor := target
	tracedPath = append(tracedPath, ancestor.Name())
	for !ancestor.Equals(origin) {
		// Go shadows var 'ancestor' if we access map assigning with := so we need to define var 'ok' separately
		// https://github.com/ashleymcnamara/gophers/blob/master/Facepalm_Picard_Gopher.png
		var ok bool
		ancestor, ok = path[ancestor]
		if !ok {
			return tracedPath, fmt.Errorf("ancestor %v not found in solution path", ancestor.Name())
		}
		tracedPath = append(tracedPath, ancestor.Name())
	}
	return reverse(tracedPath), nil
}

// needed to reverse a slice smh
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// SolutionPath is a convenience wrapper for []string to print a neat origin-to-target path string.
type SolutionPath []string

func (s SolutionPath) String() string {
	var buffer bytes.Buffer
	for i := range s {
		var separator string
		if i != len(s) - 1 {
			separator = " -> "
		} else {
			separator = ""
		}
		buffer.WriteString(s[i] + separator)
	}
	return buffer.String()
}